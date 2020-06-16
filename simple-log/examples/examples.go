package main

import "simple-log/logger"

func simpleDebugExample() {
	logger.Debug("Hello, this is a test of a debug level log message using all defaults")
}

func main() {
	simpleDebugExample()
}
