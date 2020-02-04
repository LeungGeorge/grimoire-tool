package git

import (
	"log"
	"os/exec"

	"github.com/LeungGeorge/grit/flag"
)

func commit() error {
	msg := "auto commit"
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
