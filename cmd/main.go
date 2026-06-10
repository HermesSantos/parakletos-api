package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	en := echo.New()
	en.Use(middleware.RequestLogger())
	en.Use(middleware.ContextTimeout(60 * time.Second))
	en.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20.0)))
	en.Pre(middleware.RemoveTrailingSlash())

	e := en.Group("/api")
	e.Use(middleware.RequestLogger())

	// testing route
	e.GET("/ping", func(c *echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	if err := en.Start(":8081"); err != nil {
		en.Logger.Error("failed to start server", "error", err)
	}
}
