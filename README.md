# Simple Log

Simple-log is a simple structured logger for Go that has a minimal API to keep it ... well ... simple.  

The inspiration for simple-log came from Dave Cheney's blog post concerning many logging packages providing too many options. 
You can read that blog post here -> https://dave.cheney.net/2015/11/05/lets-talk-about-logging.  As per his recommendation, 
simple-log provides only two levels of logging, debug and info. 

Simple-log only supports logging in JSON format and has no feature to change the logging format.  The default output is 
stdout but can be configured to use any output that implements the io.Writer interface.  

Debug logging is off by default, however a custom global logger can be specified.

## Examples

### Default logger

```
package main

import (
	log "github.com/mdean75/simple-log"
)

func main() {
    log.Debug("This will not display with the default logger")
    log.Info("But this will display with the default logger")
}
``` 

Produces Output: 
```bash
{"message":"But this will display with the default logger","time":"2020-06-16T20:13:36-05:00"}
```
*note the debug log does not display by default

### Custom logger with debug mode turned on

```
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

## Use method chaining to add additional context

To add additional context to a log message, first call Entry().

### Add caller information to log messages
By default the short file name is reported but can be overridden to include the full path.

```
package main

import (
	log "github.com/mdean75/simple-log"
)

func main() {
    log.Entry().WithCaller().Info("Log this message with caller information")
    log.Entry().SetLongFile().WithCaller().Info("Log this message with caller in long file format")
    // note that SetLongFile() must be called before WithCaller() when setting the long file by method chaining
    // this can also be set globally by creating custom logging settings with shortFile = false
}
```

Produces Output:
```shell script
{"message":"Log this message with caller information","caller":{"file":"main.go","line":8},"time":"2020-06-16T22:44:19-05:00"}
{"message":"Log this message with caller in long file format","caller":{"file":"/home/USERNAME/projects/my-logger-project/examples/test.go","line":9},"time":"2020-06-16T23:08:53-05:00"}
```

### Add a struct to the logging message

```
package main

import (
    log "github.com/mdean75/simple-log"
)

func main() {
    testData := struct {
        Name string `json:"name"`
        Age int `json:"age"`
        Married bool `json:"married"`
    }{Name: "Michael", Age: 47, Married: true}

    log.Entry().WithStruct(testData).Info("A logging message with an included struct")
}
```

Produces Output:

```shell script
{"message":"A logging message with an included struct","data":{"name":"Michael","age":47,"married":true},"time":"2020-06-16T23:20:30-05:00"}
```
*note: The message will always be listed first in the JSON output and additional calls to WithStruct() will overwrite any 
previous calls so only the last data will be included in the log message.

### Write logs to any io.Writer

Because simple-log uses the io.Writer interface, logs can be written to any implementation such as http.ResponseWriter.

```
func health() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusAccepted)
        
        log.Entry().SetOutStream(w).WithCaller().Info("Using simple-log, we can write to any io.Writer including http.ResponseWriter")
    }
}
```
Responds with:
```
{
  "message": "Using simple-log, we can write to any io.Writer including http.ResponseWriter",
  "caller": {
    "file": "main.go",
    "function": "main.health.func1",
    "line": 60
  },
  "time": "2020-06-18T07:41:28-05:00"
}
```

## Entry Methods

The below list contains all the methods that can be called on a logger

* Overriding methods alter the default or global logger for this instance only
  * Entry()         *note: must be the first method called when using any other modifier method
  * SetShortFile()  *note: must be called before WithCaller()
  * SetLongFile()   *note: must be called before WithCaller()
  * WithCaller()
  * WithStruct()
  * SetOutStream()

* Logging methods, send the log message at the corresponding level. In a chained method call this must be the last method called.
  * Debug()
  * Info()
