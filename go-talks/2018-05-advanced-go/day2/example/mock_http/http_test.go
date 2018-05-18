package mock_http_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/mercadolibre/advanced-go/2018-05-advanced-go/day2/example/mock_http"
	"github.com/stretchr/testify/require"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestFetchISO(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	j := `[{"id": 1, "name": "My Great Article"}]`
	mockRes, _ := httpmock.NewJsonResponder(200, json.RawMessage([]byte(j)))
	httpmock.RegisterResponder("GET", mock_http.ISO, mockRes)

	res, err := mock_http.FetchISO()
	require.NoError(t, err)

	require.Equal(t, 200, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)
	defer res.Body.Close()

	require.JSONEq(t, j, string(body))
}
