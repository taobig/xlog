package xlog

import (
	"io"
	"log/slog"
)

func SetUpSlog(w io.Writer, level slog.Level, hooks ...HookFunc) (*slog.Logger, *slog.LevelVar) {
	var levelVar slog.LevelVar
	levelVar.Set(level)
	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: &levelVar,
		//ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		//	if a.Key == slog.TimeKey && len(groups) == 0 {
		//		// 修改时间格式
		//		a.Value = slog.StringValue(a.Value.Time().Format(time.RFC3339Nano))
		//	}
		//	return a
		//},
	})
	newLogger := slog.New(NewHookHandler(handler, hooks...))

	return newLogger, &levelVar
}
