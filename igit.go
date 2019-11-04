package main

import (
	"fmt"
	"os"

	"github.com/edualb/go-generate-tag-ios-globo/git"
	"github.com/edualb/go-generate-tag-ios-globo/service"
	"github.com/edualb/go-generate-tag-ios-globo/util"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	util.CheckIfError(err)
}

func info() {
	app.Name = "iGit"
	app.Usage = "A CLI to automate the git process for iOS projects"
	app.Author = "Eduardo Albuquerque da Silva"
	app.Version = "0.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "release",
			Aliases: []string{"r"},
			Action: func(c *cli.Context) {
				if c.NArg() > 0 {
					err := godotenv.Load(c.Args()[0])
					util.CheckIfError(err)
				} else {
					util.InfoWarning("You need send the path of .env file in argument.")
					return
				}
				release()
			},
		},
	}
}

func release() {
	util.Info(fmt.Sprintf("************************** %s **************************", os.Getenv("NAME")))
	git.Stash(os.Getenv("PATH_PROJECT"))
	git.Checkout(os.Getenv("PATH_PROJECT"), os.Getenv("REFERENCE_BRANCH"))
	git.Pull(os.Getenv("PATH_PROJECT"))
	git.CreateBranch(os.Getenv("PATH_PROJECT"), os.Getenv("TAG_VERSION"))
	git.CreateRemoteBranch(os.Getenv("PATH_PROJECT"), os.Getenv("TAG_VERSION"))
	service.SetPodfileVersion(os.Getenv("PATH_PODSPEC"), os.Getenv("TAG_VERSION"))
	git.Add(os.Getenv("PATH_PROJECT"), os.Getenv("PODSPEC_FILE"))
	git.Commit(os.Getenv("PATH_PROJECT"), os.Getenv("TAG_VERSION"))
	git.Push(os.Getenv("PATH_PROJECT"))

	util.Info("************************** PODS REPOSITORY **************************")
	git.Stash(os.Getenv("PODS_PATH_PROJECT"))
	git.Checkout(os.Getenv("PODS_PATH_PROJECT"), os.Getenv("PODS_REFERENCE_BRANCH"))
	git.Pull(os.Getenv("PODS_PATH_PROJECT"))
	service.Mkdir(os.Getenv("PODS_PATH_FOLDER"))
	service.Copy(os.Getenv("PATH_PODSPEC"), os.Getenv("PODS_PATH_FOLDER"))
	git.Add(os.Getenv("PODS_PATH_PROJECT"), os.Getenv("PODS_PATH_PODSPEC"))
	git.Commit(os.Getenv("PODS_PATH_PROJECT"), os.Getenv("TAG_VERSION"))
	git.Push(os.Getenv("PODS_PATH_PROJECT"))
}
