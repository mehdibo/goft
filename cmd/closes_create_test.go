package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"goft/pkg/ftapi"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type createCloseMockAPI struct {
	t *testing.T
}
func (m *createCloseMockAPI) Get(url string) (*http.Response, error) {
	return nil, nil
}
func (m *createCloseMockAPI) Delete(url string, contentType string, body io.Reader) (*http.Response, error) {
	return nil, nil
}
func (m *createCloseMockAPI) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *createCloseMockAPI) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *createCloseMockAPI) Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *createCloseMockAPI) PatchJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *createCloseMockAPI) SetUserImage(login string, img *os.File) error {
	return nil
}
func (m *createCloseMockAPI) CreateUser(user *ftapi.User, campusID int) error {
	return nil
}
func (m *createCloseMockAPI) CreateClose(close *ftapi.Close) error {
	assert.Equal(m.t, "other", close.Kind)
	assert.Equal(m.t, "Testing purposes", close.Reason)
	assert.Equal(m.t, "spoody", close.User.Login)
	assert.Equal(m.t, 42, close.Closer.ID)
	return nil
}
func (m *createCloseMockAPI) GetUserByLogin(login string) (*ftapi.User, error) {
	assert.Equal(m.t, "spoody", login)
	return &ftapi.User{
		ID: 42,
	}, nil
}
func (m *createCloseMockAPI) UpdateUser(login string, data *ftapi.User) error  {
	return nil
}
func (m *createCloseMockAPI) AddCorrectionPoints(login string, points uint, reason string) error{
	return nil
}
func (m *createCloseMockAPI) RemoveCorrectionPoints(login string, points uint, reason string) error{
	return nil
}
func (m *createCloseMockAPI) GetUserAgus(login string) ([]ftapi.Agu, error) {
	return nil, nil
}
func (m *createCloseMockAPI) CreateFreePastAgu(login string, duration int, reason string) error {
	return nil
}

func TestNewCloseCreateCmd(t *testing.T) {
	var api ftapi.APIInterface = &createCloseMockAPI{t: t}
	createCloseCmd := NewCloseCreateCmd(&api)
	assert.IsType(t, &cobra.Command{}, createCloseCmd)
}

func TestCreateClose(t *testing.T) {
	var api ftapi.APIInterface = &createCloseMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	testCmd := NewCloseCreateCmd(&api)

	testCmd.SetArgs([]string{"spoody", "other", "Testing purposes", "spoody"})
	testCmd.SetOut(stdout)
	testCmd.SetErr(stderr)
	err := testCmd.Execute()
	assert.Nil(t, err)

	out, readErr := ioutil.ReadAll(stdout)
	if readErr != nil {
		t.Fatal(readErr)
	}

	errOut, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}

	assert.Equal(t, "", string(errOut))
	assert.Equal(t, "", string(out))
}

func TestCreateCloseInvalid(t *testing.T) {
	var api ftapi.APIInterface = &createCloseMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	testCmd := NewCloseCreateCmd(&api)

	testCmd.SetArgs([]string{"spoody", "invalid_kind", "Testing purposes", "42"})
	testCmd.SetOut(stdout)
	testCmd.SetErr(stderr)
	err := testCmd.Execute()
	assert.NotNil(t, err)
	assert.Equal(t, "'invalid_kind' is not a valid kind", err.Error())

	errOut, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}

	assert.Equal(t, "Error: 'invalid_kind' is not a valid kind\n", string(errOut))
}