package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/edualb/igit/util"
)

// TODO: Send to IGitRelease
func SetPodfileVersion(path, version string) {
	input, err := ioutil.ReadFile(path)
	util.CheckIfError(err)

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "s.version") {
			util.Info(fmt.Sprintf("Changed from '%s' to '    s.version          = \"%s\"'", lines[i], version))
			lines[i] = fmt.Sprintf("    s.version          = \"%s\"", version)
			break
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	util.CheckIfError(err)
}

func Copy(from, to string) {
	util.Info(fmt.Sprintf("Copying from %s to %s", from, to))
	util.ExecCommand(fmt.Sprintf("rsync %s %s", from, to))
}

func Mkdir(path string) {
	util.Info(fmt.Sprintf("mkdir %s", path))
	err := os.Mkdir(path, 0777)
	if err == nil || os.IsExist(err) {
		return
	} else {
		util.CheckIfError(err)
	}
}

// SOURCE: https://gist.github.com/novalagung/13c5c8f4d30e0c4bff27
func CreateFile(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		util.CheckIfErrorDefault(err)
		defer file.Close()
	}
}
