package ETHUtils

import (
	
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func ConnectToEthClient() *ethclient.Client {


	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error with .env file")
	}

	id := os.Getenv("PROJECT_ID")

    client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")

  //  _ = client // we'll use this in the upcoming sections

	return client
}




//mnem route thought column mother expose coconut race able trip hope swarm pond