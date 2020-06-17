# Simple Log

Simple-log is a simple structured logger for Go, that has a minimal API to keep it ... well ... simple.  

The inspiration for simple-log came from Dave Cheney's blog post concerning many logging packages providing too many options. 
You can read that blog post here -> https://dave.cheney.net/2015/11/05/lets-talk-about-logging.  As per his recommendation, 
simple-log provides only two levels of logging, debug and info. 

Simple-log only supports logging in JSON format and has no feature to change the logging format.  The default output is 
stdout but can be configured to use any output that implements the io.writer interface.  

Debug logging is off by default, however a custom global logger can be specified.

## Examples

### Default logger

```go
package main

import (
	log "github.com/mdean75/simple-log"
)

func main() {
    log.Debug("This will not display with the default logger")
    log.Info("But this will display with the default logger")
}

// output: {"message":"But this will display with the default logger","time":"2020-06-16T20:13:36-05:00"}

``` 

Produces Output: 
```bash
{"message":"But this will display with the default logger","time":"2020-06-16T20:13:36-05:00"}

```
*note the debug log does not display by default

### Custom logger with debug mode turned on

```go
package main

import (
	log "github.com/mdean75/simple-log"
	"os"
)

func init() {
	settings := log.NewLoggerSettings(log.NewEnabledSettings(true, false), os.Stdout)
	log.CustomLogger(settings)
}

func main() {
	log.Debug("This debug message will now display with the custom logger")
	log.Info("And of course this will also display")
}
```

Produces Output:
```bash
{"message":"This debug message will now display with the custom logger","time":"2020-06-16T20:13:36-05:00"}
{"message":"And of course this will also display","time":"2020-06-16T20:13:36-05:00"}
```
*note: now with the custom logger with debug mode enabled, debug level messages are logged.
