package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info]\033[0m", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

var (
	Info   = infoLog.Println
	Infof  = infoLog.Printf
	Error  = errorLog.Println
	Errorf = errorLog.Printf
)

// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

func InfoW(v ...interface{}) {
	var str []interface{} = make([]interface{}, 1)
	str[0] = "\033[0;30;46m"
	v = append(str, v)
	v = append(v, "\033[0m")
	Info(v...)
}

func InfofB(str string, v ...interface{}) {
	Infof("\033[0;30;47m"+str+"\033[0m", v...)
}

func InfofW(str string, v ...interface{}) {
	Infof("\033[0;30;47m"+str+"\033[0m", v...)
}

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
