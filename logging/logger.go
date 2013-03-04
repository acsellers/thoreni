package logging

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	LOGTIME       = "2006-01-02 15:04:05 -0700"
	REQUEST_INFO  = "Started %s \"%s\" for %s at %s"
	REQUEST_CLOSE = "Rendered SOMETHING in %dns"
	Log           *Logger
)

/*
This logging package was written to implement the MiniLogger
idea. Each Contextable will get it's own MiniLogger and then
at the end, it will write the entire Request's log out to the 
main log in one go. If you have a different log.Logger that
you would like to use, replace logging.Log.Output with that
log.Logger and all of thoreni will use your log.

In general, code will log to the MiniLogger included in each 
request, while modules linked to thoreni or setup code 
should use logging.Log directly, or create their own 
MiniLoggers if they are putting a bunch of text over a while 
and would perfer the text to be uniterrupted.
*/
type Logger struct {
	Output            *log.Logger
	input             chan string
	ActiveMiniLoggers []*MiniLogger
}

// Logger.Write writes data to the output directly, without regular 
// logging information like time, etc. Use this to send messages
// that aren't related to a normal request, like loading information
// for a subsystem. Strings do not need a newline, this one acts like
// Println instead of Print.
func (log *Logger) Write(data ...interface{}) {
	log.input <- fmt.Sprint(data...)
}

// Logger.Writef is the same as Write, except it acts like Printf
// instead of just Printing. Note: you do not need to add a newline
// to your format string, Logger does it automatically.
func (log *Logger) Writef(template string, data ...interface{}) {
	log.input <- fmt.Sprintf(template+"\n", data...)
}

func (log *Logger) watch() {
	for {
		output := <-log.input
		log.Output.Println(output)
	}
}

type MiniLogger struct {
	Parent *Logger
	Data   string
	Begun  time.Time
}

func NewMiniLogger() *MiniLogger {
	return &MiniLogger{Parent: Log}
}

// This writes a preamble about the request like parameters,
// method and url and time we received the request.
func (ml *MiniLogger) LogRequest(req *http.Request) {
	ml.Begun = time.Now()
	ml.Data += fmt.Sprintf(REQUEST_INFO, req.Method, req.RequestURI, req.RemoteAddr, ml.Begun.Format(LOGTIME))
	ml.Data += "\n"
}

func (ml *MiniLogger) Write(things ...interface{}) {
	ml.Data += fmt.Sprintln(things...)
}
func (ml *MiniLogger) Writef(formatString string, things ...interface{}) {
	ml.Data += fmt.Sprintf(formatString, things...)
	ml.Data += "\n"
}
func (ml *MiniLogger) CloseRequest(responseType string) {
	ml.Data += fmt.Sprintf(REQUEST_CLOSE, time.Since(ml.Begun).Nanoseconds())
	ml.Data += "\n"
}
func (ml MiniLogger) Flush() {
	ml.Parent.Write(ml.Data)
	ml.Parent.Write("")
}
func init() {
	Log = new(Logger)
	Log.input = make(chan string)
	go Log.watch()
	determineOutput()
}

func determineOutput() {
	pwd, err := os.Getwd()
	if err != nil {
		maybeLogDirectory, err := os.Stat(filepath.Join(pwd, "log"))
		if err == nil {
			if maybeLogDirectory.IsDir() {
				setLogFile(filepath.Join(pwd, "log"))
			} else {
				defer Log.Write("Couldn't open a 'log' directory, create the directory for file logs")
				setstdout()
			}
		} else {
			defer Log.Write("Couldn't open a 'log' directory, create one for file logs")
			setstdout()
		}
	} else {
		defer Log.Write("Couldn't get our working directory, using STDOUT for logging")
		setstdout()
	}
}

func setstdout() {
	Log.Output = log.New(os.Stdout, "", 0)
}

func setLogFile(directory string) {
	_, err := os.Stat(filepath.Join(directory, "web.log"))
	if err == nil {
		appendLog, err := os.OpenFile(filepath.Join(directory, "web.log"), os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
		if err == nil {
			Log.Output = log.New(appendLog, "", 0)
			Log.Write("Appending new log entries to existing log file 'web.log'")
		} else {
			defer Log.Write("Couldn't open web.log, check the permissions")
			setstdout()
		}
	} else {
		newFile, err := os.Create(filepath.Join(directory, "web.log"))
		if err == nil {
			Log.Output = log.New(newFile, "", 0)
			Log.Write("Opening new log file 'web.log'")
		} else {
			defer Log.Write("Couldn't create a web.log file, check permissions")
			setstdout()
		}
	}
}
