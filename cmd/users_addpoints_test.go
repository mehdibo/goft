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

type addPointsMockAPI struct {
	t *testing.T
}
func (m *addPointsMockAPI) Get(url string) (*http.Response, error) {
	return nil, nil
}
func (m *addPointsMockAPI) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *addPointsMockAPI) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *addPointsMockAPI) Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *addPointsMockAPI) PatchJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *addPointsMockAPI) SetUserImage(login string, img *os.File) error {
	return nil
}
func (m *addPointsMockAPI) CreateUser(user *ftapi.User, campusID int) error {
	return nil
}
func (m *addPointsMockAPI) CreateClose(close *ftapi.Close) error {
	return nil
}
func (m *addPointsMockAPI) GetUserByLogin(login string) (*ftapi.User, error) {
	return nil, nil
}
func (m *addPointsMockAPI) UpdateUser(login string, data *ftapi.User) error  {
	return nil
}
func (m *addPointsMockAPI) AddCorrectionPoints(login string, points uint, reason string) error{
	assert.Equal(m.t, "spoody", login)
	assert.Equal(m.t, uint(5), points)
	assert.Equal(m.t, "Testing purposes", reason)
	return nil
}
func (m *addPointsMockAPI) RemoveCorrectionPoints(login string, points uint, reason string) error{
	return nil
}

func TestNewAddPointsCmd(t *testing.T) {
	var api ftapi.APIInterface = &addPointsMockAPI{t: t}
	addPointsCmd := NewAddPointsCmd(&api)
	assert.IsType(t, &cobra.Command{}, addPointsCmd)
}

func TestAddPoints(t *testing.T) {
	var api ftapi.APIInterface = &addPointsMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	addPointsCmd := NewAddPointsCmd(&api)
	addPointsCmd.SetArgs([]string{"spoody", "5", "Testing purposes"})
	addPointsCmd.SetOut(stdout)
	addPointsCmd.SetErr(stderr)
	err := addPointsCmd.Execute()
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

func TestAddNegativePoints(t *testing.T) {
	var api ftapi.APIInterface = &addPointsMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	addPointsCmd := NewAddPointsCmd(&api)
	addPointsCmd.SetArgs([]string{"--", "spoody", "-5", "Testing purposes"})
	addPointsCmd.SetOut(stdout)
	addPointsCmd.SetErr(stderr)
	err := addPointsCmd.Execute()
	assert.NotNil(t, err)

	assert.Equal(t, "points must be greater than 0", err.Error())

	errOut, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}

	assert.Equal(t, "Error: points must be greater than 0\n", string(errOut))
}