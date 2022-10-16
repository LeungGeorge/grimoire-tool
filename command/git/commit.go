package git

import (
	"log"
	"os/exec"

	"github.com/LeungGeorge/grimoire-tool/flag"
)

func commit() error {
	msg := "added by robot"
	if flag.Message != "" {
		msg = flag.Message
	}

	cmd := exec.Command("git", "commit", "-m", msg)

	err := cmd.Run()
	if err != nil {
		log.Println("Execute git commit failed:" + err.Error())
		return err
	}

	log.Println("Execute git commit finished.")
	return nil
}
