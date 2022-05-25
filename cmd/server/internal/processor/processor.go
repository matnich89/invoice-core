package processor

import (
	invoiceproto "github.com/matnich89/invoiceproto/gen/go/invoice/v1"
	"invoice-core/cmd/server/internal/logger"
)

type Processor struct {
	client invoiceproto.InvoiceServiceClient
	logger *logger.Logger
}

func (p *Processor) Process(data *invoiceproto.Invoice) {

}
