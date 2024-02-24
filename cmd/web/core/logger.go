package core

import (
	"log/slog"
	"os"
)

func setLogOptions(isDev bool) *slog.HandlerOptions {
	ho := &slog.HandlerOptions{
		AddSource: true,
	}
	if isDev {
		ho.Level = slog.LevelDebug
	} else {
		ho.Level = slog.LevelInfo
	}
	return ho
}

func Logger(dev bool) *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, setLogOptions(dev)))
}
