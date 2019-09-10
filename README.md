# Safe Alarm

## Introduction

Simple program intended to be run on a raspberry pi to detect when a sensor is open or closed. It will send an email on open, or if left open for a configurable period of time.

It will also send a heartbeat at a configurable time.

I installed the sensor on the door to a Safe, but it certainly could be applied to other applications.

Why is this project written in Go? Well mostly because I wanted to try out a new language.  This is my first Go program so be Gentle ;)

## Messages

Event messages can be sent to email addresses using [mailjet](https://www.mailjet.com/). All messages will be sent to all emails specified in the config file.

By default (i.e. for free), mailjet rate limits to 10/hour, 200/day, 6000/month.  Sending to multiple emails in a single message still counts as multiple against the rate limit. The 10/hour can be lifted for free by verifying your identity.

### Text Messages

To send text messages to Verizon customers, set email to `<phone number>@vtext.com`. https://www.verizon.com/about/news/vzw/2013/06/computer-to-phone-text-messaging

To send text messages to Google Fi customers, set email to `<phone number>@msg.fi.google.com`. https://support.google.com/fi/answer/6356597

## Config

Configuration is done via a YAML file passed in as the first command line arguement.

## Test

Run unit tests

```bash
go test
```

Since so much of this program is based on waiting for things to happend, the unit tests take some time to run.

## Build

```bash
# cd to the project dir
GOOS=linux GOARCH=arm GOARM=5 go build
```

The build artifact will be called `safealarm` and will be in the current directory. SCP it to the raspberry pi.

## Run

```bash
./safealarm config.yaml
```

Note that `sudo` is not required currently, but it would be in order for edge detection to work.  If omitted it will freeze the whole pi on first edge detected and require a hard restart of the raspberry pi.

## Hardware

This code was originally written for the Raspberry Pi Model A and [Gikfun MC-38 Wired Door Sensor Magnetic Switch](https://www.amazon.com/gp/product/B0154PTDFI)

## Reference

Go Resources
- Basics of building and running go: https://golang.org/doc/code.html
- Go CLI: https://golang.org/cmd/go/
- Language Primer: https://golang.org/doc/effective_go.html
- Base libraries: https://golang.org/pkg/
- Excellent Language Reference: https://www.golang-book.com/books/intro
- Go Modules: https://github.com/golang/go/wiki/Modules#how-to-install-and-activate-module-support

Raspberry Pi
- Library used to access GPIO from Go: https://github.com/stianeikeland/go-rpio
- Build instructions to cross compile Go for the Pi: https://www.thepolyglotdeveloper.com/2017/04/cross-compiling-golang-applications-raspberry-pi/
