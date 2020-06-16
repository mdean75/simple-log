package main

import (
	log "simple-log"
)

func simpleDebugExample() {
	log.Debug("Hello, this is a test of a debug level log message using all defaults")
}

func main() {
	simpleDebugExample()
}
