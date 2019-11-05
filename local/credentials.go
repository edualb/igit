package git

import (
	"fmt"
	"os"
	"os/user"

	"github.com/edualb/igit/service"
	"github.com/edualb/igit/util"
	"github.com/zalando/go-keyring"
)

type Credentials struct {
	Username string
	Password string
}

const serviceIGit = "igit"

func (c *Credentials) Store() {
	err := keyring.Set(serviceIGit, c.Username, c.Password)
	util.CheckIfError(err)
	c.localstore()
}

func (c *Credentials) GetPassword() (string, error) {
	pass, err := keyring.Get(serviceIGit, c.Username)
	if err != nil {
		return "", err
	}
	return pass, nil
}

func (c *Credentials) DeleteStore() {
	err := keyring.Delete(serviceIGit, c.Username)
	util.CheckIfError(err)
}

// TODO Refactor
func (c *Credentials) localstore() {
	usr, err := user.Current()
	util.CheckIfError(err)
	path := fmt.Sprintf("%s/iGit/", usr.HomeDir)
	file := fmt.Sprintf("%s/username.env", path)
	service.Mkdir(path)
	service.CreateFile(file)
	c.writeUserCredential(file)
}

func (c *Credentials) writeUserCredential(path string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0777)
	util.CheckIfErrorDefault(err)
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("USERNAME=%s", c.Username))
	util.CheckIfErrorDefault(err)

	err = file.Sync()
	util.CheckIfErrorDefault(err)
}
