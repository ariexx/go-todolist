package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// LogMiddleware Handle log middleware
func LogMiddleware() func(ctx *fiber.Ctx) error {
	//call log service
	return func(ctx *fiber.Ctx) error {
		formatTime := time.Now().Format("2006-01-02")
		log := logrus.New()
		log.SetFormatter(&logrus.JSONFormatter{})

		requestID := ctx.Get("X-Request-Id")
		method := ctx.Method()
		url := ctx.OriginalURL()
		ip := ctx.IP()

		//create file log
		file, err := os.OpenFile("logs/requests/request-"+formatTime+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Log(logrus.ErrorLevel, err)
		}
		log.SetOutput(file)

		log.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     method,
			"url":        url,
			"ip":         ip,
		}).Info("request received")
		return ctx.Next()
	}
}
