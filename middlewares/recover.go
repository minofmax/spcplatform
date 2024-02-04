package middlewares

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"runtime"
	"spcplatform/services/models"
)

func CustomRecover(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.IsStopped() {
				return
			}

			var stacktrace string
			for i := 1; ; i++ {
				_, f, l, got := runtime.Caller(i)
				if !got {
					break
				}
				stacktrace += fmt.Sprintf("%s:%d\n", f, l)
			}

			errMsg := fmt.Sprintf("错误信息: %s", err)
			// when stack finished
			logMessage := fmt.Sprintf("从错误中恢复L ('%s')\n", ctx.HandlerName())
			logMessage += errMsg + "\n"
			logMessage += fmt.Sprintf("\n%s", stacktrace)
			// 打印错误日志
			ctx.Application().Logger().Warn(logMessage)
			// 返回错误信息
			ctx.JSON(models.ResponseBody{Status: models.BackendError})
			ctx.StopExecution()
		}
	}()
	ctx.Next()
}
