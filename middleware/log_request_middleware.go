package middleware

import (
	"merchant-payment-api/config"
	"merchant-payment-api/model"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogRequestMiddleware(log *logrus.Logger) gin.HandlerFunc {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	startTime := time.Now()

	return func(c *gin.Context) {
		c.Next()
		endTime := time.Since(startTime)
		requestLog := model.RequestLog{
			StartTime:  startTime,
			EndTime:    endTime,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			UserAgent:  c.Request.UserAgent(),
		}
		switch {
		case c.Writer.Status() >= 500:
			log.Error(requestLog)
			break
		case c.Writer.Status() >= 400:
			log.Warn(requestLog)
			break
		default:
			log.Info(requestLog)
			break
		}
	}
}
