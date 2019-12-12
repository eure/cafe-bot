cafebot
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Co decov Coverage][11]][12] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22] [![Downloads][15]][16]

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
[19]: https://codeclimate.com/github/eure/cafe-bot/badges/gpa.svg
[20]: https://codeclimate.com/github/eure/cafe-bot
[21]: https://bettercodehub.com/edge/badge/eure/cafe-bot?branch=master
[22]: https://bettercodehub.com/



Slack Bot for Archimedes Cafe @ eureka, Inc.


# Install

```bash
$ go get -u github.com/eure/cafe-bot
```

# Build

```bash
$ make build
```

for Raspberry Pi

```bash
$ make build-arm6
```

# Run

```bash
SLACK_BOT_TOKEN=xoxb-0000... GOOGLE_HOME_HOST=192.168.0.1 GOOGLE_HOME_LANG=ja ./bin/cafebot
```

## Environment variables

|Name|Description|
|:--|:--|
| `SLACK_RTM_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_BOT_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `BOBO_DEBUG` | Flag for debug logging. Set [boolean like value](https://golang.org/pkg/strconv/#ParseBool). |
| `GOOGLE_HOME_HOST` | Hostname or IP address of Google Home for speech feature. |
| `GOOGLE_HOME_PORT` | Port number of Google Home. Default is `8009`. |
| `GOOGLE_HOME_LANG` | Speaking language of Google Home. Default is `en`. |
| `GOOGLE_HOME_ACCENT` | Speaking accent of Google Home. Default is `us`. |


## Supported Commands

- Order drink
- Show order history
