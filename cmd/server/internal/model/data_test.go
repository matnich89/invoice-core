package model

import (
	"encoding/json"
	invoicepb "github.com/matnich89/invoiceproto/gen/go/invoice/v1"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonFile, err := ioutil.ReadFile("./testdata/testinvoice.json")
	require.NoError(t, err)
	var data invoicepb.Invoice
	json.Unmarshal(jsonFile, &data)
	require.NotNilf(t, &data, "data should not nil")
}
