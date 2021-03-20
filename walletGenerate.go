package main

import (
	"crypto/ecdsa" //https://golang.org/pkg/crypto/ecdsa/

	"fmt"
	"log"

	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore" //for generating keystores and accouns
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	//"golang.org/x/crypto/sha3" used for manual keccak256 hashing
)

type Wallet struct {

	PubKey string  `json:"Public Key"`
	PrivKey string `json:"Private Key"`
	PubAddress string `json:"Public Address"`


}


func generateWallet() {

   
	privateKey, err := crypto.GenerateKey()

	if err != nil {

		log.Fatal(err)

	}

	//turns privatekey in bytes to then conver to hex
	privateKeyInBytes := crypto.FromECDSA(privateKey)

	//makes privateKey to a hexString, with the 0x removed
	privateKeyHexString := hexutil.Encode(privateKeyInBytes[2:])

	//retrieves derived public key from generated private key
	publicKey := privateKey.Public()

	fmt.Println("\nGenerated Private Key Of Hex Value....")
	fmt.Println(privateKeyHexString)

	publicKeyECDSA, success := publicKey.(*ecdsa.PublicKey)

	if !success {

		log.Fatal("Error with publicKey ECDSA function")

	}

	publicKeyInBytes := crypto.FromECDSAPub(publicKeyECDSA)


	//omits 0x and first two characters, encodes in hex form
	publicKeyHex := hexutil.Encode(publicKeyInBytes[4:])

	fmt.Println("\nGenerated Public Key with the Hex Value....")
	fmt.Println(publicKeyHex)


	//generatinng public address
	//utilizs go-eth crypto PubkeyToAddress
	//hashes public key with keccak256 hash,
	//takes last 40 characters and prefixes with 0x
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println("\nPublic Key Address Below")
	fmt.Println(address)

	//hash -- Keccak256 hash of the generated address
	//--Manual version below
	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(publicKeyInBytes[1:])
	//address := hexutil.Envode(hash.Sum(nil)[12:])))


	//creates new wallet instance
	//allows easy export to JSON 
	wal := Wallet {

		PubKey: publicKeyHex,

		PrivKey: privateKeyHexString,

		PubAddress: address,


	} 

	file, _ := json.MarshalIndent(wal,"","")
 
	_ = ioutil.WriteFile("wallets.json", file, 0700)

	
}

func GenerateKeystore(passW string) accounts.Account { 

	ks := keystore.NewKeyStore("./users/wallets/nondetermWallets", keystore.StandardScryptN, keystore.StandardScryptP )

	account, err := ks.NewAccount(passW)

	if err != nil {

		log.Fatal(err)

	}

	//fmt.Println(account.Address.Hex())
	return account

}

func importKs(path string, passW string) {


    ks := keystore.NewKeyStore("./users/wallets/tmpNonDetermWallets", keystore.StandardScryptN, keystore.StandardScryptP)
    jsonBytes, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    account, err := ks.Import(jsonBytes, passW, passW)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

    if err := os.Remove(path); err != nil {
        log.Fatal(err)
    }






} 