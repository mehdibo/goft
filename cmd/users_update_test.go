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
	"strings"
	"testing"
)

type updateUserMockAPI struct {
	t *testing.T
}
func (m *updateUserMockAPI) Get(url string) (*http.Response, error) {
	return nil, nil
}
func (m *updateUserMockAPI) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *updateUserMockAPI) Delete(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *updateUserMockAPI) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *updateUserMockAPI) Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *updateUserMockAPI) PatchJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *updateUserMockAPI) SetUserImage(login string, img *os.File) error {
	return nil
}
func (m *updateUserMockAPI) CreateUser(user *ftapi.User, campusID int) error {
	return nil
}
func (m *updateUserMockAPI) CreateClose(close *ftapi.Close) error {
	return nil
}
func (m *updateUserMockAPI) GetUserByLogin(login string) (*ftapi.User, error) {
	return nil, nil
}
func (m *updateUserMockAPI) UpdateUser(login string, data *ftapi.User) error  {
	assert.Equal(m.t, "spoody", login)
	assert.Equal(m.t, "Spooder", data.FirstName)
	assert.Equal(m.t, "Webz", data.LastName)
	assert.Equal(m.t, "external", data.Kind)
	assert.Equal(m.t, "new_password", data.Password)
	return nil
}
func (m *updateUserMockAPI) AddCorrectionPoints(login string, points uint, reason string) error{
	return nil
}
func (m *updateUserMockAPI) RemoveCorrectionPoints(login string, points uint, reason string) error{
	return nil
}

func TestNewUpdateUserCmd(t *testing.T) {
	var api ftapi.APIInterface = &updateUserMockAPI{t: t}
	updateUserCmd := NewUpdateUserCmd(&api)
	assert.IsType(t, &cobra.Command{}, updateUserCmd)
}

func TestUpdateUser(t *testing.T) {
	var api ftapi.APIInterface = &updateUserMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	updateUserCmd := NewUpdateUserCmd(&api)
	updateUserCmd.SetArgs([]string{"spoody", "--first-name", "Spooder", "--last-name", "Webz", "--kind", "external", "--password"})
	// \n at the end to simulate pressing enter and making sure it is not passed with the password
	updateUserCmd.SetIn(strings.NewReader("new_password\n"))
	updateUserCmd.SetOut(stdout)
	updateUserCmd.SetErr(stderr)
	err := updateUserCmd.Execute()
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
	assert.Equal(t, "Enter new password: User updated\n", string(out))
}

func TestUpdateUserWithInvalidArgs(t *testing.T) {
	var api ftapi.APIInterface = &updateUserMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	updateUserCmd := NewUpdateUserCmd(&api)
	// \n at the end to simulate pressing enter and making sure it is not passed with the password
	updateUserCmd.SetOut(stdout)
	updateUserCmd.SetErr(stderr)
	err := updateUserCmd.Execute()
	assert.NotNil(t, err)


	errOut, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}

	assert.Equal(t, "accepts 1 arg(s), received 0", err.Error())
	assert.Equal(t, "Error: accepts 1 arg(s), received 0\n", string(errOut))
}