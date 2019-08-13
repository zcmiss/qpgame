package common

import (
	"fmt"
	"path"
	"time"
)

//打印日志,支持变长数组模式, 只在本地有打印效果
func PrintLog(strs ...string) {
	if path.IsAbs("/home/god") {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] ", now)
		for _, s := range strs {
			fmt.Printf("%s ", s)
		}
		fmt.Print("\n")
	}
}
