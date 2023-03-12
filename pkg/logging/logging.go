package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/rootlulu/go-gin-biu-biu-biu/internal/config"
)

var loggerFile, loggerConsole *log.Logger

// Init ...
func Init() {
	// todo, can configure the log output.
	lf, _ := logFile("test")
	loggerFile = log.New(lf, "", log.Llongfile|log.LstdFlags)
	loggerConsole = log.New(os.Stderr, "", log.Llongfile|log.LstdFlags)
}

func logFile(fileName string) (*os.File, error) {
	// todo, read the setting and decide the path.
	wd, _ := os.Getwd()
	f, err := os.OpenFile(filepath.Join(wd, config.App.LogPath, time.Now().Format(config.App.LogFormat)+config.App.LogSuffix), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("The logger file can't opened.", err)
	}
	return f, nil
}

// Debug ...
func Debug(v ...interface{}) {
	prefix("DEBUG")
	loggerFile.Println(v...)
	loggerConsole.Println(v...)
}

// Info ...
func Info(v ...interface{}) {
	prefix("INFO")
	loggerFile.Println(v...)
	loggerConsole.Println(v...)
}

// Warn ...
func Warn(v ...interface{}) {
	prefix("WARN")
	loggerFile.Println(v...)
	loggerConsole.Println(v...)
}

// Error ...
func Error(v ...interface{}) {
	prefix("ERROR")
	loggerFile.Println(v...)
	loggerConsole.Println(v...)
}

// Fatal ...
func Fatal(v ...interface{}) {
	prefix("FATAL")
	loggerFile.Println(v...)
	loggerConsole.Println(v...)
}

func prefix(level string) {
	_, file, line, ok := runtime.Caller(2)
	var _prefix string
	if ok {
		_prefix = fmt.Sprintf("[%s][%s:%d]", level, filepath.Base(file), line)
	} else {
		_prefix = fmt.Sprintf("[%s]", level)
	}
	loggerFile.SetPrefix(_prefix)
	loggerConsole.SetPrefix(_prefix)
}
