go-timber
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22]

[1]: https://godoc.org/github.com/evalphobia/go-timber?status.svg
[2]: https://godoc.org/github.com/evalphobia/go-timber
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/go-timber.svg
[6]: https://github.com/evalphobia/go-timber/releases/latest
[7]: https://travis-ci.org/evalphobia/go-timber.svg?branch=master
[8]: https://travis-ci.org/evalphobia/go-timber
[9]: https://coveralls.io/repos/evalphobia/go-timber/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/go-timber?branch=master
[11]: https://codecov.io/github/evalphobia/go-timber/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/go-timber?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/go-timber
[14]: https://goreportcard.com/report/github.com/evalphobia/go-timber
[15]: https://img.shields.io/github/downloads/evalphobia/go-timber/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/go-timber/releases
[17]: https://img.shields.io/github/stars/evalphobia/go-timber.svg
[18]: https://github.com/evalphobia/go-timber/stargazers
[19]: https://codeclimate.com/github/evalphobia/go-timber/badges/gpa.svg
[20]: https://codeclimate.com/github/evalphobia/go-timber
[21]: https://bettercodehub.com/edge/badge/evalphobia/go-timber?branch=master
[22]: https://bettercodehub.com/


Unofficial golang library for Timber.io.


# Quick Usage

```go
import (
	"github.com/evalphobia/go-timber/timber"
)

func someFunction() {
	conf := timber.Config{
		APIKey:       "",
		SourceID:     "",
		Environment:  "production",
		MinimumLevel: timber.LogLevelInfo,
		Sync:         false,
		Debug:        true,
	}

	cli := timber.New(conf)

	cli.Debug("logging...")
	cli.Trace("logging...")
	cli.Info("logging...")
	cli.Warn("logging...")
	cli.Err("logging...")
	cli.Fatal("logging...")
}

```


# Environment variables

| Name | Description |
|:--|:--|
| `TIMBER_API_KEY` | API Key of timber. |
| `TIMBER_SOURCE_ID` | Source ID of timber. |
