package markdown

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

var (
	summary string
)

const (
	// README TODO
	README      = "README.md"
	summaryFlag = "<!-- summary -->"
)

// toc executes update command and return exit code.
func toc(ctx *cli.Context) error {
	summary = ""
	curPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return err
	}

	titleLevel := int64(2)
	if len(os.Args) > 3 {
		level, _ := strconv.ParseInt(os.Args[3], 10, 64)
		if level > 0 {
			titleLevel = level
		}
	}
	DFS(curPath, titleLevel)

	if summary == "" {
		return nil
	}

	content, err := os.ReadFile(README)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if ct := strings.Count(string(content), summaryFlag); ct != 2 {
		return fmt.Errorf("未识别到目录标识对，summaryFlag[%s] ct[%v]", summaryFlag, ct)
	}
	s := strings.Index(string(content), summaryFlag)
	e := strings.LastIndex(string(content), summaryFlag)
	result := string(content)[:s] +
		summaryFlag + "\n" + summary + "\n" + summaryFlag +
		string(content)[e+len(summaryFlag):]
	return os.WriteFile(README, []byte(result), fs.FileMode(os.O_CREATE))
}

// ListAll TODO
func ListAll() {
	pwd, _ := os.Getwd() // 获取当前目录
	// 获取文件或目录相关信息
	DFS(pwd, 2)
}

// DFS 深度优先遍历文件+目录（优先展示目录）
func DFS(dirname string, titleLevel int64) {
	if titleLevel > 5 {
		return
	}
	pwd, _ := os.Getwd()

	// 获取文件或目录相关信息
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	level := strings.Repeat("#", int(titleLevel))
	// 1. 文件
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		if !strings.HasSuffix(fileInfo.Name(), ".md") {
			continue
		}
		fileURI := strings.Replace(dirname+"/"+fileInfo.Name(), pwd, ".", 1)

		markdownText := fmt.Sprintf("- [%v](%v)\n", fileInfo.Name(), fileURI)
		// fmt.Println(markdownText)
		summary += markdownText
	}
	// 2. 子目录
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			continue
		}
		markdownText := level + " " + fileInfo.Name() + "\n\n"
		// fmt.Println(markdownText)
		summary += markdownText
		DFS(dirname+"/"+fileInfo.Name(), titleLevel+1)
	}
}
