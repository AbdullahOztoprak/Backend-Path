package observability

import (
    "os"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func InitLogger() {
    // Set the log level based on the environment variable
    logLevel := os.Getenv("LOG_LEVEL")
    if logLevel == "" {
        logLevel = "info"
    }

    zerolog.SetGlobalLevel(parseLogLevel(logLevel))

    // Configure the logger to output to console
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func parseLogLevel(level string) zerolog.Level {
    switch level {
    case "debug":
        return zerolog.DebugLevel
    case "info":
        return zerolog.InfoLevel
    case "warn":
        return zerolog.WarnLevel
    case "error":
        return zerolog.ErrorLevel
    case "fatal":
        return zerolog.FatalLevel
    case "panic":
        return zerolog.PanicLevel
    default:
        return zerolog.InfoLevel
    }
}