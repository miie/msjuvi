package logger

import (
	"log"
	"os"
	"fmt"
	"runtime"
	"path/filepath"
	"github.com/miie/msjuvi/mailer"
)

const (
	VERBOSE   = 10
	INFO      = 20
	WARNING   = 30
	ERROR     = 40
	FATAL     = 50
	NOLOGGING = 60
)

type LogSettings struct {
	LogFilePath   				string
	LogLevel      				int
	LogFatalToStdOut			bool
	LogAllLevelsToStdOut		bool	
	LogWithMail   				bool
	LogWithFilenameAndLineno 	bool
}

type CommodiLogger struct {
	LogLevel      				int
	LogFatalToStdOut			bool
	LogAllLevelsToStdOut		bool	
	LogToMail   				bool
	LogWithFilenameAndLineno 	bool	
}

func New(loglevel int, logfataltostdout bool, alwayslogtostdout bool, logtomail bool, 
	logwithfilenameandlineno bool) *CommodiLogger {
	return &CommodiLogger{loglevel, logfataltostdout, alwayslogtostdout, logtomail, logwithfilenameandlineno}
}

var settings = GetStdSettings()
var std = New(10, true, true, false, true)

func SetOutputToFile(path string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0660)
		if err != nil {
			// If we fail to set logging to file warn and use log default
			log.Println("logger.SetOutputToFile: cannot set logging to file. using default (os.Stderr) filepath & err:", path, err)
			return
		}
	}
	log.SetOutput(f)
}

func (l *CommodiLogger) DoLog(level int, s ...interface{}) {
	// Check log level
	if l.LogLevel > level {
		return
	}

	// Check if we should get filename and line number
	var logstr string
	if l.LogWithFilenameAndLineno {
		logstr = getLogStringWithFilenameAndLineno(s...)
	} else {
		logstr = fmt.Sprintln(s...)
	}
	
	// Check if we should always log to stdout
	if l.LogAllLevelsToStdOut {
		//TODO: implement!!! (check if we already log to stderr - if logging is set to that...)
	} else if level == FATAL && l.LogFatalToStdOut {
		//TODO: implement!!! (check if we already log to stderr - if logging is set to that...)
	}
	
	// Log
	log.Println(logstr)
	
	// Check if we should log to mail
	if l.LogToMail {
		mailer.SendError(logstr)
	}
}

func getLogStringWithFilenameAndLineno(s ...interface{}) string {
	_, f, l, ok := runtime.Caller(3)
	if ok == false {
		log.Println("logger.getLogStringWithFilenameAndLineno: runtime.Caller failed.")
		return fmt.Sprint(s...)
	}

	return fmt.Sprint(fmt.Sprint(filepath.Base(f), ":", l), " | ", fmt.Sprint(s...))
}

func GetStdSettings() LogSettings {
	return LogSettings{
		LogFilePath: "./pleaseChangeMyName.log",
		LogLevel: VERBOSE,
		LogFatalToStdOut: true,
		LogAllLevelsToStdOut: false,	
		LogWithMail: false,
		LogWithFilenameAndLineno: true,
	}
}

func LogFatal(s string) {
	// Add description

	// Check log level
	if settings.LogLevel > FATAL {
		return
	}
	// send mail
	if settings.LogWithMail == true {
		mailer.SendError(s)
	}

	// log Fatal which calls os.Exit(1)
	if settings.LogFatalToStdOut == true {
		fmt.Println(s)
	}
	log.Fatalf(s)
}

func LogError(s string) {
	// add description

	// Check log level
	if settings.LogLevel > ERROR {
		return
	}
	// send mail
	if settings.LogWithMail == true {
		mailer.SendError(s)
	}

	if settings.LogAllLevelsToStdOut == true {
		fmt.Println(s)
	}

	// log Panic which calls panic()
	log.Panic(s)
}

func LWarn_(s ...interface{}) {
	std.DoLog(WARNING, s...)
}
func LWarn(withFilenameLineno bool, s ...interface{}) {
	// Check log level
	if settings.LogLevel > WARNING {
		return
	}

	// send mail
	if settings.LogWithMail == true {
		mailer.SendError(fmt.Sprintln(s...))
	}

	if settings.LogAllLevelsToStdOut == true {	
		fmt.Println(s...)
	}

	if withFilenameLineno == true {
		log.Println(getLogStringWithFilenameAndLineno(s...))
	} else {
		log.Println(s...)
	}
}

// add description
func LogWarning(s ...interface{}) {

	// Check log level
	if settings.LogLevel > WARNING {
		return
	}

	// send mail
	if settings.LogWithMail == true {
		mailer.SendError(fmt.Sprintln(s...))
	}

	if settings.LogAllLevelsToStdOut == true {	
		fmt.Println(s...)
	}

	// log Print
	log.Println(getLogStringWithFilenameAndLineno(s...))
}

func LogInfo(s string) {
	// add description

	// Check log level
	if settings.LogLevel > INFO {
		return
	}

	// send mail
	if settings.LogWithMail == true {
		mailer.SendError(s)
	}

	if settings.LogAllLevelsToStdOut == true {
		fmt.Println(s)
	}

	// log Print
	log.Println(s)
}

func LogVerbose(s string) {
	// add description

	// Check log level
	if settings.LogLevel > VERBOSE {
		return
	}

	// send mail
	if settings.LogWithMail == true {
		mailer.SendError(s)
	}

	if settings.LogAllLevelsToStdOut == true {
		fmt.Println(s)
	}

	// log Print
	log.Println(s)
}

func PrintToStdOut(s string) {
	fmt.Println(s)
}