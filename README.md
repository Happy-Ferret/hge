# WARNING

This API still needs input. There may still be breakage. A stable legacy API is provided for those wanting to use the C like API.

# Haaf's Game Engine (HGE) in Go!

======================

This is a binding and port of HGE (currently the hge-unix version) to Go. It currently provides a binding to the main C++ HGE class via C binding of hge-unix. The helper classes available in C++ are then ported to Go using just the core binding. Rather than binding everything through the C interface (or through SWIG), there's a fairly sizable chunk of ported/reimplemented code from C++ to Go. In the future, there may be direct bindings to the C interface, but that will be separate from the Go implementation, and functionally, should be identical.

## Requirements:
* Currently, hge-unix is all that's supported. In addition, the c_api branch must be compiled and built with -DBUILD_C_API=ON from https://github.com/losinggeneration/hge-unix/
** As a note, if it's installed to /usr/local you may need to put /usr/local/lib/pkgconfig in your PKG_CONFIG_PATH shell variable.
* You'll need a working Go 1 compiler with cgo. (I've only tested with gc Go.)
* Likely UNIX as I'm unsure if hge-unix compiles & works with Windows.

## Building:
* All you need to do is run: go build
* If you're wanting to use go get, you can do so with: go get github.com/losinggeneration/hge
* Additionally, you can: import "github.com/losinggeneration/hge" and it should work as expected.
