package xlog

import (
	"context"
	"log/slog"
)

type HookFunc func(ctx context.Context, r slog.Record) slog.Record

type HookHandler struct {
	handler slog.Handler

	hooks []HookFunc
}

func NewHookHandler(handler slog.Handler, hooks ...HookFunc) *HookHandler {
	return &HookHandler{
		handler: handler,
		hooks:   hooks,
	}
}

func (h *HookHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *HookHandler) Handle(ctx context.Context, r slog.Record) error {
	// 调用所有钩子
	for _, hook := range h.hooks {
		r = hook(ctx, r)
	}
	// 传递给底层 handler
	return h.handler.Handle(ctx, r)
}

func (h *HookHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewHookHandler(h.handler.WithAttrs(attrs), h.hooks...)
}

func (h *HookHandler) WithGroup(name string) slog.Handler {
	return NewHookHandler(h.handler.WithGroup(name), h.hooks...)
}
