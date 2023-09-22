package xlog

import "github.com/sirupsen/logrus"

type LogrusHook struct {
	levels   []logrus.Level
	callback func(entry *logrus.Entry)
}

func NewLogrusHook(levels []logrus.Level, callback func(entry *logrus.Entry)) *LogrusHook {
	return &LogrusHook{
		levels:   levels,
		callback: callback,
	}
}

func (h *LogrusHook) Levels() []logrus.Level {
	//return logrus.AllLevels
	return h.levels
}

func (h *LogrusHook) Fire(entry *logrus.Entry) error {
	//if entry.Level == logrus.ErrorLevel {}
	h.callback(entry)
	return nil
}
