package git

import (
	"log"

	"github.com/urfave/cli"
)

// sync executes github_sync command and return exit code.
func sync(ctx *cli.Context) error {
	switch status() {
	case statusEmpty:
		log.Println("Execute git status finished, nothing to commit, working tree clean")
		// pull()

	case statusAHead:
		log.Println("Execute git status finished, ahead")
		pull()
		push()

	case statusHasFile:
		log.Printf("has file\n")
		log.Println("Execute git status finished, some file will be committed")
		add()
		commit()
		pull()
		push()
	}

	log.Println("Execute Command finished.")
	return nil
}
