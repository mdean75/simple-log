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

type tempEntry interface {
	Info(v ...interface{})
}

// an entry contains the details of the logging message
type entry struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Caller  *caller     `json:"caller,omitempty"`
	Time    string      `json:"time,omitempty"`

	callerSkip int
	logger     *logger // ensure this is not serialized
}

// TODO: SEE IF I CAN ALLOW CALLING METHODS IN ANY ORDER IE. SETSHORTCALLER AFTER WITH CALLER
func newEntry() *entry {
	return &entry{}
}

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

func (entry *entry) Debug(v ...interface{}) {

	if !entry.logger.isEnabled.debugMode {
		return
	}
	entry.Time = time.Now().Format(time.RFC3339)
	entry.Message = fmt.Sprint(v...)

	entry.send()
}

func (entry *entry) Info(v ...interface{}) {

	if entry.logger.isEnabled.setCaller {
		if entry.callerSkip == 0 {
			entry.setCaller(3)
		} else {
			entry.setCaller(2)
		}

	}
	entry.Time = time.Now().Format(time.RFC3339)
	entry.Message = fmt.Sprint(v...)

	entry.send()
}

func (entry *entry) SetLongFile() *entry {
	entry.logger.isEnabled.shortFile = false

	return entry
}

func (entry *entry) SetShortFile() *entry {
	entry.logger.isEnabled.shortFile = true

	return entry
}

func (entry *entry) WithCaller() *entry {
	//entry.setCaller(2)
	entry.logger.isEnabled.setCaller = true
	entry.callerSkip = 2
	return entry

}

func (entry *entry) WithStruct(data interface{}) *entry {
	entry.Data = data
	return entry
}

func (entry *entry) send() {

	b, _ := json.Marshal(entry)
	b = append(b, 10)

	entry.logger.out.Write(b)
}

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

	entry.Caller = &caller
}

// SetOutStream will set the output stream for a specific instance of a logger
func (entry *entry) SetOutStream(out io.Writer) *entry {
	entry.logger.setOutStream(out)

	return entry
}
