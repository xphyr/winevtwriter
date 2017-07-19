Winlogger Package
=================

This package is designed to wrap the Windows Event log package (golang.org/x/sys/windows/svc/eventlog) and allow it to work with the base go logging package.  Using the "syslog" package as an example of how to do this, so you will find many of the same methoods.  Also thanks to Kardianos (github.com/kardianos) as I used some of his code as a refrence for how to use the eventlog package.

Simple Usecase
--------------

Example code of what I am trying to get to:
NOTE:  priority can be one of the following:

- EVENTLOG_ERROR_TYPE
- EVENTLOG_WARNING_TYPE
- EVENTLOG_INFORMATION_TYPE

```go
package main

import(
    "log"
    "github.com/xphyr/winlogger"
)

func main() {

    // Configure logger to write to the syslog. You could do this in init(), too.
    logwriter, e := winlogger.New(windows.EVENTLOG_ERROR_TYPE, "myprog")
    if e == nil {
        log.SetOutput(logwriter)
    }

    // Now from anywhere else in your program, you can use this:
    log.Print("Hello Logs!")
}
```

bugs
----

as of right now there are none (or perhaps the entire package could be considered a bug since it doesnt work)
