package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

const confFile = "backupexporter.yml"

var job = flag.String("job", "", "Define the job to execute")

func executeModule(module, argument string) *exec.Cmd {
	switch module {
	case "tar":
		return tarExecute(argument)
	}
	log.Fatal("No module found for name ", module)
	return nil
}

func executeJob(j JobConfig) {
	c := executeModule(j.Module, j.Argument)
	gpg := exec.Command("gpg", "-e", "--recipient", config.KeyID, "--trust-model", "always")

	gpg.Stdin, _ = c.StdoutPipe()
	gpg.Stdout = os.Stdout

	gpg.Start()
	c.Run()
	gpg.Wait()
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
