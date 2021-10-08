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

type usersGetMockAPI struct {
	t *testing.T
}
func (m *usersGetMockAPI) Get(url string) (*http.Response, error) {
	return nil, nil
}
func (m *usersGetMockAPI) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersGetMockAPI) Delete(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersGetMockAPI) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersGetMockAPI) Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersGetMockAPI) PatchJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersGetMockAPI) SetUserImage(login string, img *os.File) error {
	return nil
}
func (m *usersGetMockAPI) CreateUser(user *ftapi.User, campusID int) error {
	return nil
}
func (m *usersGetMockAPI) CreateClose(close *ftapi.Close) error {
	return nil
}
func (m *usersGetMockAPI) GetUserByLogin(login string) (*ftapi.User, error) {
	campus := ftapi.Campus{
		ID:          42,
		Name:        "Los Santos",
		TimeZone:    "Africa/Casablanca",
		Language:    nil,
		UsersCount:  42,
		VogsphereID: 42,
	}
	return &ftapi.User{
		ID:             1337,
		Login:          "spoody",
		Email:          "spoody@local.test",
		FirstName:      "Spooder",
		LastName:       "Webz",
		Phone:          "+2126666666666",
		ImageURL:       "https://avatars.githubusercontent.com/u/5004111?s=460&u=a5eae6fc0bcdc8c087af14c2e476684daaeb2bbc&v=4",
		IsStaff:        true,
		Kind:           "",
		URL:            "",
		PoolMonth:      "April",
		PoolYear:       "2019",
		Campuses:       []*ftapi.Campus{&campus},
		CampusUsers:    nil,
		Roles:          nil,
		CursusUsers:    nil,
		CorrectionPoints: 4,
		Wallet: 1337,
	}, nil
}
func (m *usersGetMockAPI) UpdateUser(login string, data *ftapi.User) error  {
	return nil
}
func (m *usersGetMockAPI) AddCorrectionPoints(login string, points uint, reason string) error{
	return nil
}
func (m *usersGetMockAPI) RemoveCorrectionPoints(login string, points uint, reason string) error{
	return nil
}

func TestNewGetUserCmd(t *testing.T) {
	var api ftapi.APIInterface = &usersGetMockAPI{t: t}
	getUserCmd := NewGetUserCmd(&api)
	assert.IsType(t, &cobra.Command{}, getUserCmd)
}

func TestGetUserCmd(t *testing.T) {
	var api ftapi.APIInterface = &usersGetMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	getUserCmd := NewGetUserCmd(&api)
	getUserCmd.SetArgs([]string{"spoody"})
	getUserCmd.SetOut(stdout)
	getUserCmd.SetErr(stderr)
	err := getUserCmd.Execute()
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
	assert.Equal(t, `Id: 1337
Login: spoody
Email: spoody@local.test
First name: Spooder
Last name: Webz
Phone: +2126666666666
Image: https://avatars.githubusercontent.com/u/5004111?s=460&u=a5eae6fc0bcdc8c087af14c2e476684daaeb2bbc&v=4
Is staff: true
Correction points: 4
Wallet: 1337
Pool Month/Year: April/2019
`, string(out))
}