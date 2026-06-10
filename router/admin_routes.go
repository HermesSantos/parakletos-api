package router

import "github.com/labstack/echo/v5"

func InitAdminRoutes(c *echo.Group) {
	ar := c.Group("/admin")

	ar.GET("/admin", func(c *echo.Context) error {
		return c.JSON(200, "admin")
	})
}
