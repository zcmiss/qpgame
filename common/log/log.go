package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var Log = logrus.New()

// 模块初始化
func init() {
	Log.Out = os.Stdout
	//输出格式为json
	Log.Formatter = new(logrus.JSONFormatter)
	Log.Level = logrus.DebugLevel
	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}
}

// 异常晶日志打印
func LogPrException(message string) {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, _ := runtime.Caller(5)
	funcname = runtime.FuncForPC(pc).Name()
	funcname = filepath.Ext(funcname)
	funcname = strings.TrimPrefix(funcname, ".")
	absPath, _ := filepath.Abs(filename) // /full/path/basename.go => basename.go
	Log.WithFields(logrus.Fields{
		"filename": absPath,
		"line":     line,
		"funcname": funcname,
	}).Info("异常原因: " + message)
}

// Panic追踪
func PanicTrace(kb int, err interface{}) {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	if stack != nil && err != nil {
		LogPrException(fmt.Sprintf("%v -------%s", err, string(stack)))
	} else {
		LogPrException(fmt.Sprintf("%v -------%s", "err is nil", "stack is nil"))
	}
}

// 延迟追踪?
func DeferRecover() {
	if err := recover(); err != nil {
		PanicTrace(4, err)
	}
}
