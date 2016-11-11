# GOPHOTOBOOTH

[![Build Status](https://travis-ci.org/jbuchbinder/gophotobooth.svg?branch=master)](https://travis-ci.org/jbuchbinder/gophotobooth)

Simple, simple tethered photobooth. Based on [Gin](https://github.com/gin-gonic/gin) and [gphoto2](http://www.gphoto.org/).

_Eventually, this needs to be rewritten using the native [gphoto2go](https://github.com/micahwedemeyer/gphoto2go) Go gphoto2 bindings. For now, it's using the gphoto2 binary._

## Documentation

> This has only been tested with a tethered EOS 40D body, so your results may vary wildly.

"Go Photo Booth" has been designed to run as a micro web server with a tethered gphoto2-enabled camera. It is meant to be triggered by a [foot switch](http://amzn.to/2eOxxvW) which has been configured to produce a "space" (0x20) character.

A gphoto2-compatible camera needs to be tethered to the computer running *gophotobooth*.

> NOTE: If the camera's filesystem is mounted by Nautilus (on Linux machines), make sure that it has been unmounted before attempting camera communication.

