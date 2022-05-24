package model

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonFile, err := ioutil.ReadFile("./testdata/testinvoice.json")
	require.NoError(t, err)
	var data Data
	json.Unmarshal(jsonFile, &data)
	require.NotNilf(t, data, "data should not nil")
}
