#Safe Alarm

##Introduction

Simple program intended to be run on a raspberry pi to detect when a sensor is open or closed. It will send an email on open, or if left open for a configurable period of time.

It will also send a heartbeat at a configurable time.

## Messages

Event messages can be sent to email addresses using the [mailjet](https://www.mailjet.com/). All messages will be sent to all emails specified in the config file.

To send text messages to verizon customers, set email to `<phone number>@@vtext.com`.

By default (i.e. for free), mailjet rate limits to 10/hour, 200/day, 6000/month.  Sending to multiple emails does not count more against this limit.  Purely the number of times opened or left open. The 10/hour can be lifted for free by verifying your identity.

##Config

Configuration is done via a YAML file passed in as the first command line arguement.

##Build

```bash
# cd to the project dir
GOOS=linux GOARCH=arm GOARM=5 go build
```

The build artifact will be called `safealarm` and will be in the current directory. SCP it to the raspberry pi.

Build instructions came from example [here](https://www.thepolyglotdeveloper.com/2017/04/cross-compiling-golang-applications-raspberry-pi/).

##Run

```bash
./safealarm config.yaml
```

Note that `sudo` is not required currently, but it would be in order for edge detection to work.  If omitted it will freeze the whole pi on first edge detected and require a hard restart of the raspberry pi.

