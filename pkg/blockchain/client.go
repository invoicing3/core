package blockchain

import "github.com/ethereum/go-ethereum/ethclient"

func GetEthClient() *ethclient.Client {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		panic(err)
	}
	return client;
}