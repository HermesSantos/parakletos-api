package admin_controller

import "github.com/labstack/echo/v5"

func InitAdminController(e *echo.Group) {
	e.GET("/admin", func(c *echo.Context) error {
		return c.JSON(200, "admin")
	})
}
