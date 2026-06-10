package router

import (
	"github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func InitAdminRoutes(c *echo.Group) {
	ar := c.Group("/admin")

	ar.GET("/login", func(c *echo.Context) error {
		return c.JSON(200, "admin")
	})

	ar.Use(echojwt.JWT([]byte("bf6c063e-7ac2-445a-bc15-6c3590e78838")))

}
