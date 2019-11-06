package git

import (
	"fmt"
	"os"
	"time"

	"github.com/edualb/igit/util"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// TODO: Remove this function when it is not more necessary
func stash(path string) {
	util.Info("git stash")
	err := util.ExecCommand(fmt.Sprintf("cd %s; git stash;", path))
	util.CheckIfError(err)
}

// TODO: Get iGit path instead of set path in parameters
func clone(url, path, username, password string) {
	util.Info("git clone %s %s", url, path)

	_, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil && err != git.ErrRepositoryAlreadyExists {
		util.CheckIfError(err)
	}
}

func pull(path, username, password string) {
	util.Info("git pull")

	w := getWorktree(path)

	err := w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		util.CheckIfError(err)
	}
}

func push(path, username, password string) {
	util.Info("git push")

	r := getRepository(path)

	err := r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
	})
	util.CheckIfError(err)
}

func checkout(path, branch string) {
	worktree := getWorktree(path)

	util.Info(fmt.Sprintf("git checkout %s", branch))

	err := worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	})
	util.CheckIfError(err)
}

func createBranch(path, branch string) {
	r := getRepository(path)

	headRef, err := r.Head()
	util.CheckIfError(err)

	branchName := plumbing.NewBranchReferenceName(branch)
	worktree, err := r.Worktree()
	util.CheckIfError(err)

	util.Info(fmt.Sprintf("git checkout -b %s", branch))

	err = worktree.Checkout(&git.CheckoutOptions{
		Hash:   headRef.Hash(),
		Branch: branchName,
		Create: true,
	})
	util.CheckIfError(err)
}

func add(path, file string) {
	worktree := getWorktree(path)
	util.Info(fmt.Sprintf("git add %s", file))
	_, err := worktree.Add(file)
	util.CheckIfError(err)
}

func commit(path, message string) {
	worktree := getWorktree(path)
	util.Info(fmt.Sprintf("git commit -m  \"%s\"", message))
	_, err := worktree.Commit(message, &git.CommitOptions{
		All:    true,
		Author: &object.Signature{Name: "iGit", When: time.Now()},
	})
	util.CheckIfError(err)
}

func getRepository(path string) *git.Repository {
	r, err := git.PlainOpen(path)
	util.CheckIfError(err)
	return r
}

func getWorktree(path string) *git.Worktree {
	r := getRepository(path)
	w, err := r.Worktree()
	util.CheckIfError(err)
	return w
}
