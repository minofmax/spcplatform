package middlewares

import (
	"github.com/kataras/iris/v12/middleware/accesslog"
	"os"
)

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
