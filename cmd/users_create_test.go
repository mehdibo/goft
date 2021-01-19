package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"goft/pkg/ftapi"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockAPI struct {
	t *testing.T
}

func (m *mockAPI) Get(url string) (*http.Response, error) {
	return nil, nil
}

func (m *mockAPI) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}

func (m *mockAPI) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}

func (m *mockAPI) CreateUser(user *ftapi.User) error {
	if user.Email == "spoody@with.login" {
		assert.Equal(m.t, "spoody", user.Login)
		assert.Equal(m.t, "Mehdi", user.FirstName)
		assert.Equal(m.t, "Bounya", user.LastName)
		assert.Equal(m.t, "admin", user.Kind)
		assert.Equal(m.t, 21, user.CampusID)
	} else if user.Email == "spoody@without.login" {
		assert.Equal(m.t, "", user.Login)
		assert.Equal(m.t, "Mehdi", user.FirstName)
		assert.Equal(m.t, "Bounya", user.LastName)
		assert.Equal(m.t, "admin", user.Kind)
		assert.Equal(m.t, 21, user.CampusID)
	} else {
		m.t.Fatal("Reached else statement in CreateUser")
	}

	return nil
}

func TestCreateUserWithLogin(t *testing.T) {
	var api ftapi.APIInterface = &mockAPI{t: t}
	stdout := bytes.NewBufferString("")
	createCmd := NewUserCreateCmd(&api)
	createCmd.SetArgs([]string{"--login", "spoody", "spoody@with.login", "Mehdi", "Bounya", "admin", "21"})
	createCmd.SetOut(stdout)
	err := createCmd.Execute()
	out, readErr := ioutil.ReadAll(stdout)
	if readErr != nil {
		t.Fatal(readErr)
	}
	assert.Nil(t, err)
	assert.Equal(t, "User created\n", string(out))
}

func TestCreateUserWithoutLogin(t *testing.T) {
	var api ftapi.APIInterface = &mockAPI{t: t}
	stdout := bytes.NewBufferString("")
	createCmd := NewUserCreateCmd(&api)
	createCmd.SetArgs([]string{"spoody@without.login", "Mehdi", "Bounya", "admin", "21"})
	createCmd.SetOut(stdout)
	err := createCmd.Execute()
	out, readErr := ioutil.ReadAll(stdout)
	if readErr != nil {
		t.Fatal(readErr)
	}
	assert.Nil(t, err)
	assert.Equal(t, "User created\n", string(out))
}

func TestCreateUserWithInvalidArgs(t *testing.T) {
	var api ftapi.APIInterface = &mockAPI{t: t}
	stderr := bytes.NewBufferString("")
	stdout := bytes.NewBufferString("")
	createCmd := NewUserCreateCmd(&api)
	createCmd.SetOut(stdout) // To prevent cmd from writing to stdout in tests
	createCmd.SetErr(stderr)
	err := createCmd.Execute()
	stderrText, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}
	assert.NotNil(t, err)
	assert.Equal(t, "accepts 5 arg(s), received 0", err.Error())
	assert.Equal(t, "Error: accepts 5 arg(s), received 0\n", string(stderrText))
}

func TestCreateUserWithInvalidKind(t *testing.T) {
	var api ftapi.APIInterface = &mockAPI{t: t}
	stderr := bytes.NewBufferString("")
	stdout := bytes.NewBufferString("")
	createCmd := NewUserCreateCmd(&api)
	createCmd.SetArgs([]string{"spoody@without.login", "Mehdi", "Bounya", "invalid_kind", "21"})
	createCmd.SetOut(stdout) // To prevent cmd from writing to stdout in tests
	createCmd.SetErr(stderr)
	err := createCmd.Execute()
	stderrText, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}
	assert.NotNil(t, err)
	assert.Equal(t, "kind must be admin, student or external", err.Error())
	assert.Equal(t, "Error: kind must be admin, student or external\n", string(stderrText))
}

func TestCreateUserWithInvalidCampusId(t *testing.T) {
	var api ftapi.APIInterface = &mockAPI{t: t}
	stderr := bytes.NewBufferString("")
	stdout := bytes.NewBufferString("")
	createCmd := NewUserCreateCmd(&api)
	createCmd.SetArgs([]string{"spoody@without.login", "Mehdi", "Bounya", "admin", "invalid_campus"})
	createCmd.SetOut(stdout) // To prevent cmd from writing to stdout in tests
	createCmd.SetErr(stderr)
	err := createCmd.Execute()
	stderrText, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}
	assert.NotNil(t, err)
	assert.Equal(t, "invalid campus_id", err.Error())
	assert.Equal(t, "Error: invalid campus_id\n", string(stderrText))
}