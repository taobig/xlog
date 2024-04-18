package xlog

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"time"
)

type SortedFieldJSONFormatter struct {
	sortedFields struct {
		Time  string `json:"time"`
		Level string `json:"level"`
		File  string `json:"file,omitempty"`
		Func  string `json:"func,omitempty"`
		Msg   string `json:"msg"`
	}
}

func (f *SortedFieldJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var file string
	var shortFuncName string
	if entry.HasCaller() {
		var frame *runtime.Frame = entry.Caller
		i := strings.LastIndex(frame.Function, "/")
		str := frame.Function[i+1:]
		j := strings.Index(str, ".")
		shortFuncName = str[j+1:]
		file = fmt.Sprintf("%s:%d", frame.File, frame.Line)
	}
	f.sortedFields.Time = entry.Time.Format(time.RFC3339)
	f.sortedFields.Level = entry.Level.String()
	f.sortedFields.File = file
	f.sortedFields.Func = shortFuncName
	f.sortedFields.Msg = entry.Message

	serialized, err := json.Marshal(f.sortedFields)
	if err != nil {
		return nil, err
	}

	return append(serialized, '\n'), nil
}
