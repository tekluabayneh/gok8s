package utils

import (
	"fmt"
	"log/slog"
	"os"
)

var log *slog.Logger

func InitLog(debug, jsonFormat bool) {
	loglevel := slog.LevelInfo

	if debug {
		loglevel = slog.LevelDebug
	}

	otps := &slog.HandlerOptions{
		Level:     loglevel,
		AddSource: debug,
	}

	var handler slog.Handler

	if jsonFormat {
		handler = slog.NewJSONHandler(os.Stdout, otps)
	} else {
		handler = slog.NewTextHandler(os.Stdout, otps)
	}

	log = slog.New(handler)
	slog.SetDefault(log)

	fmt.Println("initialized Logs")
}

func Log() *slog.Logger {
	return log
}
