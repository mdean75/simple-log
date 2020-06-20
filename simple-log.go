package logger

import "reflect"

func Debug(v ...interface{}) {
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

	// only process the logging message if debug mode is enabled
	if e.logger.isEnabled.debugMode {
		e.Debug(v...)
	}
}

func Info(v ...interface{}) {
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

	// TODO: see about moving logging functions to methods of entry instead of logger and create new entry with the logger in the struct
	e.Info(v...)
}
