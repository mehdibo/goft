package ftapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// APIInterface interface for a struct that talks to the 42 API
type APIInterface interface {
	Get(url string) (*http.Response, error)
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
	PostJSON(url string, data interface{}) (resp *http.Response, err error)
	Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error)
	PatchJSON(url string, data interface{}) (resp *http.Response, err error)
	CreateUser(user *User) error
	SetUserImage(login string, img *os.File) error
	CreateClose(close *Close) error
}

// API This is a struct to send authenticated requests to the 42 API
type API struct {
	apiEndpoint string
	httpClient *http.Client
}

// New Creates an API instance
func New(apiEndpoint string, authenticatedClient *http.Client) APIInterface  {
	return &API{
		apiEndpoint: apiEndpoint,
		httpClient:  authenticatedClient,
	}
}

// NewFromCredentials Creates an API instance with an authenticated client using the given oAuth2 credentials
func NewFromCredentials(apiEndpoint string, oauthCredentials *clientcredentials.Config) APIInterface {
	ctx := context.Background()
	authenticatedClient := oauthCredentials.Client(ctx)
	return New(apiEndpoint, authenticatedClient)
}

func (ft *API) newRequest(method string, contentType string, url string, body io.Reader) (*http.Request, error)  {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	return req, nil
}

// Execute the request
func (ft *API) do(req *http.Request) (*http.Response, error)  {
	for {
		resp, err := ft.httpClient.Do(req)
		if err != nil {
			return resp, err
		}
		if resp.StatusCode == http.StatusTooManyRequests {
			// Check if exceeded max hourly rate
			retryAfter, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
			if err != nil {
				retryAfter = 0
			}
			if retryAfter <= 0 {
				return nil, errors.New("exceeded rate limit")
			}
			// We wait for the duration set by the header
			time.Sleep(time.Duration(retryAfter) * time.Second)
			continue
		}
		return resp,err
	}
}

// Get sends a get request to the given URL
func (ft *API) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", ft.apiEndpoint+url, nil)
	if err != nil {
		return nil, err
	}
	return ft.do(req)
}

// Post sends a POST request to the given url
func (ft *API) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)  {
	req, err := ft.newRequest("POST", contentType, ft.apiEndpoint+url, body)
	if err != nil {
		return nil, err
	}
	return ft.do(req)
}

// Patch sends a PATCH request to the given url
func (ft *API) Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := ft.newRequest("PATCH", contentType, ft.apiEndpoint+url, body)
	if err != nil {
		return nil, err
	}
	return ft.do(req)
}

// PatchJSON this method will automatically turn data into a json and send a PATCH request to the given url
func (ft *API) PatchJSON(url string, data interface{}) (resp *http.Response, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return ft.Patch(url, "application/json", bytes.NewReader(jsonData))
}

// PostJSON this method will automatically turn data into a json and send a post request to the given url
func (ft *API) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return ft.Post(url, "application/json", bytes.NewReader(jsonData))
}

// TODO: better handle errors

// CreateUser creates a new user and sets `user` id and url to the one returned by the API
func (ft *API) CreateUser(user *User) error  {
	payload := map[string]User{
		"user": *user,
	}
	resp, err := ft.PostJSON("/users", payload)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return errors.New("failed creating user")
	}
	var createdUser User
	_ = json.NewDecoder(resp.Body).Decode(&createdUser)
	user.ID = createdUser.ID
	user.URL = createdUser.URL
	return nil
}

// SetUserImage set a profile image to the user
func (ft *API) SetUserImage(login string, img *os.File) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("user[image]", filepath.Base(img.Name()))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, img)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	resp, err := ft.Patch("/users/"+login, writer.FormDataContentType(), body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		switch resp.StatusCode {
		case http.StatusNotFound:
			return errors.New("user not found")
		default:
			return errors.New("failed setting profile image")
		}
	}
	return nil
}

func (ft *API) CreateClose(close *Close) error  {
	payload := map[string]map[string]interface{}{
		"close": {
			"kind": close.Kind,
			"reason": close.Reason,
		},
	}
	resp, err := ft.PostJSON("/users/"+close.User.Login+"/closes", payload)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusCreated {
		return nil
	}
	switch resp.StatusCode {
	case http.StatusNotFound:
		return errors.New("user not found")
	default:
		return errors.New("failed creating close")
	}
}