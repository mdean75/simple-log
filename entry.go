package logger

import (
	"encoding/json"
	"fmt"
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

	logger *logger // ensure this is not serialized
}

func (entry *entry) SetLongFileNew() *entry {
	entry.logger.isEnabled.shortFile = false

	return entry
}

func (entry *entry) SetShortFileNew() *entry {
	entry.logger.isEnabled.shortFile = true

	return entry
}

func (entry *entry) WithCallerNew() *entry {
	entry.setCallerNew(2)
	return entry
}

func (entry *entry) WithStructNew(data interface{}) *entry {
	entry.Data = data
	return entry
}

func (entry *entry) Info(v ...interface{}) {

	entry.Time = time.Now().Format(time.RFC3339)
	entry.Message = fmt.Sprint(v...)

	entry.sendEntry()
}

func (entry *entry) sendEntry() {

	b, _ := json.Marshal(entry)
	b = append(b, 10)

	entry.logger.out.Write(b)
}

func (entry *entry) setCallerNew(n int) {
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

func EntryNew() *entry {
	var e *entry
	e = newEntry()

	var l *logger

	if reflect.DeepEqual(globalLogger, logger{}) {
		//globalLogger = *createDefaultLogger()
		l = createDefaultLogger()

	} else {
		gl := globalLogger
		l = &gl
	}

	e.logger = l
	return e
}

func newEntry() *entry {
	return &entry{}
}
