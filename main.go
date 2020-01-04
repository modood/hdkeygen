package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {
	compress := flag.Bool("c", false, "generate a compressed public key")
	number := flag.Int("n", 20, "set number of keys to generate")
	flag.Parse()

	fmt.Printf("\n%-34s %-s\n", "Bitcoin address", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 86))

	for i := 0; i < *number; i++ {
		wif, addr, err := generate(*compress)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-34s %s\n", addr, wif)
	}
	fmt.Println()
}

func generate(compress bool) (wif, addr string, err error) {
	prvKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}

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
