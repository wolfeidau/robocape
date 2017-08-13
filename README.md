# robocape

[![GoDoc](https://godoc.org/github.com/wolfeidau/robocape?status.svg)](https://godoc.org/github.com/wolfeidau/robocape)

This is a Go wrapper for the [robotcape C library](https://github.com/StrawsonDesign/Robotics_Cape_Installer) which is designed to work with the [beaglebone blue](https://beagleboard.org/blue) board. It is designed to make it easy to take advantage of this great library without rewriting it in pure Go.

# usage

As one would imagine this library must be built on the native platform to enable linking to robotcape library using CGO.

```
debian@beaglebone:~$ go get -u -v github.com/wolfeidau/robocape
...
debian@beaglebone:~$ cd go/src/github.com/wolfeidau/robocape
debian@beaglebone:~$ go install ./examples/...
debian@beaglebone:~$ sudo /home/debian/go/bin/rc_buttons
```

# license

This code is released under MIT license and is Copyright Mark Wolfe.