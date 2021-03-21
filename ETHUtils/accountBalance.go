package ETHUtils

import (

    "context"
    "log"
    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

)

func GetAccountBalance(addr common.Address, client *ethclient.Client ) *big.Float {

	balanceInWei, err := client.BalanceAt(context.Background(), addr, nil)

	if err != nil {
  log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balanceInWei.String())

	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue
}