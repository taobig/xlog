package xlog

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestHook(t *testing.T) {
	logger := logrus.New()
	myHook := NewLogrusHook(logrus.AllLevels, func(entry *logrus.Entry) {
		var bufferData []byte
		var err error
		if entry.Buffer != nil {
			bufferData, err = io.ReadAll(entry.Buffer)
			assert.NoError(t, err)
		}
		t.Logf("myHook called, time:%+v, level:%+v, message:%+v, fields:%+v, buffer:%+v",
			entry.Time.Format(time.RFC3339), entry.Level, entry.Message, entry.Data, string(bufferData))
	})
	logger.AddHook(myHook)
	logger.Info("should call hook")
	logger.Error("should call hook")
	logger.WithField("key", "value").Error("should call hook")
	logger.Error("should call hook")
}
