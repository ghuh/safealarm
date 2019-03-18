#Safe Alarm

##Introduction

Simple program intended to be run on a raspberry pi to detect when a sensor is open or closed. It will send an email on open, or if left open for a configurable period of time.

It will also send a heartbeat at a configurable time.

## Messages

Event messages can be sent to email addresses using the [mailjet](https://www.mailjet.com/). All messages will be sent to all emails specified in the config file.

To send text messages to verizon customers, set email to `<phone number>@@vtext.com`.

By default, mailjet rate limits to 10 messages an hour.  Sending to multiple emails does not count more against this limit.  Purely the number of times openend or left open.

##Config

Configuration is done via a YAML file passed in as the first command line arguement.

##Build

```bash
# cd to the project dir
GOOS=linux GOARCH=arm GOARM=5 go build
```

The build artifact will be called `safealarm` and will be in the current directory. SCP it to the raspberry pi.

##Run

`./safealarm config.yaml`
