package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var logger *log.Logger

// Init ...
func Init() {
	// todo
	lf, _ := logFile("test")
	logger = log.New(lf, "", log.Llongfile|log.LstdFlags)
}

func logFile(fileName string) (*os.File, error) {
	// todo, read the setting and decide the path.
	wd, _ := os.Getwd()
	f, _ := os.OpenFile(wd+"logs"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	return f, nil
}

// Debug ...
func Debug(v ...interface{}) {
	prefix("DEBUG")
	logger.Println(v...)
}

// Info ...
func Info(v ...interface{}) {
	prefix("INFO")
	logger.Println(v...)
}

// Warn ...
func Warn(v ...interface{}) {
	prefix("WARN")
	logger.Println(v...)
}

// Error ...
func Error(v ...interface{}) {
	prefix("ERROR")
	logger.Println(v...)
}

// Fatal ...
func Fatal(v ...interface{}) {
	prefix("FATAL")
	logger.Println(v...)
}

func prefix(level string) {
	_, file, line, ok := runtime.Caller(2)
	var _prefix string
	if ok {
		_prefix = fmt.Sprintf("[%s][%s:%d]", level, filepath.Base(file), line)
	} else {
		_prefix = fmt.Sprintf("[%s]", level)
	}
	logger.SetPrefix(_prefix)
}
