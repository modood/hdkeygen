package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

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
const (
	Apostrophe uint32 = 0x80000000 // 0'
	Purpose    uint32 = 0x8000002C // 44'
)

// CoinType SLIP-0044 : Registered coin types for BIP-0044
// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
type CoinType = uint32

const (
	CoinTypeBTC CoinType = 0x80000000
	CoinTypeLTC CoinType = 0x80000002
	CoinTypeETH CoinType = 0x8000003c
	CoinTypeEOS CoinType = 0x800000c2
)

type Key struct {
	path     string
	bip32Key *bip32.Key
}

func (k *Key) GetWIF(compress bool) (wif, address string, err error) {
	prvKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), k.bip32Key.Key)
	return GenerateFromBytes(prvKey, compress)
}

// https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
// bip44 define the following 5 levels in BIP32 path:
// m / purpose' / coin_type' / account' / change / address_index

func (k *Key) GetPath() string {
	return k.path
}

type KeyManager struct {
	mnemonic   string
	passphrase string
	keys       map[string]*bip32.Key
	mux        sync.Mutex
}

func NewKeyManager(bitSize int, passphrase string) (*KeyManager, error) {
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	km := &KeyManager{
		mnemonic:   mnemonic,
		passphrase: passphrase,
		keys:       make(map[string]*bip32.Key, 0),
	}
	return km, nil
}

func (km *KeyManager) GetMnemonic() string {
	return km.mnemonic
}

func (km *KeyManager) GetPassphrase() string {
	return km.passphrase
}

func (km *KeyManager) GetSeed() []byte {
	return bip39.NewSeed(km.GetMnemonic(), km.GetPassphrase())
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

func (km *KeyManager) GetPurposeKey() (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'`, Purpose-Apostrophe)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetMasterKey()
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(Purpose)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetCoinTypeKey(coinType CoinType) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'`, Purpose-Apostrophe, coinType-Apostrophe)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetPurposeKey()
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

func (km *KeyManager) GetAccountKey(coinType CoinType, account uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'`, Purpose-Apostrophe, coinType-Apostrophe, account)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetCoinTypeKey(coinType)
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
func (km *KeyManager) GetChangeKey(coinType CoinType, account, change uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d`, Purpose-Apostrophe, coinType-Apostrophe, account, change)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetAccountKey(coinType, account)
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

func (km *KeyManager) GetKey(coinType CoinType, account, change, index uint32) (*Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d/%d`, Purpose-Apostrophe, coinType-Apostrophe, account, change, index)

	key, ok := km.getKey(path)
	if ok {
		return &Key{path: path, bip32Key: key}, nil
	}

	parent, err := km.GetChangeKey(coinType, account, change)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(index)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return &Key{path: path, bip32Key: key}, nil
}

func main() {
	compress := flag.Bool("compress", true, "generate a compressed public key")
	bip39 := flag.Bool("bip39", false, "mnemonic code for generating deterministic keys")
	pass := flag.String("pass", "", "protect bip39 mnemonic with a passphrase")
	number := flag.Int("n", 20, "set number of keys to generate")

	flag.Parse()

	if *bip39 {
		km, err := NewKeyManager(128, *pass)
		if err != nil {
			log.Fatal(err)
		}
		masterKey, err := km.GetMasterKey()
		if err != nil {
			log.Fatal(err)
		}
		accountKey, err := km.GetAccountKey(CoinTypeBTC, 0)
		if err != nil {
			log.Fatal(err)
		}
		changeKey, err := km.GetChangeKey(CoinTypeBTC, 0, 0)
		if err != nil {
			log.Fatal(err)
		}
		passphrase := km.GetPassphrase()
		if passphrase == "" {
			passphrase = "<none>"
		}
		fmt.Printf("\n%-35s %s\n", "BIP39 Mnemonic:", km.GetMnemonic())
		fmt.Printf("%-35s %s\n", "BIP39 Passphrase:", passphrase)
		fmt.Printf("%-35s %x\n", "BIP39 Seed:", km.GetSeed())
		fmt.Printf("%-35s %-16s %s\n", "BIP32 Root Key:", "m", masterKey.B58Serialize())
		fmt.Printf("%-35s %-16s %s\n", "BIP44 Account Extended Private Key:", "m/44'/0'/0'", accountKey.B58Serialize())
		fmt.Printf("%-35s %-16s %s\n", "BIP44 Account Extended Public  Key:", "", accountKey.PublicKey().B58Serialize())
		fmt.Printf("%-35s %-16s %s\n", "BIP32 Extended Private Key:", "m/44'/0'/0'/0", changeKey.B58Serialize())
		fmt.Printf("%-35s %-16s %s\n", "BIP32 Extended Public  Key:", "", changeKey.PublicKey().B58Serialize())

		fmt.Printf("\n%-18s %-34s %-s\n", "Path", "Bitcoin address", "WIF(Wallet Import Format)")
		fmt.Println(strings.Repeat("-", 106))
		for i := 0; i < *number; i++ {
			key, err := km.GetKey(CoinTypeBTC, 0, 0, uint32(i))
			if err != nil {
				log.Fatal(err)
			}
			wif, address, err := key.GetWIF(*compress)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%-18s %-34s %s\n", key.GetPath(), address, wif)
		}
		fmt.Println()
		return
	}

	fmt.Printf("\n%-34s %-s\n", "Bitcoin address", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 86))

	for i := 0; i < *number; i++ {
		wif, addr, err := Generate(*compress)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-34s %s\n", addr, wif)
	}
	fmt.Println()
}

func Generate(compress bool) (wif, addr string, err error) {
	prvKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}
	return GenerateFromBytes(prvKey, compress)
}

func GenerateFromBytes(prvKey *btcec.PrivateKey, compress bool) (wif, addr string, err error) {
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, compress)
	if err != nil {
		return "", "", err
	}

	addressPubKey, err := btcutil.NewAddressPubKey(btcwif.SerializePubKey(), &chaincfg.MainNetParams)
	if err != nil {
		return "", "", err
	}

	wif = btcwif.String()
	addr = addressPubKey.EncodeAddress()

	return wif, addr, nil
}
