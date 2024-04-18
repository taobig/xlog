package xlog

import (
	"github.com/sirupsen/logrus"
)

func SetUp(reportCaller bool) {
	SetUpWithLogger(logrus.StandardLogger(), reportCaller)
}

func SetUpWithLogger(logger *logrus.Logger, reportCaller bool) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}
	//logger := logrus.New()
	logger.SetReportCaller(reportCaller) //设置在输出日志中添加文件名和方法信息。默认显示的是长文件名，函数名和行号
	//logger.SetOutput(os.Stdout)
	//logger.SetFormatter(&logrus.JSONFormatter{})
	//logger.SetFormatter(&logrus.JSONFormatter{
	//	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
	//		// origin: {"file":"/Users/xxx/xlog/logrus_test.go:17","func":"github.com/taobig/xlog.TestLogrus","level":"error","msg":"error log","time":"2024-04-13T16:00:00Z"}
	//		// origin: {"file":"/Users/xxx/xlog/logrus_test.go:17","func":"github.com/taobig/xlog.(*parent).Test","level":"error","msg":"error log","time":"2024-04-13T16:00:00Z"}
	//		//i := strings.LastIndex(frame.Function, ".")
	//		i := strings.LastIndex(frame.Function, "/")
	//		str := frame.Function[i+1:]
	//		j := strings.Index(str, ".")
	//		shortFuncName := str[j+1:]
	//		//_, fileName := path.Split(frame.File)
	//		//filepath := frame.Function[:i] + "/" + fileName
	//		//return shortFuncName, fmt.Sprintf("%s:%d", filepath, frame.Line)
	//		return shortFuncName, fmt.Sprintf("%s:%d", frame.File, frame.Line)
	//	},
	//})
	logger.SetFormatter(&SortedFieldJSONFormatter{})
}
