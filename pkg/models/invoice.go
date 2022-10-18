package models

type Invoice struct {
	Id string;
	CreatorName string;
	CreatorId string;
	CompanyId string;
	Items []InvoiceItem
}
