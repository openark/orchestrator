## shared-stopwatch

shared-stopwatch - a simple package to provide shared timer functionality.

There are several implementations of a stopwatch in Go, but I haven't found
any that supports sharing the stopwatch between many routines running in parallel.

This is intended for collecting total run time of a task executed (possibly) in parallel
inside an application.

### Installation

Install the package with:
`go get github.com/maurosr/shared-stopwatch`

### Contributing

Patches and improvements to this package are welcome.

### Licensing

BSD 2-Clause License

### Feedback

Feedback and patches welcome.

Mauro Schilman
<mauro.schilman@booking.com>

### Code Documentation

[godoc.org/github.com/maurosr/shared-stopwatch](http://godoc.org/github.com/maurosr/shared-stopwatch)

### Acknowledgement

This software was originally developed at Booking.com. With approval from Booking.com, this software was released as open source, for which the authors would like to express their gratitude.
