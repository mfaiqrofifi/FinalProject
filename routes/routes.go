package routes

import (
	"social_media/controllers/user"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	ev1.POST("register", user.Register)
	ev1.GET("login/:id", user.LoginUser)
	ev1.POST("login/:id/buatstatus", user.BuatStatus)
	return e
}
