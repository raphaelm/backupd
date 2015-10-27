package main

import (
	"os/exec"
)

func tarExecute(argument string) *exec.Cmd {
	return exec.Command("tar", "zcf", "-", argument)
}
