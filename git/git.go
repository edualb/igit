package git

import (
	"fmt"
	"time"

	"github.com/edualb/go-generate-tag-ios-globo/util"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func Stash(path string) {
	util.Info("git stash")
	err := util.ExecCommand(fmt.Sprintf("cd %s; git stash;", path))
	util.CheckIfError(err)
}

func Pull(path string) {
	util.Info("git pull")
	err := util.ExecCommand(fmt.Sprintf("cd %s; git pull;", path))
	util.CheckIfError(err)
}

func Checkout(path string, branch string) {
	worktree := getWorktree(path)

	util.Info(fmt.Sprintf("git checkout %s", branch))

	err := worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	})
	util.CheckIfError(err)
}

func CreateBranch(path string, version string) {
	r := getRepository(path)

	headRef, err := r.Head()
	util.CheckIfError(err)

	branchName := plumbing.NewBranchReferenceName(fmt.Sprintf("release/%s", version))
	worktree, err := r.Worktree()
	util.CheckIfError(err)

	util.Info(fmt.Sprintf("git checkout -b release/%s", version))

	err = worktree.Checkout(&git.CheckoutOptions{
		Hash:   headRef.Hash(),
		Branch: branchName,
		Create: true,
	})
	util.CheckIfError(err)
}

func CreateRemoteBranch(path string, version string) {
	util.Info(fmt.Sprintf("git push --set-upstream origin release/%s", version))
	err := util.ExecCommand(fmt.Sprintf("cd %s; git push --set-upstream origin release/%s", path, version))
	util.CheckIfError(err)
}

func Add(path string, file string) {
	worktree := getWorktree(path)
	util.Info(fmt.Sprintf("git add %s", file))
	_, err := worktree.Add(file)
	util.CheckIfError(err)
}

func Commit(path string, version string) {
	worktree := getWorktree(path)
	util.Info(fmt.Sprintf("git commit -m  \"update to release/%s\"", version))
	_, err := worktree.Commit(fmt.Sprintf("update to release/%s", version), &git.CommitOptions{
		All:    true,
		Author: &object.Signature{Name: "iGit", When: time.Now()},
	})
	util.CheckIfError(err)
}

func Push(path string) {
	util.Info("git push")
	err := util.ExecCommand(fmt.Sprintf("cd %s; git push", path))
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
