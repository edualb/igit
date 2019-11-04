# iGit
A CLI that automate process of git for iOS projects.

## Install
* Download `igit` binary

* Open folder `usr/local/bin` folder

* Put the binary there

* Run `$ igit -v` to test

## Commands
* `$ igit release | r`

This command prepare a release for iOS projects.

### Usage
You need to create a `.env` file with these variables:

```env
# Name of the project (get this name from Pods Repository) 
NAME=Example

# Release version
TAG_VERSION=0.0.1

# Name of .podspec file
PODSPEC_FILE=Example.podspec

# Project path
PATH_PROJECT=/Users/me/Documents/Projects/test-project-repository

# Branch which will be to push the folder in project
REFERENCE_BRANCH=develop

# Pods repository path
PODS_PATH_PROJECT=/Users/me/Documents/Projects/test-pods-repository

# Branch which will be to push the folder in Pods Repository
PODS_REFERENCE_BRANCH=master

PATH_PODSPEC=${PATH_PROJECT}/${PODSPEC_FILE}
PODS_PATH_FOLDER=${PODS_PATH_PROJECT}/${NAME}/${TAG_VERSION}/
PODS_PATH_PODSPEC=${NAME}/${TAG_VERSION}/${PODSPEC_FILE}
```

After you have created `.env` file, you can run:

```bash
$ igit r ./file.env
```
