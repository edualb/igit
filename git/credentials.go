package git

import (
	"os"

	"github.com/edualb/igit/util"
	"github.com/zalando/go-keyring"
)

type Credentials struct {
	Username string
	Password string
}

const service = "igit"

func (c *Credentials) Store() {
	err := keyring.Set(service, c.Username, c.Password)
	util.CheckIfError(err)
}

func (c *Credentials) GetPassword() (string, error) {
	pass, err := keyring.Get(service, c.Username)
	if err != nil {
		return "", err
	}
	return pass, nil
}

func (c *Credentials) DeleteStore() {
	err := keyring.Delete(service, c.Username)
	util.CheckIfError(err)
}

func (c *Credentials) localstore() {
	// service.Mkdir("/usr/local/iGit/")
	// service.CreateFile("/usr/local/iGit/.username")
	writeUserCredential("/usr/local/iGit/.username", c.Username)
}

func writeUserCredential(path, user string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0777)
	util.CheckIfErrorDefault(err)
	defer file.Close()

	_, err = file.WriteString(user)
	util.CheckIfErrorDefault(err)

	err = file.Sync()
	util.CheckIfErrorDefault(err)
}
