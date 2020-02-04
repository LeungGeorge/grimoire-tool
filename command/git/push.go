package git

import (
	"log"
	"os/exec"
)

func push() error {
	cmd := exec.Command("git", "push")

	err := cmd.Run()
	if err != nil {
		log.Println("Execute Command failed:" + err.Error())
		return err
	}

	log.Println("Execute git push finished.")
	return nil
}
