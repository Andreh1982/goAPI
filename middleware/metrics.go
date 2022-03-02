package middleware

import (
	"strconv"

	"goAPI/metric"

	"github.com/gin-gonic/gin"
)

func Metrics(mService metric.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		appMetric := metric.NewHTTP(context.FullPath(), context.Request.Method)

		appMetric.Started()

		context.Next()

		appMetric.Finished()

		appMetric.StatusCode = strconv.Itoa(context.Writer.Status())

		mService.SaveHTTP(appMetric)
	}
}
