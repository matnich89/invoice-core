package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	invoiceproto "github.com/matnich89/invoiceproto/gen/go/invoice/v1"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"invoice-core/cmd/server/internal/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	zapLog, _ := zap.NewProduction()
	handler := New(logger.New(zapLog))
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	handler.HealthCheck(c)
	require.Equal(t, http.StatusOK, recorder.Code)
}

func TestInvoiceCreate(t *testing.T) {
	zapLog, _ := zap.NewProduction()
	handler := New(logger.New(zapLog))
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonFile, err := ioutil.ReadFile("./testdata/testinvoice.json")
	require.NoError(t, err)
	var data invoiceproto.Invoice
	err = json.Unmarshal(jsonFile, &data)
	require.NoError(t, err)
	jsonbytes, err := json.Marshal(&data)
	c.Request = &http.Request{}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	handler.CreateInvoice(c)
	require.Equal(t, http.StatusOK, recorder.Code)
}
