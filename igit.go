package main

import (
	"fmt"
	"os"

	"github.com/edualb/igit/git"
	"github.com/edualb/igit/service"
	"github.com/edualb/igit/util"
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
			Name:        "release",
			Aliases:     []string{"r"},
			Description: "Create a release in a project and create a folder for this release in a pods repository",
			Action: func(c *cli.Context) {
				if c.NArg() > 0 {
					err := godotenv.Load(c.Args()[0])
					util.CheckIfError(err)
				} else {
					util.InfoWarning("You need to send the path of .env file in argument.")
					return
				}
				release()
			},
		},
		{
			Name:        "store-credential",
			Aliases:     []string{"sc"},
			Description: "Store credentials from git",
			Action: func(c *cli.Context) {
				user := ""
				pass := ""
				if c.NArg() > 1 {
					user = c.Args()[0]
					pass = c.Args()[1]
				} else {
					util.InfoWarning("You need to set your username and password.")
					return
				}
				git.Store(user, pass)
				password, err := git.GetPassword(user)
				util.CheckIfError(err)
				fmt.Println(fmt.Sprintf("password: %s", password))
			},
		},
	}
}

func release() {
	mainProject := &git.IGitRelease{
		Path:        os.Getenv("PATH_PROJECT"),
		Release:     os.Getenv("TAG_VERSION"),
		BranchRef:   os.Getenv("REFERENCE_BRANCH"),
		PodspecFile: os.Getenv("PODSPEC_FILE"),
	}

	util.Info(fmt.Sprintf("* %s: ", os.Getenv("PROJECT_NAME")))
	mainProject.Stash()
	mainProject.Checkout()
	mainProject.Pull()
	mainProject.CreateBranch()
	mainProject.CreateRemoteBranch()
	service.SetPodfileVersion(os.Getenv("PATH_PODSPEC"), mainProject.Release)
	mainProject.Add()
	mainProject.Commit()
	mainProject.Push()

	podsProject := &git.IGitRelease{
		Path:        os.Getenv("PODS_PATH_PROJECT"),
		Release:     os.Getenv("TAG_VERSION"),
		BranchRef:   os.Getenv("PODS_REFERENCE_BRANCH"),
		PodspecFile: os.Getenv("PODSPEC_FILE"),
	}

	util.Info("* PODS REPOSITORY:")
	podsProject.Stash()
	podsProject.Checkout()
	podsProject.Pull()
	service.Mkdir(os.Getenv("PODS_PATH_FOLDER"))
	service.Copy(os.Getenv("PATH_PODSPEC"), os.Getenv("PODS_PATH_FOLDER"))
	podsProject.AddPods(os.Getenv("PROJECT_NAME"))
	podsProject.Commit()
	podsProject.Push()
}
