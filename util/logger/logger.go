package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Level int

const (
	NON Level = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

var loggerLevel = map[string]Level{"ERROR": ERROR, "WARN": WARN, "INFO": INFO, "DEBUG": DEBUG, "TRACE": TRACE, "FATAL": FATAL}

var logLevel = NON
var logPath = ""

var currentYear, currentDay int
var currentMonth time.Month

var loggerFile *os.File
var logLocker sync.RWMutex

var traceLog *log.Logger
var debugLog *log.Logger
var infoLog *log.Logger
var warnLog *log.Logger
var errorLog *log.Logger
var fataLog *log.Logger

func init() {
	if logLevel == NON {
		level := "TRACE"
		if logLevel = loggerLevel[strings.ToUpper(level)]; logLevel == NON {
			logLevel = INFO
		}
	}
	currentYear, currentMonth, currentDay = GetNow().Date()
	initLoggger(getIOWriter(), getIOWriter2())
}

func init0(level string, lp string) {
	if logLevel = loggerLevel[strings.ToUpper(level)]; logLevel == NON {
		logLevel = INFO
	}
	logPath = lp
	if lp != "" && lp != "console" {
		logLocker.Lock()
		defer logLocker.Unlock()

		writer := getIOWriter()
		initLoggger(writer, writer)
	}
}

func initLoggger(writer io.Writer, errWriter io.Writer) {
	traceLog = log.New(writer, "[TRACE]", log.Ldate|log.Ltime|log.Lshortfile)
	debugLog = log.New(writer, "[DEBUG]", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog = log.New(writer, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	warnLog = log.New(writer, "[WARN]", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(writer, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
	fataLog = log.New(writer, "[FATAL]", log.Ldate|log.Ltime|log.Lshortfile)
}

func splitter() string {
	if runtime.GOOS == "windows" {
		return `\`
	} else {
		return "/"
	}
}

func getIOWriter() io.Writer {
	if logPath != "" && strings.ToLower(logPath) != "console" {
		_, err := os.Stat(logPath)
		if os.IsNotExist(err) {
			if err = os.MkdirAll(logPath, os.ModePerm); err != nil {
				fmt.Println("Make log dir error")
				return nil
			}
		}
		file := fmt.Sprintf("%s%slog-%d-%d-%d.log", logPath, splitter(), currentYear, currentMonth, currentDay)
		logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		if err != nil {
			panic(err)
		}
		loggerFile = logFile
		return logFile
	} else {
		return os.Stdout
	}
}

func checkDate() {
	cYear, cMonth, cDay := GetNow().Date()
	if cDay != currentDay || cMonth != currentMonth || cYear != currentYear {
		logLocker.Lock()
		defer logLocker.Unlock()

		if loggerFile != nil {
			_ = loggerFile
		}
		writer := getIOWriter()
		if logPath != "" && strings.ToLower(logPath) != "console" {
			initLoggger(writer, writer)
		}
	}
}

func getIOWriter2() io.Writer {
	if logPath != "" && strings.ToLower(logPath) != "console" {
		return nil
	} else {
		return os.Stderr
	}
}

func GetNow() time.Time {
	cst, err := time.LoadLocation("Asia/shanghai")
	if err != nil {
		return time.Now()
	}
	return time.Now().In(cst)
}

func Trace(contens ...interface{}) {
	checkDate()

	logLocker.RLock()
	defer logLocker.RUnlock()

	if logLevel >= TRACE {
		_ = traceLog.Output(2, fmt.Sprintln(contens...))
	}
}

func Tracef(contents string, v ...interface{}) {
	Trace(fmt.Sprintf(contents, v...))
}

func Debug(contents ...interface{}) {
	checkDate()

	logLocker.RLock()
	defer logLocker.RUnlock()

	if logLevel >= DEBUG {
		_ = debugLog.Output(2, fmt.Sprintln(contents...))
	}
}

func Debugf(contents string, v ...interface{}) {
	Debug(fmt.Sprintf(contents, v...))
}

func Info(contents ...interface{}) {
	checkDate()

	logLocker.RLock()
	defer logLocker.RUnlock()
	if logLevel >= INFO {
		_ = infoLog.Output(2, fmt.Sprintln(contents))
	}
}

func Infof(contents string, v ...interface{}) {
	Info(fmt.Sprintf(contents, v))
}

func Warn(contents ...interface{}) {
	checkDate()

	logLocker.RLock()
	defer logLocker.RUnlock()

	if logLevel >= WARN {
		_ = warnLog.Output(2, fmt.Sprintln(contents...))
	}
}

func Warnf(contents string, v ...interface{}) {
	Warn(fmt.Sprintf(contents, v))
}

func Error(contents ...interface{}) {
	checkDate()

	logLocker.RLock()
	defer logLocker.RUnlock()

	if logLevel >= ERROR {
		_ = errorLog.Output(2, fmt.Sprintln(contents))
	}
}

func Errorf(contents string, v ...interface{}) {
	checkDate()
	logLocker.RLock()
	defer logLocker.RUnlock()

	if logLevel >= ERROR {
		_ = errorLog.Output(2, fmt.Sprintf(contents, v))
	}
}

func Fatal(contents ...interface{}) {
	checkDate()

	logLocker.RLock()
	defer logLocker.RUnlock()

	if logLevel >= FATAL {
		s := fmt.Sprint(contents)
		fataLog.Fatal(s)
	}
}
