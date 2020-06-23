package logger

import (
	"io"
	"os"
	"sync"
)

var globalLogger = *createDefaultLogger()
var mu = sync.Mutex{}

type logging interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	//Debugf(format string, v ...interface{})
	//Infof(format string, v ...interface{})
}

// a logger specifies configuration for the logger
type logger struct {
	isEnabled Enabled
	out       io.Writer
	mu        *sync.Mutex
}

// logger settings
type Enabled struct {
	debugMode bool
	shortFile bool
	setCaller bool
}

// NewEnabled creates a new instance of the Enabled (logger settings) struct
func NewEnabled(debug, shortFile, caller bool) *Enabled {
	return &Enabled{
		debugMode: debug,
		shortFile: shortFile,
		setCaller: caller,
	}
}

type caller struct {
	File     string `json:"file,omitempty"`
	Function string `json:"function,omitempty"`
	Line     int    `json:"line,omitempty"`
}

func newLogger(isEnabled *Enabled, out io.Writer) *logger {
	return &logger{
		isEnabled: *isEnabled,
		out:       out,
		mu:        &sync.Mutex{},
	}
}

func createDefaultLogger() *logger {
	return newLogger(&Enabled{
		debugMode: false,
		shortFile: true,
		setCaller: false,
	}, os.Stdout)

}

// CustomLogger creates a logger to override the default and sets this as the globalLogger
func CustomLogger(settings Enabled, out io.Writer) {
	globalLogger = logger{
		isEnabled: Enabled{
			debugMode: settings.debugMode,
			shortFile: settings.shortFile,
			setCaller: settings.shortFile,
		},
		out: out,
		mu:  &sync.Mutex{},
	}
}

func (lgr *logger) setOutStream(out io.Writer) *logger {
	lgr.out = out

	return lgr
}
