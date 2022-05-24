package processor

import (
	invoiceproto "github.com/matnich89/invoiceproto/gen/go/invoice/v1"
	"invoice-core/cmd/server/internal/model"
	"strconv"
)

type Processor struct {
	client invoiceproto.InvoiceServiceClient
}

func (p *Processor) Process(data model.Data) {
}

func (p *Processor) buildInvoice(data model.Data) {
	var invoiceRows []invoiceproto.InvoiceRow = make([]invoiceproto.InvoiceRow, 0)
	for _, val := range data.Rows {
		invoiceRows = append()
	}

	invoice := invoiceproto.Invoice{
		Client:        data.Client,
		Address:       data.ClientAddress,
		Rate:          data.Rate,
		Total:         data.Total,
		Date:          data.Date,
		InvoiceNumber: strconv.Itoa(data.InvoiceNumber),
		Rows:          data.Rows,
	}
}

func (p *Processor) buildInvoiceRow(data model.Rows) invoiceproto.InvoiceRow {

}
