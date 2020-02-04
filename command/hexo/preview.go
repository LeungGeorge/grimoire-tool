package hexo

import (
	"os/exec"

	"github.com/urfave/cli"
)

// preview executes hexo_preview command and return exit code.
func preview(ctx *cli.Context) error {
	c := `
#!/bin/bash
hexo clean
hexo g
hexo s --debug
	`

	cmd := exec.Command("bash", "-c", c)
	stdOut, stdOutErr := cmd.StdoutPipe()
	if stdOutErr != nil {
		panic(stdOutErr)
	}

	stdErr, stdErrErr := cmd.StderrPipe()
	if stdErrErr != nil {
		panic(stdErrErr)
	}

	go syncLog(stdOut)
	go syncLog(stdErr)

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	return nil
}
