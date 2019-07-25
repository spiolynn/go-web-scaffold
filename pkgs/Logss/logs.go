package main

import (
	"fmt"
	"github.com/keepeye/logrus-filename"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-web-scaffold/pkgs/setting"
	"os"
	"path/filepath"
	"time"
)

/*
需求:
两个logger
*/

var (
	Hostname           string
	Logs               *logrus.Entry
	DefaultCallerDepth = 1
)

func init() {
	fmt.Println(" init logging ...")
	Setting()
}

func Setting() {

	hname, err := os.Hostname()
	if err != nil {
		Hostname = "template"
	} else {
		Hostname = hname
	}

	// init logpath
	logpath := filepath.Join(setting.G_cfg_yaml.Logs.Logsrootpath,
		setting.G_cfg_yaml.Logs.Logsfilename)

	Logs = newLogger(logpath)

}

// 定义日志
func newLogger(v_path string) *logrus.Entry {

	// 日志日期分割
	path := v_path
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	)
	if err != nil {
		fmt.Errorf("日志日期分割 error")
	}

	// 日志类型分割
	pathMap := lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}

	log := logrus.New()

	// 设置日志级别为warn以上
	switch setting.G_cfg_yaml.Logs.Logslevel {
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	// 加行号hook
	filenameHook := filename.NewHook()
	filenameHook.Field = "lineno"
	log.AddHook(filenameHook)

	// 文件hook
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.DisableTimestamp = false
	customFormatter.DisableLevelTruncation = true

	log.Hooks.Add(lfshook.NewHook(pathMap, customFormatter))

	//logs := log.WithFields(logrus.Fields{"host": Hostname})
	logs := log.WithFields(logrus.Fields{})

	return logs
}

// 封装 info
//func Info(args ...interface{}) {
//
//	_, file, line, ok := runtime.Caller(1)
//	if ok {
//		lineinfo := fmt.Sprintf("[%s:%-d]", file, line)
//		Log.WithFields(logrus.Fields{
//			"host": Hostname, "line": lineinfo,
//		}).Info(args)
//	}else{
//		Log.WithFields(logrus.Fields{
//			"host": Hostname,
//		}).Info(args)
//	}
//
//}

// 定义日志输出格式
//
//func logfomat() {
//	Log = &logrus.Logger{
//		Out:   os.Stderr,
//		Level: logrus.DebugLevel,
//		Formatter: &easy.Formatter{
//			TimestampFormat: "2006-01-02 15:04:05",
//			LogFormat:       "[%lvl%]: %time% - %msg% \n",
//		},
//	}
//	//customFormatter := new(logrus.TextFormatter)
//	//customFormatter.TimestampFormat = "2006-01-02 15:04:05"
//}

func main() {

	Logs.Info("sdfadf")
	Logs.Debug("sdfadf")
	Logs.Warn("sdfsdf")
}
