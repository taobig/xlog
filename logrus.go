package xlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
)

func SetUpLogrus(logger *logrus.Logger) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}
	//logger := logrus.New()
	logger.SetReportCaller(true) //设置在输出日志中添加文件名和方法信息。默认显示的是长文件名，函数名和行号
	//logger.SetOutput(os.Stdout)
	//logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			i := strings.LastIndex(frame.Function, ".")
			shortFuncName := frame.Function[i+1:]
			filepath := frame.Function[:i]
			_, fileName := path.Split(frame.File)
			filepath = filepath + "/" + fileName
			return shortFuncName, fmt.Sprintf("%s:%d", filepath, frame.Line)
		},
	})
}
