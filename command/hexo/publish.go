package hexo

import (
	"os/exec"

	"github.com/urfave/cli"
)

// publish executes hexo_publish command and return exit code.
func publish(ctx *cli.Context) error {
	c := `
#!/bin/bash

hexo clean
hexo g
hexo d

hexo algolia

git add --all
git commit -m "auto commit"
git pull
git push origin hexo
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
