package git

import (
	"fmt"
)

type IGitRelease struct {
	Path        string
	Release     string
	BranchRef   string
	PodspecFile string
}

func (gitR *IGitRelease) Stash() {
	Stash(gitR.Path)
}

func (gitR *IGitRelease) Pull() {
	Pull(gitR.Path)
}

func (gitR *IGitRelease) CreateBranch() {
	CreateBranch(gitR.Path, fmt.Sprintf("release/%s", gitR.Release))
}

func (gitR *IGitRelease) CreateRemoteBranch() {
	CreateRemoteBranch(gitR.Path, fmt.Sprintf("release/%s", gitR.Release))
}

func (gitR *IGitRelease) Checkout() {
	Checkout(gitR.Path, gitR.BranchRef)
}

func (gitR *IGitRelease) Commit() {
	Commit(gitR.Path, fmt.Sprintf("update to release/%s", gitR.Release))
}

func (gitR *IGitRelease) Push() {
	Push(gitR.Path)
}

func (gitR *IGitRelease) Add() {
	Add(gitR.Path, gitR.PodspecFile)
}
