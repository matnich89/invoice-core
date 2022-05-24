package model

type Data struct {
	Client        string `json:"client"`
	ClientAddress string `json:"clientAddress"`
	Rate          string `json:"rate"`
	Total         string `json:"total"`
	Date          string `json:"date"`
	InvoiceNumber int    `json:"invoiceNumber"`
	Rows          []Rows `json:"rows"`
}

type Row struct {
	Dates  string `json:"dates"`
	Desc   string `json:"desc"`
	Hours  string `json:"hours"`
	Rate   string `json:"rate"`
	Amount string `json:"amount"`
}
