package main

import (
	"log"
	"runtime"

	"github.com/xphyr/winlogger"
)

func main() {

	if runtime.GOOS == "windows" {
		// Configure logger to write to the syslog. You could do this in init(), too.
		logwriter, e := winlogger.New(winlogger.EVENTLOG_ERROR_TYPE, "myprog")
		if e == nil {
			log.SetOutput(logwriter)
		}

		// Now from anywhere else in your program, you can use this:
		log.Print("Hello Windows Event Logs!")
	} else {
		log.Print("Hello something that isnt windows")
	}

}
