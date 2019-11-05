package git

import (
	"fmt"
	"time"

	"github.com/edualb/igit/util"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// TODO: Replace ExecCommand to go-git.v4 lib
func stash(path string) {
	util.Info("git stash")
	err := util.ExecCommand(fmt.Sprintf("cd %s; git stash;", path))
	util.CheckIfError(err)
}

// TODO: Replace ExecCommand to go-git.v4 lib
func pull(path string) {
	util.Info("git pull")
	err := util.ExecCommand(fmt.Sprintf("cd %s; git pull;", path))
	util.CheckIfError(err)
}

// TODO: Replace ExecCommand to go-git.v4 lib
func createRemoteBranch(path string, branch string) {
	util.Info(fmt.Sprintf("git push --set-upstream origin %s", branch))
	err := util.ExecCommand(fmt.Sprintf("cd %s; git push --set-upstream origin %s", path, branch))
	util.CheckIfError(err)
}

// TODO: Replace ExecCommand to go-git.v4 lib
func push(path string) {
	util.Info("git push")
	err := util.ExecCommand(fmt.Sprintf("cd %s; git push", path))
	util.CheckIfError(err)
}

func checkout(path string, branch string) {
	worktree := getWorktree(path)

	util.Info(fmt.Sprintf("git checkout %s", branch))

	err := worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	})
	util.CheckIfError(err)
}

func createBranch(path string, branch string) {
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

func add(path string, file string) {
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
