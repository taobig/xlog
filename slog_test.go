package xlog

import (
	"log/slog"
	"os"
	"testing"
)

func TestSetUpSlog(t *testing.T) {
	slog.Debug("debug log 1")
	slog.Info("info log 1")

	logger, levelVar := SetUpSlog(os.Stdout, slog.LevelInfo)
	logger.Debug("debug log 2")
	logger.Info("info log 2")

	slog.SetDefault(logger)
	levelVar.Set(slog.LevelDebug)
	slog.Debug("debug log 3")
	slog.Info("info log 3")
}
