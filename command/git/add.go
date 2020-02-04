package git

import (
	"log"
	"os/exec"
)

func add() error {
	cmd := exec.Command("git", "add", ".")

	err := cmd.Run()
	if err != nil {
		log.Println("Execute Command failed:" + err.Error())
		return err
	}

	log.Println("Execute git add finished.")
	return nil
}
