package space

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/LeungGeorge/grimoire-tool/lib/gtools"
	"github.com/urfave/cli"
)

// Space 查看磁盘空间
var Space = cli.Command{
	Name:   "space",
	Usage:  "space of current directory.",
	Action: space,
}

// space executes update command and return exit code.
func space(ctx *cli.Context) error {
	curPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return err
	}

	files, err := ioutil.ReadDir(curPath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	pList := make(PairList, len(files))
	for idx, file := range files {
		fullPath := fmt.Sprintf("%s/%s", curPath, file.Name())
		if file.IsDir() {
			sz, err := gtools.DirSize(fullPath)
			if err != nil {
				log.Fatal(err)
				return err
			}
			pList[idx] = Pair{
				Key:   file.Name(),
				Value: sz,
			}
		} else {
			pList[idx] = Pair{
				Key:   file.Name(),
				Value: uint64(file.Size()),
			}
		}
	}

	sort.Sort(pList)

	for _, e := range pList {
		fmt.Printf("%s\t%s\n", gtools.HumanBytes(e.Value), e.Key)
	}

	return nil
}
