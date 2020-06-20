package main

import (
	log "github.com/mdean75/simple-log"
	"os"
)

func simpleDebugExample() {
	log.Debug("Hello, this is a test of a debug level log message using all defaults")
}

func simpleInfoExample() {
	log.Info("Hello, this is a test of an info level log message using all defaults")
}

func customLogger() {
	settings := log.NewLoggerSettings(log.NewEnabledSettings(true, false), os.Stdout)
	log.CustomLogger(settings)
}
func main() {

	//customLogger()
	//
	//simpleDebugExample()
	//simpleInfoExample()
	//
	//log.Entry().WithCaller().Debug("this is a test debug with caller")
	//log.Entry().SetShortFile().WithCaller().Debug("this should now have the short file for the caller")
	//log.Entry().WithCaller().SetShortFile().Debug("but this will still have the long file")
	log.Info("test")
	log.Info("test 2")

}
