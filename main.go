package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"

	btcec "github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
)

// Purpose BIP43 - Purpose Field for Deterministic Wallets
// https://github.com/bitcoin/bips/blob/master/bip-0043.mediawiki
//
// Purpose is a constant set to 44' (or 0x8000002C) following the BIP43 recommendation.
// It indicates that the subtree of this node is used according to this specification.
//
// What does 44' mean in BIP44?
// https://bitcoin.stackexchange.com/questions/74368/what-does-44-mean-in-bip44
//
// 44' means that hardened keys should be used. The distinguisher for whether
// a key a given index is hardened is that the index is greater than 2^31,
// which is 2147483648. In hex, that is 0x80000000. That is what the apostrophe (') means.
// The 44 comes from adding it to 2^31 to get the final hardened key index.
// In hex, 44 is 2C, so 0x80000000 + 0x2C = 0x8000002C.
type Purpose = uint32

const (
	PurposeBIP44 Purpose = 0x8000002C // 44' BIP44
	PurposeBIP49 Purpose = 0x80000031 // 49' BIP49
	PurposeBIP84 Purpose = 0x80000054 // 84' BIP84
	PurposeBIP86 Purpose = 0x80000056 // 86' BIP86 //taprrot
)

// CoinType SLIP-0044 : Registered coin types for BIP-0044
// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
type CoinType = uint32

const (
	CoinTypeBTC CoinType = 0x80000000
	CoinTypeETH CoinType = 0x8000003c
)

const (
	Apostrophe uint32 = 0x80000000 // 0'
)

type Key struct {
	Path     string
	bip32Key *bip32.Key
}

func (k *Key) Calculate(compress bool) (wif, address, segwitBech32, segwitNested, taproot string, err error) {
	prvKey, _ := btcec.PrivKeyFromBytes(k.bip32Key.Key)
	return CalculateFromPrivateKey(prvKey, compress)
}

// https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
// bip44 define the following 5 levels in BIP32 path:
// m / purpose' / coin_type' / account' / change / address_index

type KeyManager struct {
	Mnemonic   string
	Passphrase string
	keys       map[string]*bip32.Key
	mux        sync.Mutex
}

// NewKeyManager return new key manager
// if mnemonic is not provided, it will generate a new mnemonic with 128 bits of entropy, which is 12 words
func NewKeyManager(mnemonic, passphrase string) (*KeyManager, error) {
	if mnemonic == "" {
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			return nil, err
		}
		mnemonic, err = bip39.NewMnemonic(entropy)
		if err != nil {
			return nil, err
		}
	}

	km := &KeyManager{
		Mnemonic:   mnemonic,
		Passphrase: passphrase,
		keys:       make(map[string]*bip32.Key, 0),
	}
	return km, nil
}

func (km *KeyManager) GetSeed() []byte {
	return bip39.NewSeed(km.Mnemonic, km.Passphrase)
}

func (km *KeyManager) getKey(path string) (*bip32.Key, bool) {
	km.mux.Lock()
	defer km.mux.Unlock()

	key, ok := km.keys[path]
	return key, ok
}

func (km *KeyManager) setKey(path string, key *bip32.Key) {
	km.mux.Lock()
	defer km.mux.Unlock()

	km.keys[path] = key
}

func (km *KeyManager) GetMasterKey() (*bip32.Key, error) {
	path := "m"

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	key, err := bip32.NewMasterKey(km.GetSeed())
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetPurposeKey(purpose uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'`, purpose-Apostrophe)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetMasterKey()
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(purpose)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetCoinTypeKey(purpose, coinType uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'`, purpose-Apostrophe, coinType-Apostrophe)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetPurposeKey(purpose)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(coinType)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetAccountKey(purpose, coinType, account uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'`, purpose-Apostrophe, coinType-Apostrophe, account)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetCoinTypeKey(purpose, coinType)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(account + Apostrophe)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

// GetChangeKey ...
// https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki#change
// change constant 0 is used for external chain
// change constant 1 is used for internal chain (also known as change addresses)
func (km *KeyManager) GetChangeKey(purpose, coinType, account, change uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d`, purpose-Apostrophe, coinType-Apostrophe, account, change)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetAccountKey(purpose, coinType, account)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(change)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetKey(purpose, coinType, account, change, index uint32) (*Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d/%d`, purpose-Apostrophe, coinType-Apostrophe, account, change, index)

	key, ok := km.getKey(path)
	if ok {
		return &Key{Path: path, bip32Key: key}, nil
	}

	parent, err := km.GetChangeKey(purpose, coinType, account, change)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(index)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return &Key{Path: path, bip32Key: key}, nil
}

func CalculateFromPrivateKey(prvKey *btcec.PrivateKey, compress bool) (wif, address, segwitBech32, segwitNested, taproot string, err error) {
	// generate the wif(wallet import format) string
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, compress)
	if err != nil {
		return "", "", "", "", "", err
	}
	wif = btcwif.String()

	// generate a normal p2pkh address
	serializedPubKey := btcwif.SerializePubKey()
	addressPubKey, err := btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", "", "", "", err
	}
	address = addressPubKey.EncodeAddress()

	// generate a normal p2wkh address from the pubkey hash
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", "", "", "", err
	}
	segwitBech32 = addressWitnessPubKeyHash.EncodeAddress()

	// generate an address which is
	// backwards compatible to Bitcoin nodes running 0.6.0 onwards, but
	// allows us to take advantage of segwit's scripting improvments,
	// and malleability fixes.
	serializedScript, err := txscript.PayToAddrScript(addressWitnessPubKeyHash)
	if err != nil {
		return "", "", "", "", "", err
	}
	addressScriptHash, err := btcutil.NewAddressScriptHash(serializedScript, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", "", "", "", err
	}
	segwitNested = addressScriptHash.EncodeAddress()

	// generate a taproot address
	tapKey := txscript.ComputeTaprootKeyNoScript(prvKey.PubKey())
	addressTaproot, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(tapKey), &chaincfg.MainNetParams)
	if err != nil {
		return "", "", "", "", "", err
	}
	taproot = addressTaproot.EncodeAddress()

	return wif, address, segwitBech32, segwitNested, taproot, nil
}

func main() {
	compress := true // generate a compressed public key

	pass := flag.String("pass", "", "protect bip39 mnemonic with a passphrase")
	number := flag.Int("n", 10, "set number of keys to generate")
	mnemonic := flag.String("mnemonic", "", "optional list of words to re-generate a root key")

	wifInput := flag.String("wif", "", "decode the private key from wif, then generate the bitcoin address.")

	flag.Parse()

	if *wifInput != "" {
		wif, err := btcutil.DecodeWIF(*wifInput)
		if err != nil {
			log.Fatal(err)
		}

		wifCompressed, addressCompressed, segwitBech32, segwitNested, taproot, err := CalculateFromPrivateKey(wif.PrivKey, true)
		if err != nil {
			log.Fatal(err)
		}

		wifUncompressed, addressUncompressed, _, _, _, err := CalculateFromPrivateKey(wif.PrivKey, false)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n Wallet Import Format:")
		fmt.Printf(" *   %-24s %s\n", "WIF(compressed):", wifCompressed)
		fmt.Printf(" *   %-24s %s\n", "WIF(uncompressed):", wifUncompressed)

		fmt.Println("\n Public Addresses:")
		fmt.Printf(" *   %-24s %s\n", "Legacy(compresed):", addressCompressed)
		fmt.Printf(" *   %-24s %s\n", "Legacy(uncompressed):", addressUncompressed)
		fmt.Printf(" *   %-24s %s\n", "SegWit(nested):", segwitNested)
		fmt.Printf(" *   %-24s %s\n", "SegWit(bech32):", segwitBech32)
		fmt.Printf(" *   %-24s %s\n", "Taproot(bech32m):", taproot)
		fmt.Println()
		return
	}

	km, err := NewKeyManager(*mnemonic, *pass)
	if err != nil {
		log.Fatal(err)
	}
	masterKey, err := km.GetMasterKey()
	if err != nil {
		log.Fatal(err)
	}
	passphrase := km.Passphrase
	if passphrase == "" {
		passphrase = "<none>"
	}
	fmt.Printf("\n%-18s %s\n", "BIP39 Mnemonic:", km.Mnemonic)
	fmt.Printf("%-18s %s\n", "BIP39 Passphrase:", passphrase)
	fmt.Printf("%-18s %x\n", "BIP39 Seed:", km.GetSeed())
	fmt.Printf("%-18s %s\n", "BIP32 Root Key:", masterKey.B58Serialize())

	fmt.Printf("\n%-18s %-34s %-52s\n", "Path(BIP44)", "Legacy(P2PKH, compresed)", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 106))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP44, CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		wif, address, _, _, _, err := key.Calculate(compress)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-18s %-34s %s\n", key.Path, address, wif)
	}

	fmt.Printf("\n%-18s %-34s %s\n", "Path(BIP49)", "SegWit(P2WPKH-nested-in-P2SH)", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 106))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP49, CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		wif, _, _, segwitNested, _, err := key.Calculate(compress)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-18s %s %s\n", key.Path, segwitNested, wif)
	}

	fmt.Printf("\n%-18s %-42s %s\n", "Path(BIP84)", "SegWit(P2WPKH, bech32)", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 114))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP84, CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		wif, _, segwitBech32, _, _, err := key.Calculate(compress)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-18s %s %s\n", key.Path, segwitBech32, wif)
	}

	fmt.Printf("\n%-18s %-62s %s\n", "Path(BIP86)", "Taproot(P2TR, bech32m)", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 134))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP86, CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		wif, _, _, _, taproot, err := key.Calculate(compress)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-18s %s %s\n", key.Path, taproot, wif)
	}

	fmt.Printf("\n%-18s %-42s %-52s\n", "Path(BIP44)", "Ethereum(EIP55)", "Private Key(hex)")
	fmt.Println(strings.Repeat("-", 126))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP44, CoinTypeETH, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}

		address := ethereumAddress(key.bip32Key.Key)
		fmt.Printf("%-18s %s %x\n", key.Path, address, key.bip32Key.Key)
	}

	fmt.Println()
}

// ethereumAddress generates an ethereum address from a private key.
// The private key must be 32 bytes. The address is returned with the 0x prefix and in EIP55 checksum format.
func ethereumAddress(privateKeyBytes []byte) (address string) {
	_, pubKey := btcec.PrivKeyFromBytes(privateKeyBytes)

	// Public ECDSA Key
	publicKey := pubKey.ToECDSA()

	// ethereum public key must be 64byte (32byte x 32byte y coordinates
	// this is uncompressed ECDSA public key without 04 prefix
	publicKeyBytes := append(publicKey.X.FillBytes(make([]byte, 32)), publicKey.Y.FillBytes(make([]byte, 32))...)

	// Keccak-256 hash of the public key
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes)
	addr := hash.Sum(nil)

	// Ethereum uses the last 20 bytes of the Keccak-256 hash of the public key
	// this is ethereum address(without 0x prefix) but currently have not checksum
	addr = addr[len(addr)-20:]

	return eip55checksum(fmt.Sprintf("0x%x", addr))
}

// eip55checksum implements the EIP55 checksum address encoding.
// https://github.com/ethereum/ercs/blob/master/ERCS/erc-55.md
// In English, convert the address to hex, but if the i th digit is a letter (ie. it's one of abcdef)
// print it in uppercase if the 4*i th bit of the hash of the lowercase hexadecimal address is 1 otherwise print it in lowercase.
// this function is copied from the go-ethereum library: go-ethereum/common/types.go checksumHex method
func eip55checksum(address string) string {
	buf := []byte(strings.ToLower(address))
	sha := sha3.NewLegacyKeccak256()
	sha.Write(buf[2:])
	hash := sha.Sum(nil)
	for i := 2; i < len(buf); i++ {
		hashByte := hash[(i-2)/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if buf[i] > '9' && hashByte > 7 {
			buf[i] -= 32
		}
	}
	return string(buf[:])
}
