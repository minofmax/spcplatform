package middlewares

import (
	"bufio"
	rotatelogs "github.com/iproj/file-rotatelogs"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"os"
	"time"
)

func MakeRotateAccessLog() *accesslog.AccessLog {
	pathToAccessLog := "logs/request.log.%Y%m%d%H"
	logs, err := rotatelogs.New(pathToAccessLog, rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(1*time.Hour))
	if err != nil {
		panic(err)
	}
	ac := accesslog.New(bufio.NewWriter(logs))
	ac.AddOutput(os.Stdout)

	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	ac.SetFormatter(&accesslog.JSON{
		Indent:    "",
		HumanTime: true,
	})

	return ac
}

func MakeAccessLog() *accesslog.AccessLog {
	ac := accesslog.File("logs/access.log")
	ac.AddOutput(os.Stdout)

	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	ac.SetFormatter(&accesslog.JSON{
		Indent:    "",
		HumanTime: true,
	})

	return ac
}
