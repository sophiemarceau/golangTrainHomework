package imitateHystrix

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Wrapper(size, reqThreshold int, failThreshold float64, duration time.Duration) gin.HandlerFunc {
	s := InitSlideWindow(size, reqThreshold, failThreshold, duration)
	s.Launch()
	s.Monitor()
	s.ShowStatus()
	return func(context *gin.Context) {
		if s.Broken() {
			context.String(http.StatusBadRequest, " reject! Fuck")
			context.Abort()
			return
		}
		context.Next()
		if context.Writer.Status() != http.StatusOK {
			s.RecordReqResult(false)
		} else {
			s.RecordReqResult(true)
		}
	}
}
