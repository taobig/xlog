package xlog

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestNewLogrusHook(t *testing.T) {
	// 创建钩子函数
	hookHandlers := []HookFunc{
		func(ctx context.Context, r slog.Record) slog.Record {
			// 添加额外字段
			r.Add("request_id", time.Now().UnixNano())
			r.Add("timestamp", slog.TimeValue(r.Time))

			return r
		},
		func(ctx context.Context, r slog.Record) slog.Record {
			// 过滤或处理特定级别的日志
			if r.Level >= slog.LevelError {
				// 发送错误通知等
			}
			return r
		},
		func(ctx context.Context, r slog.Record) slog.Record {
			{
				// --- 组装 JSON 并发送到外部平台 ---
				logData := map[string]any{
					"time":  r.Time.Format(time.RFC3339Nano),
					"level": r.Level.String(),
					"msg":   r.Message,
				}

				// 把 slog 的 Attrs 也转成 JSON
				r.Attrs(func(a slog.Attr) bool {
					logData[a.Key] = a.Value.Any()
					return true
				})

				payload, _ := json.Marshal(logData)
				//_, err := os.Stdout.Write(payload)
				//require.NoError(t, err)
				t.Logf("payload: %s", payload)
			}
			return r
		},
	}

	// 创建一个可变的 LevelVar
	var levelVar slog.LevelVar
	//levelVar.Set(slog.LevelDebug) // 初始级别设为 Debug，缺省为info

	// 创建带钩子的 handler
	handler := NewHookHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: &levelVar}), hookHandlers...)
	logger := slog.New(handler)
	logger.Debug("这是一条带钩子的debug日志")
	logger.Info("这是一条带钩子的info日志")

	levelVar.Set(slog.LevelDebug)

	logger.Debug("debug log 1")
	logger.Info("info log 1", "user", "alice")
	logger.Warn("err log 1", "user", "bob")
}
