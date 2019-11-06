package git

import (
	"errors"
	"fmt"
	"os"
	"os/user"

	"github.com/edualb/igit/local"

	"github.com/edualb/igit/util"
	"github.com/joho/godotenv"
)

type IGitRelease struct {
	Path        string
	Release     string
	BranchRef   string
	PodspecFile string

	c local.Credentials
}

// TODO: Refactor
func (gitR *IGitRelease) SetAuth() {
	usr, err := user.Current()
	util.CheckIfError(err)
	pathFile := fmt.Sprintf("%s/iGit/username.env", usr.HomeDir)
	_, err = os.Stat(pathFile)

	if err == nil {
		err := godotenv.Load(pathFile)
		util.CheckIfError(err)
		gitR.c.Username = os.Getenv("USERNAME")
		pass, err := gitR.c.GetPassword()
		if len(pass) > 0 {
			gitR.c.Password = pass
		} else {
			util.CheckIfError(errors.New("You need to set your credentials. Run $ igit store-credentials"))
		}
	} else {
		util.CheckIfError(errors.New("You need to set your credentials. Run $ igit store-credentials"))
	}
}

func (gitR *IGitRelease) Stash() {
	stash(gitR.Path)
}

func (gitR *IGitRelease) Pull() {
	pull(gitR.Path)
}

func (gitR *IGitRelease) CreateBranch() {
	createBranch(gitR.Path, fmt.Sprintf("release/%s", gitR.Release))
}

func (gitR *IGitRelease) Checkout() {
	checkout(gitR.Path, gitR.BranchRef)
}

func (gitR *IGitRelease) Commit() {
	commit(gitR.Path, fmt.Sprintf("update to release/%s", gitR.Release))
}

func (gitR *IGitRelease) Push() {
	push(gitR.Path, gitR.c.Username, gitR.c.Password)
}

func (gitR *IGitRelease) Add() {
	add(gitR.Path, gitR.PodspecFile)
}

func (gitR *IGitRelease) AddPods(projectName string) {
	pathPodspecFile := fmt.Sprintf("%s/%s/%s", projectName, gitR.Release, gitR.PodspecFile)
	add(gitR.Path, pathPodspecFile)
}
