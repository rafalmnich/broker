# IQCC Broker - API server for sending mqtt actions for home automation

[![Go Report Card](https://goreportcard.com/badge/github.com/rafalmnich/broker)](https://goreportcard.com/report/github.com/rafalmnich/broker)
[![Build Status](https://travis-ci.org/rafalmnich/broker.svg?branch=master)](https://travis-ci.org/rafalmnich/broker)
[![Coverage Status](https://coveralls.io/repos/github/rafalmnich/broker/badge.svg?branch=master)](https://coveralls.io/github/rafalmnich/broker?branch=master)
[![GoDoc](https://godoc.org/github.com/rafalmnich/broker?status.svg)](https://godoc.org/github.com/rafalmnich/broker)
[![GitHub release](https://img.shields.io/github/release/rafalmnich/broker.svg)](https://github.com/rafalmnich/broker/releases)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/rafalmnich/broker/master/LICENSE)

## Starting up the server:

```shell script
go build -o broker ./cmd/broker
```

## Sending messages through API
### endpoint: `/action`
### message format:
```json
[
  {
    "name": "some name",
    "value": 213
  },
  {
    "name": "other name",
    "value": 10
  }
]
```



#### usefull shortcuts for later use:
for turning denon off: http://192.168.8.135:8080/goform/formiPhoneAppPower.xml?1+PowerOff
