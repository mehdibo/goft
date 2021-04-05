package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"goft/pkg/ftapi"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type usersLsMockAPI struct {
	t *testing.T
}
func (m *usersLsMockAPI) Get(url string) (*http.Response, error) {
	return nil, nil
}
func (m *usersLsMockAPI) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersLsMockAPI) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersLsMockAPI) Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersLsMockAPI) PatchJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *usersLsMockAPI) SetUserImage(login string, img *os.File) error {
	return nil
}
func (m *usersLsMockAPI) CreateUser(user *ftapi.User, campusID int) error {
	return nil
}
func (m *usersLsMockAPI) CreateClose(close *ftapi.Close) error {
	return nil
}
func (m *usersLsMockAPI) GetUserByLogin(login string) (*ftapi.User, error) {
	return nil, nil
}
func (m *usersLsMockAPI) UpdateUser(login string, data *ftapi.User) error  {
	return nil
}
func (m *usersLsMockAPI) AddCorrectionPoints(login string, points uint, reason string) error{
	return nil
}
func (m *usersLsMockAPI) RemoveCorrectionPoints(login string, points uint, reason string) error{
	return nil
}
func (m *usersLsMockAPI) GetUsers(page int, filters *map[string]string, sort *map[string]string) ([]*ftapi.User, error) {

	if page == 2 {
		return nil, nil
	}
	fmt.Println(filters)
	return []*ftapi.User{
		{
			ID:    1,
			Login: "spooda",
			URL:   "urla",
				},
		{
			ID:    2,
			Login: "norminet",
			URL:   "urlb",
				},
	}, nil
}
func (m *usersLsMockAPI) ExpandUsers(users []*ftapi.User) (int, error) {
	return 0, nil
}

func TestNewListUsersCmd(t *testing.T)  {
	var api ftapi.APIInterface = &usersLsMockAPI{t: t}
	lsUsers := NewListUsersCmd(&api)
	assert.IsType(t, &cobra.Command{}, lsUsers)
}

func TestListUsersFilters(t *testing.T)  {
	var api ftapi.APIInterface = &usersLsMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	lsUsers := NewListUsersCmd(&api)
	lsUsers.SetArgs([]string{
		"--pool-year", "2021",
		"--pool-month", "march",
		"--kind", "student",
		"--primary-campus", "21",
	})
	lsUsers.SetOut(stdout)
	lsUsers.SetErr(stderr)
	err := lsUsers.Execute()
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