package issuedmrtemp_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	issuedmrtemp "github.com/oapi-codegen/oapi-codegen/v2/internal/test/issues/issue-dmrtemp"
)

type fakeHttpClient struct {
	capturedRequest *http.Request
}

func (f *fakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	f.capturedRequest = req
	return nil, assert.AnError
}

func TestClient_GetRoot_DoesNotEmitQueryParamForOptionalList(t *testing.T) {
	t.Parallel()

	httpClient := &fakeHttpClient{}
	client, err := issuedmrtemp.NewClient("http://example.test", issuedmrtemp.WithHTTPClient(httpClient))
	require.NoError(t, err)

	_, _ = client.GetRoot(t.Context(), &issuedmrtemp.GetRootParams{
		Id: ptr("some-id"),
	})

	require.NotNil(t, httpClient.capturedRequest)

	expectedParams := url.Values{
		"id": []string{"some-id"},
	}
	assert.Equal(t, expectedParams, httpClient.capturedRequest.URL.Query())
	expectedParamString := "id=some-id"
	assert.Equal(t, expectedParamString, httpClient.capturedRequest.URL.RawQuery)
}

func ptr[T any](v T) *T {
	return &v
}
