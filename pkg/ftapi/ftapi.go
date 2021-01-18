package ftapi

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// FtAPI This is a struct to send authenticated requests to the 42 API
type FtAPI struct {
	apiEndpoint string
	httpClient *http.Client
}

// New Creates an FtAPI instance
func New(apiEndpoint string, authenticatedClient *http.Client) *FtAPI  {
	return &FtAPI{
		apiEndpoint: apiEndpoint,
		httpClient:  authenticatedClient,
	}
}

// NewFromCredentials Creates an FtAPI instance with an authenticated client using the given oAuth2 credentials
func NewFromCredentials(apiEndpoint string, oauthCredentials *clientcredentials.Config) *FtAPI {
	ctx := context.Background()
	authenticatedClient := oauthCredentials.Client(ctx)
	return New(apiEndpoint, authenticatedClient)
}