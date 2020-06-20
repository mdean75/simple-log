package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

var std = createDefaultLogger()
var globalLogger logger

type logging interface {
	Debug(v ...interface{})
	//Info(v ...interface{})
	//Debugf(format string, v ...interface{})
	//Infof(format string, v ...interface{})
}

// a logger specifies configuration for the logger
type logger struct {
	isEnabled enabled // not data to be displayed
	out       io.Writer

	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Caller  *caller     `json:"caller,omitempty"`
	Time    string      `json:"time,omitempty"`
}

type enabled struct {
	debugMode bool
	shortFile bool
	setCaller bool
}

type loggerSettings struct {
	isEnabled enabled
	out       io.Writer
}

type caller struct {
	File     string `json:"file,omitempty"`
	Function string `json:"function,omitempty"`
	Line     int    `json:"line,omitempty"`
}

func NewLoggerSettings(isEnabled *enabled, out io.Writer) *loggerSettings {
	return &loggerSettings{
		isEnabled: *isEnabled,
		out:       out,
	}
}

func NewEnabledSettings(debug, shortFile bool) *enabled {
	return &enabled{
		debugMode: debug,
		shortFile: shortFile,
	}
}

func createDefaultLogger() *logger {
	return NewLogger(&loggerSettings{
		isEnabled: enabled{
			debugMode: false,
			shortFile: true,
			setCaller: true,
		},
		out: os.Stdout,
	})
}

func CustomLogger(settings *loggerSettings) {
	globalLogger = logger{
		isEnabled: enabled{
			debugMode: settings.isEnabled.debugMode,
			shortFile: settings.isEnabled.shortFile,
			setCaller: settings.isEnabled.setCaller,
		},
		out: settings.out,
	}
}

func NewLogger(settings *loggerSettings) *logger {
	return &logger{isEnabled: enabled{
		debugMode: settings.isEnabled.debugMode,
		shortFile: settings.isEnabled.shortFile,
		setCaller: settings.isEnabled.setCaller,
	}, out: settings.out}
}

func (lgr *logger) send() {

	b, _ := json.Marshal(lgr)
	b = append(b, 10)

	lgr.out.Write(b)
}

// deprecated
func (lgr *logger) Debug(v ...interface{}) {

	if !lgr.isEnabled.debugMode {
		return
	}
	lgr.Time = time.Now().Format(time.RFC3339)

	lgr.Message = fmt.Sprint(v...)
	lgr.send()
}

// deprecated
func (lgr *logger) WithStruct(data interface{}) *logger {
	lgr.Data = data
	return lgr
}

// deprecated
func (lgr *logger) WithCaller() *logger {
	lgr.setCaller(2)
	return lgr
}

// deprecated
func (lgr *logger) SetShortFile() *logger {
	lgr.isEnabled.shortFile = true

	return lgr
}

// deprecated
func (lgr *logger) SetLongFile() *logger {
	lgr.isEnabled.shortFile = false

	return lgr
}

// TODO: ensure out is a valid type or is this even needed
// deprecated
func (lgr *logger) SetOutStream(out io.Writer) *logger {
	lgr.out = out

	return lgr
}

// deprecated
func (lgr *logger) setCaller(n int) {
	pc, file, line, _ := runtime.Caller(n)
	fn := runtime.FuncForPC(pc)

	var caller caller
	if lgr.isEnabled.shortFile {
		caller.File = filepath.Base(file)
	} else {
		caller.File = file
	}

	if fn != nil {
		caller.Function = fn.Name()

	}

	caller.Line = line

	lgr.Caller = &caller

}

// deprecated
// Entry must be called first in order to override logger settings for a specific log message.
func Entry() *logger {
	var l *logger

	if reflect.DeepEqual(globalLogger, logger{}) {
		l = createDefaultLogger()
	} else {
		gl := globalLogger
		l = &gl
	}

	return l
}

/*
	Public functions to initiate the logging start here
*/

//func Debug(v ...interface{}) {
//	var l *logger
//
//	// check if the globalLogger is a zero value logger and create a new default logger if needed
//	if reflect.DeepEqual(globalLogger, logger{}) {
//		globalLogger = *createDefaultLogger()
//		l = &globalLogger
//	} else {
//		gl := globalLogger
//		l = &gl
//	}
//
//	// only process the logging message if debug mode is enabled
//	if l.isEnabled.debugMode {
//		l.Debug(v...)
//	}
//}
//
//func Info(v ...interface{}) {
//	var e *entry
//	e = newEntry()
//
//	// check if the globalLogger is a zero value logger and create a new default logger if needed
//	if reflect.DeepEqual(globalLogger, logger{}) {
//		globalLogger = *createDefaultLogger()
//		e.logger = &globalLogger
//	} else {
//		gl := globalLogger
//		e.logger = &gl
//	}
//
//	// TODO: see about moving logging functions to methods of entry instead of logger and create new entry with the logger in the struct
//	e.Info(v...)
//}
