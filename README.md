cafe-bot
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Downloads][15]][16]

[1]: https://godoc.org/github.com/eure/cafe-bot?status.svg
[2]: https://godoc.org/github.com/eure/cafe-bot
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/eure/cafe-bot.svg
[6]: https://github.com/eure/cafe-bot/releases/latest
[7]: https://travis-ci.org/eure/cafe-bot.svg?branch=master
[8]: https://travis-ci.org/eure/cafe-bot
[9]: https://coveralls.io/repos/eure/cafe-bot/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/eure/cafe-bot?branch=master
[11]: https://codecov.io/github/eure/cafe-bot/coverage.svg?branch=master
[12]: https://codecov.io/github/eure/cafe-bot?branch=master
[13]: https://goreportcard.com/badge/github.com/eure/cafe-bot
[14]: https://goreportcard.com/report/github.com/eure/cafe-bot
[15]: https://img.shields.io/github/downloads/eure/cafe-bot/total.svg?maxAge=1800
[16]: https://github.com/eure/cafe-bot/releases
[17]: https://img.shields.io/github/stars/eure/cafe-bot.svg
[18]: https://github.com/eure/cafe-bot/stargazers


Slack Bot for Archimedes Cafe @ eureka, Inc.


# Install

```bash
$ go get -u github.com/eure/cafe-bot
```

# Build

```bash
$ cd /path/to/cafe-bot
$ go build -o cafebot
```

for Raspberry Pi

```bash
$ GOOS=linux GOARCH=arm GOARM=6 go build -o cafebot
```

# Run

```bash
SLACK_BOT_TOKEN=xoxb-0000... GOOGLE_HOME_HOST=192.168.0.1 ./cafebot
```

## Environment variables

|Name|Description|
|:--|:--|
| `SLACK_RTM_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_BOT_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_BOT_SPEECH` | Flag for speech feature. Set [boolean like value](https://golang.org/pkg/strconv/#ParseBool). |
| `SLACK_BOT_DEBUG` | Flag for debug logging. Set [boolean like value](https://golang.org/pkg/strconv/#ParseBool). |
| `GOOGLE_HOME_HOST` | Hostname or IP address of Google Home for speech feature. |
| `GOOGLE_HOME_PORT` | Port number of Google Home. Default is `8009`. |

# Development

## Directory Structure

```
# Entry point
├── main.go

# SlackBot Daemon and Config
├── bot.go
├── config.go

# Third party clients. (Slack, Google Home)
├── clients.go

# Helper methods for Command
├── command__data.go
├── command__factory.go
├── command__helper.go

# Original Commands
├── command_xxx.go
├── command_yyy.go
├── command_zzz.go
...

# Tasks.
├── task.go
├── task_xxx.go

# Cafe
├── cafe_menu.go    # menu
├── cafe_order.go   # order history
```

## How to create other commands?

`Command` is composed of multiple `task`.
Create new command and add it on `CreateCommand` function.

1. Copy `command_xxx.go`

2. Edit it for new command and adds existing (or new) task.

3. Edit `CreateCommand` function on `command__factory.go`, and add new `case` for the new command.

## Supported tasks

- Send Slack message
- Send Slack message as a thread reply
- Speak something on Google Home
- Reload SlackBot
- Change flags of SlackBot
