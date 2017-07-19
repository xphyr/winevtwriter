// Copyright 2017 Mark DeNeve.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// +build windows

// Package winlogger implements a Writer interface to use
// the Windows event log with the standard go logger

package winlogger

// A Writer is a connection to a Windows Event log server.
type Writer struct {
	priority int    // this should hold one of the windows.EVENTLOG* constants
	source   string // this defines what the source will be listed as in eventlog
	hostname string
	network  string
	raddr    string

	mu  sync.Mutex    // guards conn
	evt *eventlog.Log // holds the event log connection
}

// Write sends a log message to the syslog daemon.
func (w *Writer) Write(b []byte) (int, error) {
	return w.writeAndRetry(w.priority, string(b))
}

func (w *Writer) writeAndRetry(p int, s string) (int, error) {

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.evt != nil {
		if n, err := w.write(p, s); err == nil {
			return n, err
		}
	}
	if err := w.connect(); err != nil {
		return 0, err
	}
	return w.write(p, s)
}

// Close closes a connection to the Windows Event Log.
func (w *Writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.evt != nil {
		err := w.evt.Close()
		w.evt = nil
		return err
	}
	return nil
}

// New establishes a new connection to the Windows Event System. Each
// write to the returned writer sends a log message with the given
// priority and prefix.
func New(priority int, source string) (*Writer, error) {
	return Dial("", "", priority, source)
}

// Dial establishes a connection to a log daemon by connecting to
// address raddr on the specified network. Each write to the returned
// writer sends a log message with the given facility, severity and
// source.
// If network is empty, Dial will connect to the local syslog server.
// Otherwise, see the documentation for net.Dial for valid values
// of network and raddr.
func Dial(network, raddr string, priority int, source string) (*Writer, error) {
	if priority < 0 || priority > windows.EVENTLOG_INFORMATION_TYPE {
		return nil, errors.New("xphyr/winlogger: invalid eventlog priority")
	}

	if source == "" {
		source = os.Args[0]
	}
	hostname, _ := os.Hostname()

	w := &Writer{
		priority: priority,
		source:   source,
		hostname: hostname,
		network:  network,
		raddr:    raddr,
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	err := w.connect()
	if err != nil {
		return nil, err
	}
	return w, err
}

// connect makes a connection to the syslog server.
// It must be called with w.mu held.
func (w *Writer) connect() (err error) {
	if w.evt != nil {
		// ignore err from close, it makes sense to continue anyway
		w.evt.Close()
		w.evt = nil
	}

	if w.network == "" {
		w.evt, err = eventlog.Open(w.source)
		if w.hostname == "" {
			w.hostname = "localhost"
		}
	} else {
		if err == nil {
			w.evt, err = eventlog.OpenRemote(w.hostname, w.source)
			if err != nil {
				return fmt.Errorf("xphyr/winlogger: unable to connect to remote eventlogger %v", w.hostname)
			}
		}
	}
	return
}

// write generates and writes a syslog formatted string. The
// format is as follows: <PRI>TIMESTAMP HOSTNAME source[PID]: MSG
func (w *Writer) write(p int, msg string) (int, error) {
	// ensure it ends in a \n
	// I dont think this is needed
	/* nl := ""
	if !strings.HasSuffix(msg, "\n") {
		nl = "\n"
	}
	*/

	switch p {
	case windows.EVENTLOG_ERROR_TYPE:
		err := w.evt.Error(1, msg)
		if err != nil {
			return 0, err
		}
	case windows.EVENTLOG_WARNING_TYPE:
		err := w.evt.Warning(1, msg)
		if err != nil {
			return 0, err
		}
	case windows.EVENTLOG_INFORMATION_TYPE:
		err := w.evt.Info(1, msg)
		if err != nil {
			return 0, err
		}
	}

	// Note: return the length of the input, not the number of
	// bytes printed by Fprintf, because this must behave like
	// an io.Writer.
	return len(msg), nil
}
