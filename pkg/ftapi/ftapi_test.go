package ftapi

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
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