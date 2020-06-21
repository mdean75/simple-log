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

	if e.logger.isEnabled.setCaller {
		e.setCaller(2)
	}

	// only process the logging message if debug mode is Enabled
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

	if e.logger.isEnabled.setCaller {
		e.setCaller(2)
	}

	e.Info(v...)
}
