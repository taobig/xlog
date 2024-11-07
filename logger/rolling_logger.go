package logger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

func NewRollingLogger(logPath string, maxSizeInMegabytes, maxBackups, maxAge int, localTime, compress bool) (*lumberjack.Logger, error) {
	err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create logfile parent dir:%s, err: %+v", filepath.Dir(logPath), err)
	}
	return &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    maxSizeInMegabytes,
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
		// LocalTime determines if the time used for formatting the timestamps in
		// backup files is the computer's local time.  The default is to use UTC
		// time.
		LocalTime: localTime,
		// Compress determines if the rotated log files should be compressed using gzip. The default is not to perform compression.
		Compress: compress, // disabled by default
	}, nil
}
