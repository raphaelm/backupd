package main

import (
	"os/exec"
)

func tarExecute(j JobConfig) *exec.Cmd {
	return exec.Command("tar", "zcf", "-", j.Argument)
}
