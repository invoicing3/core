package core

type InvoiceItem struct {
	id string;
	item_name string;
	quantity float32
	price float64
	total float32
}

type Invoice struct {
	id string;
	creator_name string;
	creator_id string;
	company_id string;
	items []InvoiceItem
}
