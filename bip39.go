package main

import (

  "fmt"
  "github.com/tyler-smith/go-bip39"
  "github.com/tyler-smith/go-bip32"
)

func generateMnemonic() string {

  entropy, _ := bip39.NewEntropy(256)
  mnemonic, _ := bip39.NewMnemonic(entropy)

  return mnemonic

}

func NewHDWallet( passphrase string) {

	mnemonic := generateMnemonic()

	seed := bip39.NewSeed(mnemonic, passphrase)

	masterKey, _ := bip32.NewMasterKey(seed)

	publicKey :=  masterKey.PublicKey()

	fmt.Println("Mnemonic: ", mnemonic)
 	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)

}

