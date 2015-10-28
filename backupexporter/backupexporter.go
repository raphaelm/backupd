package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

const confFile = "backupexporter.yml"

var job = flag.String("job", "", "Define the job to execute")

func executeModule(j JobConfig) (*exec.Cmd, string) {
	switch j.Module {
	case "tar":
		return tarExecute(j)
	}
	log.Fatal("No module found for name ", j.Module)
	return nil, ""
}

func executeJob(j JobConfig) {
	c, fname := executeModule(j)
	gpg := exec.Command("gpg", "-e", "--recipient", config.KeyID, "--trust-model", "always")
	hash := sha256.New()

	gpg.Stdin, _ = c.StdoutPipe()
	gpg.Stdout = io.MultiWriter(os.Stdout, hash)
	gpg.Stderr = os.Stderr
	c.Stderr = c.Stderr

	gpg.Start()
	if err := c.Run(); err != nil {
		log.Fatal(err)
		return
	}
	if err := gpg.Wait(); err != nil {
		log.Fatal(err)
		gpg.Wait()
		return
	}

	// Write filename and checksum to stderr, as stdout is meant to be piped into a file
	fmt.Fprintln(os.Stderr, "name:", fname)
	fmt.Fprintln(os.Stderr, "checksum:", hex.EncodeToString(hash.Sum(nil)))
}

func main() {
	flag.Parse()

	readConfig()

	if *job == "" {
		log.Fatal("Usage: -job JOBNAME")
	}

	for _, j := range config.Jobs {
		if j.Name == *job {
			executeJob(j)
			return
		}
	}
	log.Fatal("No job configuration found for name ", *job)
}
