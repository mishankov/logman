# Logman

[![codecov](https://codecov.io/github/mishankov/logman/graph/badge.svg?token=3KHQU1BLMV)](https://codecov.io/github/mishankov/logman)
[![CI](https://github.com/mishankov/logman/actions/workflows/ci.yml/badge.svg)](https://github.com/mishankov/logman/actions/workflows/ci.yml)

---

<p align="center">
  <img src="./images/logo.jpg" alt="logo" width="200"/>


<p align="center">
 Highly configurable logging library for Go with sane defaults
</p>

# Installation

```bash
go get -u github.com/mishankov/logman
```

# Usage
## Default logger

For basic log output to stdout default logger can be used:

```go
package main

import "github.com/mishankov/logman/loggers"

logger := loggers.NewDefaultLogger()

logger.Info("Hello, world!")
```

## Available logging functions

```go
package main

import "github.com/mishankov/logman"
import "github.com/mishankov/logman/loggers"

logger := loggers.NewDefaultLogger()

logger.Debug("Hello,", "world!")
logger.Debugf("Hello, %s!", "world")

logger.Info("Hello,", "world!")
logger.Infof("Hello, %s!", "world")

logger.Warn("Hello,", "world!")
logger.Warnf("Hello, %s!", "world")

logger.Error("Hello,", "world!")
logger.Errorf("Hello, %s!", "world")

logger.Fatal("Hello,", "world!")
logger.Fatalf("Hello, %s!", "world")

logger.Log(logmnan.Info, "Hello,", "world!")
logger.Logf(logman.Info, "Hello, %s!", "world")

```

## Custom logger

You can use `logman.NewLogger` to create custom logger. For example, this how to create mimic `loggers.NewDefaultLogger` using `logman.NewLogger`:

```go
package main

import "github.com/mishankov/logman"
import "github.com/mishankov/logman/formatters"

logger := logman.NewLogger(os.Stdout, formatters.NewDefaultFormatter(formatters.DefaultFormat, formatters.DefaultTimeLayout), nil)
```

## Writers

Writers output logs to some destination. Every writer should implement `io.Writer` interface. Logman provides `FileWriter`:

```go
package main

import "github.com/mishankov/logman"
import "github.com/mishankov/logman/writers"

fw, _ := writers.NewFileWriter("test.log")
```

`DefaultLogger` uses `os.Stdout` as writer

## Formatters

Formatters format log messages before they passed to writer. Every formatter should implement `logman.Formatter` interface. Logman provides `DefaultFormatter` and `JSONFormatter`:

```go
package main

import "github.com/mishankov/logman/formatters"

formatter := formatters.NewDefaultFormatter(formatters.DefaultFormat, formatters.DefaultTimeLayout)

jsonFormatter := formatters.NewJSONFormatter()
```

## Filters

Filters filter log messages. Every filter should implement `logman.Filter` interface. Logman provides `LevelFilter`:

```go
package main

import "github.com/mishankov/logman/filters"

filter := filters.NewLevelFilter(logman.Info)
```

## Examples

Let's see how to create custom logger that outputs Error or higher level messages in JSON format to file:

```go
package main

import "github.com/mishankov/logman"
import "github.com/mishankov/logman/formatters"
import "github.com/mishankov/logman/writers"
import "github.com/mishankov/logman/filters"

func main() {
  fw, _ := writers.NewFileWriter("error.log")
  formatter := formatters.NewJSONFormatter()
  filter := filters.NewLevelFilter(logman.Error)
  logger := logman.NewLogger(fw, formatter, filter)

  logger.Error("Hello,", "world!")
  logger.Debug("I am not logged")
}
```
