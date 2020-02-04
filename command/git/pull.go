package git

import (
	"log"
	"os/exec"
)

func pull() error {
	cmd := exec.Command("git", "pull")

	err := cmd.Run()
	if err != nil {
		log.Println("Execute Command failed:" + err.Error())
		return err
	}

	log.Println("Execute git pull finished.")
	return nil
}
