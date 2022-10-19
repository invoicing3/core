package main

import (
	"log"

	"github.com/invoicing3/core/pkg/blockchain"
	"github.com/invoicing3/core/pkg/models"
)
var items = []models.InvoiceItem{
	{
		Id: string("c8623bc0-4f27-11ed-bdc3-0242ac120002"),
		ItemName: "Product 1",
		Price: 35.25,
		Quantity: 5,
		Total: 35.25 * 5,	
	},
}
var invoice = models.Invoice{
	Id: "57131160-4f27-11ed-bdc3-0242ac120002",
	CreatorId: "66cc396a-4f27-11ed-bdc3-0242ac120002",
	CompanyId: "6d206d90-4f27-11ed-bdc3-0242ac120002",
	Items: items,
}
func main() {
	ethClient := blockchain.GetEthClient()

	log.Fatal(blockchain.GetAccountAddress(ethClient))
}