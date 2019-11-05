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
	stash(gitR.Path)
}

func (gitR *IGitRelease) Pull() {
	pull(gitR.Path)
}

func (gitR *IGitRelease) CreateBranch() {
	createBranch(gitR.Path, fmt.Sprintf("release/%s", gitR.Release))
}

func (gitR *IGitRelease) CreateRemoteBranch() {
	createRemoteBranch(gitR.Path, fmt.Sprintf("release/%s", gitR.Release))
}

func (gitR *IGitRelease) Checkout() {
	checkout(gitR.Path, gitR.BranchRef)
}

func (gitR *IGitRelease) Commit() {
	commit(gitR.Path, fmt.Sprintf("update to release/%s", gitR.Release))
}

func (gitR *IGitRelease) Push() {
	push(gitR.Path)
}

func (gitR *IGitRelease) Add() {
	add(gitR.Path, gitR.PodspecFile)
}

func (gitR *IGitRelease) AddPods(projectName string) {
	pathPodspecFile := fmt.Sprintf("%s/%s/%s", projectName, gitR.Release, gitR.PodspecFile)
	add(gitR.Path, pathPodspecFile)
}
