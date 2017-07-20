WinEvtWriter Package
=================

This package is designed to wrap the Windows Event log package (golang.org/x/sys/windows/svc/eventlog) and allow it to work with the base go logging package.  Using the "syslog" package as an example of how to do this, so you will find many of the same methoods.  Also thanks to Kardianos (github.com/kardianos) as I used some of his code as a refrence for how to use the eventlog package.

Simple Usecase
--------------

Example code of what I am trying to get to:
NOTE:  priority can be one of the following:

- EVENTLOG_ERROR_TYPE
- EVENTLOG_WARNING_TYPE
- EVENTLOG_INFORMATION_TYPE

See the code in the example directory, which should compile on any OS.

bugs
----

as of right now there are none.  Of course I have only used the test code, so I am sure there are bugs all over the place.
