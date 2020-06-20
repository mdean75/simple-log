package logger

import (
	"io"
	"os"
)

var std = createDefaultLogger()
var globalLogger logger

type logging interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	//Debugf(format string, v ...interface{})
	//Infof(format string, v ...interface{})
}

// a logger specifies configuration for the logger
type logger struct {
	isEnabled Enabled // not data to be displayed
	out       io.Writer
}

type Enabled struct {
	debugMode bool
	shortFile bool
	setCaller bool
}

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
	}
}

func createDefaultLogger() *logger {
	return newLogger(&Enabled{
		debugMode: false,
		shortFile: true,
		setCaller: true,
	}, os.Stdout)

}

func CustomLogger(settings Enabled, out io.Writer) {
	globalLogger = logger{
		isEnabled: Enabled{
			debugMode: settings.debugMode,
			shortFile: settings.shortFile,
			setCaller: settings.shortFile,
		},
		out: out,
	}
}

// TODO: ensure out is a valid type or is this even needed
func (lgr *logger) setOutStream(out io.Writer) *logger {
	lgr.out = out

	return lgr
}
