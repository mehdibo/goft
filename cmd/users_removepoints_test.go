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

type removePointsMockAPI struct {
	t *testing.T
}
func (m *removePointsMockAPI) Get(url string) (*http.Response, error) {
	return nil, nil
}
func (m *removePointsMockAPI) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *removePointsMockAPI) Delete(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *removePointsMockAPI) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *removePointsMockAPI) Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *removePointsMockAPI) PatchJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *removePointsMockAPI) SetUserImage(login string, img *os.File) error {
	return nil
}
func (m *removePointsMockAPI) CreateUser(user *ftapi.User, campusID int) error {
	return nil
}
func (m *removePointsMockAPI) CreateClose(close *ftapi.Close) error {
	return nil
}
func (m *removePointsMockAPI) GetUserByLogin(login string) (*ftapi.User, error) {
	return nil, nil
}
func (m *removePointsMockAPI) UpdateUser(login string, data *ftapi.User) error  {
	return nil
}
func (m *removePointsMockAPI) AddCorrectionPoints(login string, points uint, reason string) error{
	return nil
}
func (m *removePointsMockAPI) RemoveCorrectionPoints(login string, points uint, reason string) error{
	assert.Equal(m.t, "norminet", login)
	assert.Equal(m.t, uint(2), points)
	assert.Equal(m.t, "Meow", reason)
	return nil
}

func TestNewRemovePointsCmd(t *testing.T) {
	var api ftapi.APIInterface = &removePointsMockAPI{t: t}
	removePointsCmd := NewRemovePointsCmd(&api)
	assert.IsType(t, &cobra.Command{}, removePointsCmd)
}

func TestRemovePoints(t *testing.T) {
	var api ftapi.APIInterface = &removePointsMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	removePointsCmd := NewRemovePointsCmd(&api)
	removePointsCmd.SetArgs([]string{"norminet", "2", "Meow"})
	removePointsCmd.SetOut(stdout)
	removePointsCmd.SetErr(stderr)
	err := removePointsCmd.Execute()
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

func TestRemoveNegativePoints(t *testing.T) {
	var api ftapi.APIInterface = &removePointsMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	removePointsCmd := NewRemovePointsCmd(&api)
	removePointsCmd.SetArgs([]string{"--", "spoody", "-5", "Testing purposes"})
	removePointsCmd.SetOut(stdout)
	removePointsCmd.SetErr(stderr)
	err := removePointsCmd.Execute()
	assert.NotNil(t, err)

	assert.Equal(t, "points must be greater than 0", err.Error())

	errOut, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}

	assert.Equal(t, "Error: points must be greater than 0\n", string(errOut))
}