# Golog

Golog is a simple log for golang.

## Features

- log levels: DEBUG, INFO, WARN, ERROR
- deferent log handlers: ConsoleHander, FileHandler, RotatingHandler
- basic on golang inner pkg - log

## Installation

    go get github.com/jander/golog/logger

## Example

    // rotating hander, max log files is 4, max file size is 4M
    rotatingHandler := logger.NewRotatingHandler("./", "test.log", 4, 4*1024*1024)

    // logger set handlers: console, rotating
    logger.SetHandlers(logger.Console, rotatingHandler)

    defer logger.Close()

    // logger set flags
    logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

    // logger set log level
    logger.SetLevel(logger.INFO)

    logger.Debug("something", "debug")
    logger.Info("something")
    logger.Warn("something")
    logger.Error("something")

## License

Golog is available under the MIT License.

