# Logman

[![Coverage Status](https://coveralls.io/repos/github/mishankov/logman/badge.svg)](https://coveralls.io/github/mishankov/logman)
[![CI](https://img.shields.io/github/actions/workflow/status/mishankov/logman/ci.yml)](https://github.com/mishankov/logman/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mishankov/logman)](https://goreportcard.com/report/github.com/mishankov/logman)

---

<p align="center">
  <img src="./images/logo.png" alt="logo" width="200"/>
</p>

<p align="center">
 Deeply configurable logging library for Go with no configuration required to start
</p>

# Installation

```bash
go get -u github.com/mishankov/logman
```

# Usage
## Default logger

For basic log output to stdout, the default logger can be used:

```go
package main

import "github.com/mishankov/logman/loggers"

func main() {
	logger := loggers.NewDefaultLogger()
	logger.Info("Hello, world!")
}
```

Output:

```
[2009-11-10 23:00:00 GMT+0000] [main.main:7] [Info] - Hello, world!
```

## Available logging functions

### Default functions

```go
package main

import (
	"github.com/mishankov/logman"
	"github.com/mishankov/logman/loggers"
)

func main() {
	logger := loggers.NewDefaultLogger()

	logger.Debug("Hello,", "world!")
	logger.Info("Hello,", "world!")
	logger.Warn("Hello,", "world!")
	logger.Error("Hello,", "world!")
	logger.Fatal("Hello,", "world!")

	logger.Log(logman.Info, "Hello,", "world!")
}
```

Output:
```
[2009-11-10 23:00:00 GMT+0000] [main.main:11] [Debug] - Hello, world!
[2009-11-10 23:00:00 GMT+0000] [main.main:12] [Info] - Hello, world!    
[2009-11-10 23:00:00 GMT+0000] [main.main:13] [Warn] - Hello, world!    
[2009-11-10 23:00:00 GMT+0000] [main.main:14] [Error] - Hello, world!   
[2009-11-10 23:00:00 GMT+0000] [main.main:15] [Fatal] - Hello, world!   
[2009-11-10 23:00:00 GMT+0000] [main.main:17] [Info] - Hello, world!
```

### Functions with string formatting

```go
package main

import (
	"github.com/mishankov/logman"
	"github.com/mishankov/logman/loggers"
)

func main() {
	logger := loggers.NewDefaultLogger()

	logger.Debugf("Hello, %s!", "world")
	logger.Infof("Hello, %s!", "world")
	logger.Warnf("Hello, %s!", "world")
	logger.Errorf("Hello, %s!", "world")
	logger.Fatalf("Hello, %s!", "world")

	logger.Logf(logman.Info, "Hello, %s!", "world")
}
```

Output:
```
[2009-11-10 23:00:00 GMT+0000] [main.main:11] [Debug] - Hello, world!
[2009-11-10 23:00:00 GMT+0000] [main.main:12] [Info] - Hello, world!
[2009-11-10 23:00:00 GMT+0000] [main.main:13] [Warn] - Hello, world!
[2009-11-10 23:00:00 GMT+0000] [main.main:14] [Error] - Hello, world!
[2009-11-10 23:00:00 GMT+0000] [main.main:15] [Fatal] - Hello, world!
[2009-11-10 23:00:00 GMT+0000] [main.main:17] [Info] - Hello, world!
```

### Structured logging

Structured logging functions allow to add key-value pairs to messages

```go
package main

import (
	"github.com/mishankov/logman"
	"github.com/mishankov/logman/loggers"
)

func main() {
	logger := loggers.NewDefaultLogger()

	logger.Debugs("Hello, world!", "key1", "value", "key2", 1234)
	logger.Infos("Hello, world!", "key1", "value", "key2", 1234)
	logger.Warns("Hello, world!", "key1", "value", "key2", 1234)
	logger.Errors("Hello, world!", "key1", "value", "key2", 1234)
	logger.Fatals("Hello, world!", "key1", "value", "key2", 1234)

	logger.Logs(logman.Info, "Hello, world!", "key1", "value", "key2", 1234)
}
```

Output:
```
[2009-11-10 23:00:00 GMT+0000] [main.main:11] [Debug] - Hello, world! key1=value key2=1234
[2009-11-10 23:00:00 GMT+0000] [main.main:12] [Info] - Hello, world! key1=value key2=1234
[2009-11-10 23:00:00 GMT+0000] [main.main:13] [Warn] - Hello, world! key1=value key2=1234
[2009-11-10 23:00:00 GMT+0000] [main.main:14] [Error] - Hello, world! key1=value key2=1234
[2009-11-10 23:00:00 GMT+0000] [main.main:15] [Fatal] - Hello, world! key1=value key2=1234
[2009-11-10 23:00:00 GMT+0000] [main.main:17] [Info] - Hello, world! key1=value key2=1234
```

## Custom logger

You can use `logman.NewLogger` to create a custom logger. For example, this is how to mimic `loggers.NewDefaultLogger` using `logman.NewLogger`:

```go
package main

import (
	"os"
	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
)

func main() {
	logger := logman.NewLogger(os.Stdout, formatters.NewDefaultFormatter(formatters.DefaultFormat, formatters.DefaultTimeLayout), nil)
}
```

## Writers

Writers output logs to some destination. Every writer should implement the `io.Writer` interface. Logman provides `FileWriter`:

```go
package main

import "github.com/mishankov/logman/writers"

func main() {
	fw, _ := writers.NewFileWriter("test.log")
}
```

`DefaultLogger` uses `os.Stdout` as writer.

## Formatters

Formatters format log messages before they are passed to the writer. Every formatter should implement the `logman.Formatter` interface. Logman provides `DefaultFormatter` and `JSONFormatter`:

```go
package main

import "github.com/mishankov/logman/formatters"

func main() {
	formatter := formatters.NewDefaultFormatter(formatters.DefaultFormat, formatters.DefaultTimeLayout)
	jsonFormatter := formatters.NewJSONFormatter()
}
```

## Filters

Filters filter log messages. Every filter should implement the `logman.Filter` interface. Logman provides `LevelFilter`:

```go
package main

import "github.com/mishankov/logman/filters"

func main() {
	filter := filters.NewLevelFilter(logman.Info)
}
```

## Examples

Let's see how to create a custom logger that outputs Error or higher level messages in JSON format to a file:

```go
package main

import (
	"github.com/mishankov/logman"
	"github.com/mishankov/logman/filters"
	"github.com/mishankov/logman/formatters"
	"github.com/mishankov/logman/writers"
)

func main() {
	fw, _ := writers.NewFileWriter("error.log")
	formatter := formatters.NewJSONFormatter()
	filter := filters.NewLevelFilter(logman.Error)
	logger := logman.NewLogger(fw, formatter, filter)

	logger.Error("Hello,", "world!")
	logger.Debug("I am not logged")
}
```

`error.log` content:
```json
{"callLocation":"main.main:16","logLevel":"Error","message":"Hello, world!","time":"2009-11-10 23:00:00 GMT+0000"}
```

# Motivation

- Practice Golang and TDD skills
- Simple but flexible logging for personal projects
