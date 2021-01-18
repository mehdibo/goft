package ftapi

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	ftAPI := New("https://api.intra.42.fr", &http.Client{})
	assert.IsType(t, ftAPI, &FtAPI{}, ftAPI)
}

func TestNewFromCredentials(t *testing.T) {
	ftAPI := NewFromCredentials("https://api.intra.42.fr", &clientcredentials.Config{})
	assert.IsType(t, ftAPI, &FtAPI{}, ftAPI)
}

func getBody(body io.ReadCloser) string {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return ""
	}
	return string(bodyBytes)
}

func TestGet(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/v2/users", req.URL.String())
		rw.Header().Add("X-Test", "test_value")
		_, _ = rw.Write([]byte(`OK`))
	}))
	defer server.Close()
	ftAPI := New(server.URL, server.Client())
	resp, err := ftAPI.Get("/v2/users")
	assert.Nil(t, err)
	assert.Equal(t, "OK", getBody(resp.Body))
	assert.Equal(t, "test_value", resp.Header.Get("X-Test"))
}

func TestPost(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/v1/users", req.URL.String())
		assert.Equal(t, "custom/content-type", req.Header.Get("Content-Type"))
		assert.Equal(t, "this_is_the_body", getBody(req.Body))
		rw.Header().Add("X-Test", "test_value")
		_, _ = rw.Write([]byte(`OK`))
	}))
	defer server.Close()
	ftAPI := New(server.URL, server.Client())
	body := bytes.NewReader([]byte("this_is_the_body"))
	resp, err := ftAPI.Post("/v1/users", "custom/content-type", body)
	assert.Nil(t, err)
	assert.Equal(t, "OK", getBody(resp.Body))
	assert.Equal(t, "test_value", resp.Header.Get("X-Test"))
}