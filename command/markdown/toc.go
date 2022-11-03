package markdown

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/urfave/cli"
)

// Toc 查看目录下 markdown 文件
var Toc = cli.Command{
	Name:     "markdown",
	Usage:    "markdown commands: toc.",
	Category: "markdown",
	Subcommands: []cli.Command{
		{
			Name:   "toc",
			Usage:  "create toc for all markdown files",
			Action: toc,
		},
	},
}

// toc executes update command and return exit code.
func toc(ctx *cli.Context) error {
	curPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return err
	}

	var result []string
	baseDir := curPath
	filepath.Walk(curPath, func(filepath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		result = append(result, filepath)
		return nil
	})

	sort.Strings(result)
	fmt.Println("- toc")
	for _, filepath := range result {
		if strings.HasSuffix(filepath, ".md") {
			showName := strings.TrimLeft(filepath, baseDir)
			if i := strings.LastIndex(showName, "/"); i > 0 {
				showName = showName[i+1:]
			}
			fileURI := strings.Replace(filepath, baseDir, ".", 1)
			markdownText := fmt.Sprintf("- [%v](%v)", showName, fileURI)
			fmt.Println(markdownText)
		}
	}

	return nil
}
