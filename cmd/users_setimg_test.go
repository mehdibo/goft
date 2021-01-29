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

type setImgMockAPI struct {
	t *testing.T
}
func (m *setImgMockAPI) Get(url string) (*http.Response, error) {
	return nil, nil
}
func (m *setImgMockAPI) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *setImgMockAPI) PostJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *setImgMockAPI) Patch(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}
func (m *setImgMockAPI) PatchJSON(url string, data interface{}) (resp *http.Response, err error) {
	return nil, nil
}
func (m *setImgMockAPI) SetUserImage(login string, img *os.File) error {
	assert.Equal(m.t, "spoody", login)
	assert.Equal(m.t, "../tests/profile_photo.png", img.Name())
	stat, err := img.Stat()
	assert.Nil(m.t, err)
	assert.NotNil(m.t, stat)
	assert.Equal(m.t, "profile_photo.png", stat.Name())
	assert.Equal(m.t, int64(99412), stat.Size())
	return nil
}
func (m *setImgMockAPI) CreateUser(user *ftapi.User, campusID int) error {
	return nil
}
func (m *setImgMockAPI) CreateClose(close *ftapi.Close) error {
	return nil
}
func (m *setImgMockAPI) GetUserByLogin(login string) (*ftapi.User, error) {
	return nil, nil
}

func TestNewSetImgCmd(t *testing.T) {
	var api ftapi.APIInterface = &setImgMockAPI{t: t}
	setimgCmd := NewSetImgCmd(&api)
	assert.IsType(t, &cobra.Command{}, setimgCmd)
}

func TestSetUserImg(t *testing.T) {
	var api ftapi.APIInterface = &setImgMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	setimgCmd := NewSetImgCmd(&api)
	setimgCmd.SetArgs([]string{"spoody", "../tests/profile_photo.png"})
	setimgCmd.SetOut(stdout)
	setimgCmd.SetErr(stderr)
	err := setimgCmd.Execute()
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

func TestSetUserImgWrongFile(t *testing.T) {
	var api ftapi.APIInterface = &setImgMockAPI{t: t}
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")

	setimgCmd := NewSetImgCmd(&api)
	setimgCmd.SetArgs([]string{"spoody", "/inexisting_dir/inexisting_file.png"})
	setimgCmd.SetOut(stdout)
	setimgCmd.SetErr(stderr)
	err := setimgCmd.Execute()
	assert.NotNil(t, err)

	out, readErr := ioutil.ReadAll(stdout)
	if readErr != nil {
		t.Fatal(readErr)
	}

	errOut, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}

	assert.Equal(t, "Error: stat /inexisting_dir/inexisting_file.png: no such file or directory\n", string(errOut))
	assert.Equal(t, "stat /inexisting_dir/inexisting_file.png: no such file or directory", err.Error())
	assert.Equal(t, "Usage:\n  setimg login image_path [flags]\n\nFlags:\n  -h, --help   help for setimg\n\n", string(out))
}

func generateFile(size int64, fileName string) (*os.File, error) {
	content := make([]byte, size)
	tmpFile, err := ioutil.TempFile("/tmp", fileName)
	if err != nil {
		return nil, err
	}
	if _, err := tmpFile.Write(content); err != nil {
		return nil, err
	}
	return tmpFile, nil
}

func TestCreateUserWithSmallFile(t *testing.T) {
	var api ftapi.APIInterface = &mockAPI{t: t}
	stderr := bytes.NewBufferString("")
	stdout := bytes.NewBufferString("")

	tmpFile, err := generateFile(10, "small.png")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())


	setimgCmd := NewSetImgCmd(&api)
	setimgCmd.SetArgs([]string{"spoody", tmpFile.Name()})
	setimgCmd.SetOut(stdout)
	setimgCmd.SetErr(stderr)
	err = setimgCmd.Execute()
	assert.NotNil(t, err)

	out, readErr := ioutil.ReadAll(stdout)
	if readErr != nil {
		t.Fatal(readErr)
	}

	errOut, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}

	assert.Equal(t, "Error: image file size is invalid\n", string(errOut))
	assert.Equal(t, "image file size is invalid", err.Error())
	assert.Equal(t, "Usage:\n  setimg login image_path [flags]\n\nFlags:\n  -h, --help   help for setimg\n\n", string(out))
}

func TestCreateUserWithBigFile(t *testing.T) {
	var api ftapi.APIInterface = &mockAPI{t: t}
	stderr := bytes.NewBufferString("")
	stdout := bytes.NewBufferString("")

	tmpFile, err := generateFile(1050000, "big.png")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	setimgCmd := NewSetImgCmd(&api)
	setimgCmd.SetArgs([]string{"spoody", tmpFile.Name()})
	setimgCmd.SetOut(stdout)
	setimgCmd.SetErr(stderr)
	err = setimgCmd.Execute()
	assert.NotNil(t, err)

	out, readErr := ioutil.ReadAll(stdout)
	if readErr != nil {
		t.Fatal(readErr)
	}

	errOut, readErr := ioutil.ReadAll(stderr)
	if readErr != nil {
		t.Fatal(readErr)
	}

	assert.Equal(t, "Error: image file size is invalid\n", string(errOut))
	assert.Equal(t, "image file size is invalid", err.Error())
	assert.Equal(t, "Usage:\n  setimg login image_path [flags]\n\nFlags:\n  -h, --help   help for setimg\n\n", string(out))
}