package util

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckIfError(err error) {
	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.Exit(1)
	}
}

func ExecCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	return err
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func InfoWarning(format string, args ...interface{}) {
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
