package hexo

import (
	"io"
	"log"
	"strings"
)

//通过管道同步获取日志的函数
func syncLog(reader io.ReadCloser) {
	//因为logger的print方法会自动添加一个换行，所以我们需要一个cache暂存不满一行的log
	cache := ""
	buf := make([]byte, 1024, 1024)
	for {
		strNum, err := reader.Read(buf)
		if strNum > 0 {
			outputByte := buf[:strNum]
			//这里的切分是为了将整行的log提取出来，然后将不满整行和下次一同打印
			outputSlice := strings.Split(string(outputByte), "\n")
			logText := strings.Join(outputSlice[:len(outputSlice)-1], "\n")
			log.Printf("%s%s", cache, logText)
			cache = outputSlice[len(outputSlice)-1]
		}
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "file already closed") {
				err = nil
			}
		}
	}
}
