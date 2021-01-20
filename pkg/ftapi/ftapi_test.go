package ftapi

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	ftAPI := New("https://api.intra.42.fr", &http.Client{})
	assert.IsType(t, ftAPI, &API{}, ftAPI)
}

func TestNewFromCredentials(t *testing.T) {
	ftAPI := NewFromCredentials("https://api.intra.42.fr", &clientcredentials.Config{})
	assert.IsType(t, ftAPI, &API{}, ftAPI)
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
		assert.Equal(t, "GET", req.Method)
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
		assert.Equal(t, "POST", req.Method)
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

func TestPatch(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "PATCH", req.Method)
		assert.Equal(t, "/v1/users", req.URL.String())
		assert.Equal(t, "custom/content-type", req.Header.Get("Content-Type"))
		assert.Equal(t, "this_is_the_body", getBody(req.Body))
		rw.Header().Add("X-Test", "test_value")
		_, _ = rw.Write([]byte(`OK`))
	}))
	defer server.Close()
	ftAPI := New(server.URL, server.Client())
	body := bytes.NewReader([]byte("this_is_the_body"))
	resp, err := ftAPI.Patch("/v1/users", "custom/content-type", body)
	assert.Nil(t, err)
	assert.Equal(t, "OK", getBody(resp.Body))
	assert.Equal(t, "test_value", resp.Header.Get("X-Test"))
}

type testData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Array []string `json:"array"`
}

func TestPostJson(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method)
		assert.Equal(t, "/v1/users", req.URL.String())
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t, "{\"id\":10,\"name\":\"Spoody\",\"array\":[\"test_1\",\"test_2\"]}", getBody(req.Body))
		rw.Header().Add("X-Test", "test_value")
		_, _ = rw.Write([]byte(`OK`))
	}))
	defer server.Close()
	ftAPI := New(server.URL, server.Client())
	resp, err := ftAPI.PostJSON("/v1/users", testData{
		ID:   10,
		Name: "Spoody",
		Array: []string{"test_1", "test_2"},
	})
	assert.Nil(t, err)
	assert.Equal(t, "OK", getBody(resp.Body))
	assert.Equal(t, "test_value", resp.Header.Get("X-Test"))
}

func TestPatchJson(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "PATCH", req.Method)
		assert.Equal(t, "/v1/users", req.URL.String())
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t, "{\"id\":10,\"name\":\"Spoody\",\"array\":[\"test_1\",\"test_2\"]}", getBody(req.Body))
		rw.Header().Add("X-Test", "test_value")
		_, _ = rw.Write([]byte(`OK`))
	}))
	defer server.Close()
	ftAPI := New(server.URL, server.Client())
	resp, err := ftAPI.PatchJSON("/v1/users", testData{
		ID:   10,
		Name: "Spoody",
		Array: []string{"test_1", "test_2"},
	})
	assert.Nil(t, err)
	assert.Equal(t, "OK", getBody(resp.Body))
	assert.Equal(t, "test_value", resp.Header.Get("X-Test"))
}


func TestHourlyLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/v2/users", req.URL.String())
		rw.Header().Add("X-Hourly-Ratelimit-Remaining", "0")
		rw.WriteHeader(http.StatusTooManyRequests)
		_, _ = rw.Write([]byte(`Not ok`))
	}))
	defer server.Close()
	ftAPI := New(server.URL, server.Client())
	resp, err := ftAPI.Get("/v2/users")
	assert.NotNil(t, err)
	assert.Equal(t, "exceeded rate limit", err.Error())
	assert.Nil(t, resp)
}

func TestCreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method)
		assert.Equal(t, "/users", req.URL.String())
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t,
			"{\"user\":{\"login\":\"spoody\",\"email\":\"spoody@test.local\",\"first_name\":\"Spooder\",\"last_name\":\"Webz\",\"kind\":\"admin\",\"campus_id\":21}}",
			getBody(req.Body),
		)
		rw.WriteHeader(http.StatusCreated)
		_, _ = rw.Write([]byte("{\"id\": 127,\"login\":\"spoody\",\"url\": \"https://api.intra.42.fr/v2/users/spoody\"}"))
	}))
	defer server.Close()
	ftAPI := New(server.URL, server.Client())
	user := User{
		Login:     "spoody",
		Email:     "spoody@test.local",
		FirstName: "Spooder",
		LastName:  "Webz",
		Kind:      "admin",
		CampusID:  21,
	}
	expectedUser := user
	expectedUser.ID = 127
	expectedUser.URL = "https://api.intra.42.fr/v2/users/spoody"
	err := ftAPI.CreateUser(&user)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestSetUserImage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		err := req.ParseMultipartForm(10 << 20)
		assert.Nil(t, err)
		assert.Equal(t, "PATCH", req.Method)
		assert.Equal(t, "/users/spoody", req.URL.String())
		assert.Contains(t, req.Header.Get("Content-Type"), "multipart/form-data; boundary=")
		file, fileHeader, err := req.FormFile("user[image]")
		assert.Nil(t, err)
		assert.NotNil(t, file)
		assert.NotNil(t, fileHeader)
		assert.Equal(t, "profile_photo.png", fileHeader.Filename)
		assert.Equal(t, int64(99412), fileHeader.Size)
		assert.Equal(t, "form-data; name=\"user[image]\"; filename=\"profile_photo.png\"", fileHeader.Header.Get("Content-Disposition"))
		rw.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	imgFile, err := os.Open("../../tests/profile_photo.png")
	if err != nil {
		t.Fatal(err)
	}
	ftAPI := New(server.URL, server.Client())
	err = ftAPI.SetUserImage("spoody", imgFile)
	imgFile.Close()
	assert.Nil(t, err)
}

func TestSetUserImageWithNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()
	imgFile, err := os.Open("../../tests/profile_photo.png")
	if err != nil {
		t.Fatal(err)
	}
	defer imgFile.Close()
	ftAPI := New(server.URL, server.Client())
	err = ftAPI.SetUserImage("spoody", imgFile)
	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())
}

func TestSetUserImageWithFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	imgFile, err := os.Open("../../tests/profile_photo.png")
	if err != nil {
		t.Fatal(err)
	}
	defer imgFile.Close()
	ftAPI := New(server.URL, server.Client())
	err = ftAPI.SetUserImage("spoody", imgFile)
	assert.NotNil(t, err)
	assert.Equal(t, "failed setting profile image", err.Error())
}