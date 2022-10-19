package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

func GetAccountAddress(client gethclient.Client) string {
	address :=  common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	return address.Hex();
}