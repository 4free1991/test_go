package logger

import (
	"bytes"
	"context"
	"fmt"
	filename "github.com/keepeye/logrus-filename"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/petermattis/goid"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var Logger = logrus.New()

func getTraceId(ctx context.Context, key string) string {
	u, _ := ctx.Value(key).(string)
	return u
}

//=========>>>

func Init() {
	filenameHook := filename.NewHook()
	Logger.AddHook(filenameHook)
	Logger.SetFormatter(&MyFormatter{})
	Logger.SetReportCaller(true) //定位行号
	logFile := "logs/"
	if !Exists(logFile) {
		_ = os.MkdirAll(logFile, os.ModePerm)
	}
	logFileName := logFile + "/info.log"
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  writer(logFileName, 7),
		logrus.WarnLevel:  writer(logFileName, 7),
		logrus.ErrorLevel: writer(logFileName, 7),
		logrus.PanicLevel: writer(logFileName, 7),
	}, &MyFormatter{})

	Logger.AddHook(lfHook)

	//err log 配置
	//增加错误日志输出
	errLogFileName := logFile + "/error.log"
	errFile, _ := os.OpenFile(path.Join(logFile, errLogFileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errlfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.ErrorLevel: errFile,
	}, &MyFormatter{})
	Logger.AddHook(errlfHook)
}

/*
*
文件设置
*/
func writer(logPath string, save uint) *rotatelogs.RotateLogs {
	newPath := logPath[:len(logPath)-4]
	fileSuffix := "%Y-%m-%d.log"
	logier, err := rotatelogs.New(
		newPath+"_"+fileSuffix,
		rotatelogs.WithLinkName(logPath),          // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(save),        // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return logier
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

type MyFormatter struct {
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	timestamp = replaceAtPosition(timestamp, 20, ",")
	location := "xxx.go"
	function := "method"
	if entry.Caller != nil {
		location = afterLast(entry.Caller.File, "/") + ":" + strconv.Itoa(entry.Caller.Line)
		function = entry.Caller.Function
	}
	//Logger.WithContext(ctx).Infof("==>eeee")
	var trace, span string
	if ctx := entry.Context; ctx != nil {
		trace = getTraceId(ctx, "trace")
		span = getTraceId(ctx, "span")
	}
	var goId = goid.Get()
	var newLog string
	newLog = fmt.Sprintf("%s [%d] %s [%s] [%s] [trace=%s,span=%s,parent=] - %s\n", timestamp,
		goId,
		strings.ToUpper(entry.Level.String()), function, location, trace, span, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func replaceAtPosition(originaltext string, indexofcharacter int, replacement string) string {
	runes := []rune(originaltext)
	partOne := string(runes[0 : indexofcharacter-1])
	partTwo := string(runes[indexofcharacter:])
	return partOne + replacement + partTwo
}

func afterLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[i+len(char):]
}
