package git

import (
	"log"
	"os/exec"
	"strings"
)

func status() int {
	cmd := exec.Command("git", "status", "-sb")
	bt, _ := cmd.Output()
	sLine := strings.Split(string(bt), "\n")

	// 是否存在未 push 的 commit
	ahead := false
	hasFile := false
	for i, s := range sLine {
		if i == 0 {
			ahead = strings.Contains(strings.ToLower(s), "ahead")
		} else if len(s) > 0 {
			hasFile = true
		}
	}

	log.Println("status:")
	log.Printf("%s\n", string(bt))
	if hasFile {
		return statusHasFile
	}

	if ahead {
		return statusAHead
	}

	return statusEmpty
}
