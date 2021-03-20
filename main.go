package main

import (

	"fmt"
    //"log"
	"github.com/Xowmandon/XowGeth/ETHUtils"

    //"github.com/ethereum/go-ethereum/ethclient"

   // "client.go"
)

func main() {


	//connects to eth Client, returns client/connection.
	//function stored in client.go
	client := ConnectToEthClient()

	fmt.Printf("v is of type %T\n", client)

	//addrr := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"


	//converts address to common.address type in go-eth
	//allows for member functions to act on the address
	//addr := getAddress(addrr)


	//balance := getAccountBalance(addr, client)

	//fmt.Println(balance)

	generateWallet()

	//firstWallet :=  "./wallets/UTC--2021-03-20T11-27-57.093412400Z--58f90b71050621d5627767de0a7f76505f0337de"
	password := "SECRET"

	//gens new keystore file and returns public address of new account
    pubAddress := GenerateKeystore(password)

	fmt.Println("PUBLIC ADDRESS PF ACCOUNT :", pubAddress)

	//importKs(firstWallet, password)

	NewHDWallet("joshua")

}