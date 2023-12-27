package logging

import (
	"fmt"
	"time"
)

var (
	Loggr     Logger
	LogLevels map[int32]string
)

const (
	INFO    = iota
	ERROR   = iota
	WARNING = iota
	DEBUG   = iota
)

type Logger struct{}

func init() {
	LogLevels := make(map[int]string)
	LogLevels[INFO] = "INFO"
	LogLevels[ERROR] = "ERROR"
	LogLevels[WARNING] = "WARNING"
	LogLevels[DEBUG] = "DEBUG"
	Loggr = Logger{}
}

func (l *Logger) writeLog(level int32, msg string) {
	currentTime := time.Now()
	fmt.Printf("[%s][@%s][Message: %s]", LogLevels[level], currentTime.Format("2006.01.02 15:04:05"), msg)
}

func (l *Logger) Info(msg string) {
	l.writeLog(INFO, msg)
}

func (l *Logger) Err(msg string) {
	l.writeLog(ERROR, msg)
}

func (l *Logger) Warn(msg string) {
	l.writeLog(WARNING, msg)
}

func (l *Logger) Debug(msg string) {
	l.writeLog(DEBUG, msg)
}
