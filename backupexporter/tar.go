package main

import (
	"os/exec"
)

func tarExecute(j JobConfig) (*exec.Cmd, string) {
	return exec.Command("tar", "zcf", "-", j.Argument), j.Name + ".tar.gz"
}
