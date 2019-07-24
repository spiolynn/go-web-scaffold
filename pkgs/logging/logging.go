package logging

/*
日志模块
*/

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

var (
	hostname string
	Logs     *logrus.Entry
	Logm     *logrus.Entry
)

func init() {
	fmt.Println(" init logging ...")
	Setting()
}

func Setting() {

	hname, err := os.Hostname()
	if err != nil {
		hostname = "template"
	} else {
		hostname = hname
	}

	// init logpath
	logpath := filepath.Join(setting.G_cfg_yaml.Logs.Logsrootpath,
		setting.G_cfg_yaml.Logs.Logsfilename)
	Logs = newLogger(logpath)

	// init Metric Logpath
	if setting.G_cfg_yaml.Logs.Metriclogsfilename != "empty" ||
		len(setting.G_cfg_yaml.Logs.Metriclogsfilename) != 0 {
		fmt.Println("--- need LogM ---")
		metriclogpath := filepath.Join(setting.G_cfg_yaml.Logs.Logsrootpath,
			setting.G_cfg_yaml.Logs.Metriclogsfilename)
		Logm = newLogger(metriclogpath)
	}

	// 设置日志格式
	//Log.Formatter = &logrus.JSONFormatter{}
	//Log.Formatter = &logrus.TextFormatter{}
	//logfomat()

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	// Log.SetOutput(os.Stdout)

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

//// 封装 info
//func Info(args ...interface{}) {
//	Log.WithFields(logrus.Fields{
//		"host": hostname,
//	}).Info(args)
//}
//
//func Warn(args ...interface{}) {
//	Log.WithFields(logrus.Fields{
//		"host": hostname,
//	}).Warn(args)
//}
//
//func Fatal(args ...interface{}) {
//	Log.WithFields(logrus.Fields{
//		"host": hostname,
//	}).Fatal(args)
//}
//
//func Debug(args ...interface{}) {
//	Log.WithFields(logrus.Fields{
//		"host": hostname,
//	}).Debug(args)
//}
//
//func Panic(args ...interface{}) {
//	Log.WithFields(logrus.Fields{
//		"host": hostname,
//	}).Panic(args)
//}

// 定义日志输出格式

//func logfomat() {
//
//	Log = &logrus.Logger{
//		Out:   os.Stderr,
//		Level: logrus.DebugLevel,
//		Formatter: &easy.Formatter{
//			TimestampFormat: "2006-01-02 15:04:05",
//			LogFormat:       "[%lvl%]: %time% - %msg% \n",
//		},
//	}
//
//	customFormatter := new(logrus.TextFormatter)
//	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
//
//}

//func main() {
//
//	Info("i am info log", "sdfadf")
//	Warn("i am info log")
//	Debug("i am debug log")
//	//Panic("i am panic")
//	//Fatal("i am panic")
//
//}
