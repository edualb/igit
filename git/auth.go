package git

import (
	"github.com/edualb/igit/util"
	"github.com/zalando/go-keyring"
)

const service = "igit"

func Store(user, pass string) {
	if len(user) > 0 && len(pass) > 0 {
		err := keyring.Set(service, user, pass)
		util.CheckIfError(err)
	}
}

func GetPassword(user string) (string, error) {
	pass, err := keyring.Get(service, user)
	if err != nil {
		return "", err
	}
	return pass, nil
}

func DeleteStore(user string) {
	if len(user) > 0 {
		err := keyring.Delete(service, user)
		util.CheckIfError(err)
	}
}
