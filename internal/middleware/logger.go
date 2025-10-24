package middleware

import (
	"time"
	"user-management-api/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		logger.Log.Info("HTTP request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("ip", c.IP()),
			zap.String("request_id", c.Locals("requestID").(string)),
		)

		return err
	}
}
