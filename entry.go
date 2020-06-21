package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

// an entry contains the details of the logging message
type entry struct {
	message string
	data    interface{}
	caller  *caller
	time    string

	callerSkip int
	logger     *logger
}

// MarshalJSON overrides default marshalling to allow for unexported field names
func (entry *entry) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Caller  *caller     `json:"caller,omitempty"`
		Time    string      `json:"time,omitempty"`
	}{
		Message: entry.message,
		Data:    entry.data,
		Caller:  entry.caller,
		Time:    entry.time,
	})

	if err != nil {
		return nil, err
	}

	return b, nil
}

// newEntry returns a new empty logging entry instance
func newEntry() *entry {
	return &entry{}
}

// Entry can be used to start a custom logging entry instance and must be the first method in a chained call.
// This is only required when not using the public functions that return *entry.
func Entry() *entry {
	var e *entry
	e = newEntry()

	var l *logger

	if reflect.DeepEqual(globalLogger, logger{}) {
		l = createDefaultLogger()

	} else {
		gl := globalLogger
		l = &gl
	}

	e.logger = l
	return e
}

// Debug sets the message to a debug level logging entry and calls send to log the message.
func (entry *entry) Debug(v ...interface{}) {

	if !entry.logger.isEnabled.debugMode {
		return
	}

	entry.time = time.Now().Format(time.RFC3339)
	entry.message = fmt.Sprint(v...)

	if entry.logger.isEnabled.setCaller && entry.caller == nil {
		entry.setCaller(2)
	}

	entry.send()
}

// Info sets the message to an info level logging entry and call send to log the message.
func (entry *entry) Info(v ...interface{}) {

	entry.time = time.Now().Format(time.RFC3339)
	entry.message = fmt.Sprint(v...)

	if entry.logger.isEnabled.setCaller && entry.caller == nil {
		entry.setCaller(2)
	}

	entry.send()
}

// SetLongFile sets the logger to use long file format with full path.
func (entry *entry) SetLongFile() *entry {
	entry.logger.isEnabled.shortFile = false

	return entry
}

// SetShortFile set the logger to use short file format with file name only.
func (entry *entry) SetShortFile() *entry {
	entry.logger.isEnabled.shortFile = true

	return entry
}

// WithCaller adds the caller information to a logging instance and should only be used when overriding logger behavior,
// ie. it should not be called by any other method in the simple-log package otherwise it will not report the correct
// caller function. Use setLogger(n int) within the package.
func (entry *entry) WithCaller() *entry {
	entry.setCaller(2)

	return entry

}

// WithStruct adds a struct of data to a logging instance.
func (entry *entry) WithStruct(data interface{}) *entry {
	entry.data = data
	return entry
}

// SetOutStream will set the output stream for a specific instance of a logger
func (entry *entry) SetOutStream(out io.Writer) *entry {
	entry.logger.setOutStream(out)

	return entry
}

// setCaller will add caller information of file, function, and line number to a logging instance.
func (entry *entry) setCaller(n int) {
	pc, file, line, _ := runtime.Caller(n)
	fn := runtime.FuncForPC(pc)

	var caller caller

	if entry.logger.isEnabled.shortFile {
		caller.File = filepath.Base(file)
	} else {
		caller.File = file
	}

	if fn != nil {
		caller.Function = fn.Name()

	}

	caller.Line = line

	entry.caller = &caller
}

// send is used to send the message to the output stream.
func (entry *entry) send() {

	b, _ := json.Marshal(entry)
	b = append(b, 10)

	entry.logger.out.Write(b)
}
