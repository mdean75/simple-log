package simple_log

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
	Info(v ...interface{})
	//Debugf(format string, v ...interface{})
	//Infof(format string, v ...interface{})
}

type logger struct {
	isEnabled enabled // not data to be displayed
	out       io.Writer

	Message string        `json:"message"`
	Data    interface{}   `json:"data,omitempty"`
	Fields  []interface{} `json:"fields,omitempty"`
	Caller  *caller       `json:"caller,omitempty"`
	Time    string        `json:"time,omitempty"`
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
	File string `json:"file,omitempty"`
	Line int    `json:"line,omitempty"`
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

func (lgr *logger) Debug(v ...interface{}) {

	if lgr.isEnabled.debugMode {
		lgr.Time = time.Now().Format(time.RFC3339)

		lgr.Message = fmt.Sprint(v...)

	}
	lgr.send()
}

func (lgr *logger) Info(v ...interface{}) {

	lgr.Time = time.Now().Format(time.RFC3339)
	lgr.Message = fmt.Sprint(v...)

	lgr.send()
}

func (lgr *logger) WithStruct(data interface{}) *logger {
	lgr.Data = data
	return lgr
}

func (lgr *logger) WithCaller() *logger {
	lgr.setCaller(2)
	return lgr
}

func (lgr *logger) SetShortFile() *logger {
	lgr.isEnabled.shortFile = true

	return lgr
}

func (lgr *logger) setCaller(n int) {
	_, file, line, _ := runtime.Caller(n)

	var caller caller
	if lgr.isEnabled.shortFile {
		caller.File = filepath.Base(file)
	} else {
		caller.File = file
	}
	caller.Line = line

	lgr.Caller = &caller

}

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

func Debug(v ...interface{}) {
	var l *logger
	if reflect.DeepEqual(globalLogger, logger{}) {
		l = createDefaultLogger()
	} else {
		gl := globalLogger
		l = &gl
	}
	if l.isEnabled.debugMode {
		l.Debug(v...)
	}

}

func Info(v ...interface{}) {
	var l *logger
	if reflect.DeepEqual(globalLogger, logger{}) {
		l = createDefaultLogger()
	} else {
		gl := globalLogger
		l = &gl
	}
	l.Info(v...)

}
