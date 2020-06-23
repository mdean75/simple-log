package logger

import (
	"io"
	"reflect"
)

// Debug is the public accessor for a debug level log message. This is used when not modifying the default or custom
// logger behaviour.
func Debug(v ...interface{}) {
	var e *entry
	e = newEntry()

	// use a copy of the globalLogger for the logging entry
	gl := globalLogger
	e.logger = &gl

	if e.logger.isEnabled.setCaller {
		e.setCaller(2)
	}

	// only process the logging message if debug mode is Enabled
	if e.logger.isEnabled.debugMode {
		e.Debug(v...)
	}
}

// Info is the public accessor for an info level log message. This is used when not modifying the default or custom
// logger behaviour.
func Info(v ...interface{}) {
	var e *entry
	e = newEntry()

	// use a copy of the globalLogger for the logging entry
	gl := globalLogger
	e.logger = &gl

	if e.logger.isEnabled.setCaller {
		e.setCaller(2)
	}

	e.Info(v...)
}

// WithCaller is used to add caller information to a logging instance and can be the first function in a chained call.
func WithCaller() *entry {
	var e *entry
	e = newEntry()

	// check if the globalLogger is a zero value logger and create a new default logger if needed
	if reflect.DeepEqual(globalLogger, logger{}) {
		globalLogger = *createDefaultLogger()
		e.logger = &globalLogger
	} else {
		gl := globalLogger
		e.logger = &gl
	}

	e.setCaller(2)

	return e
}

// WithStruct is used to add a struct of data to a logging instance and can be the first function in a chained call.
func WithStruct(data interface{}) *entry {
	var e *entry
	e = newEntry()

	// check if the globalLogger is a zero value logger and create a new default logger if needed
	if reflect.DeepEqual(globalLogger, logger{}) {
		globalLogger = *createDefaultLogger()
		e.logger = &globalLogger
	} else {
		gl := globalLogger
		e.logger = &gl
	}

	e.data = data

	return e
}

// SetOutStream is used to specify a different output stream for a logging instance and can be the first function in a chained call.
func SetOutStream(out io.Writer) *entry {
	var e *entry
	e = newEntry()

	// check if the globalLogger is a zero value logger and create a new default logger if needed
	if reflect.DeepEqual(globalLogger, logger{}) {
		globalLogger = *createDefaultLogger()
		e.logger = &globalLogger
	} else {
		gl := globalLogger
		e.logger = &gl
	}

	e.logger.out = out

	return e
}

// SetShortFile sets the caller information to use the short file format for a logging instance and can be the first function in a chained call.
func SetShortFile() *entry {
	var e *entry
	e = newEntry()

	// check if the globalLogger is a zero value logger and create a new default logger if needed
	if reflect.DeepEqual(globalLogger, logger{}) {
		globalLogger = *createDefaultLogger()
		e.logger = &globalLogger
	} else {
		gl := globalLogger
		e.logger = &gl
	}

	e.logger.isEnabled.shortFile = true

	return e
}

// SetLongFile sets the caller information to use the long file format for a logging instance and can be the first function in a chained call.
func SetLongFile() *entry {
	var e *entry // SetShortFile sets the caller information to use the short file format for a logging instance and can be the first function in a chained call.

	e = newEntry()

	// check if the globalLogger is a zero value logger and create a new default logger if needed
	if reflect.DeepEqual(globalLogger, logger{}) {
		globalLogger = *createDefaultLogger()
		e.logger = &globalLogger
	} else {
		gl := globalLogger
		e.logger = &gl
	}

	e.logger.isEnabled.shortFile = false

	return e
}
