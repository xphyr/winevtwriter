// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.package service

//+build !windows

package winlogger

import (
	"errors"
)

// A Writer is a connection to a Windows Event log server.
type Writer struct {
	priority int    // this should hold one of the windows.EVENTLOG* constants
	source   string // this defines what the source will be listed as in eventlog
	hostname string
	network  string
	raddr    string
}

func New(priority int, source string) (*Writer, error) {
	return nil, errors.New("xphyr/winlogger: not sure how you got here this is not a windows machine")
}

// Write sends a log message to the syslog daemon.
func (w *Writer) Write(b []byte) (int, error) {
	return 0, nil
}
