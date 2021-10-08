package routes

import (
	"social_media/controllers/users/http/deteilStatus"
	"social_media/controllers/users/http/userRegister"
	"social_media/controllers/users/http/userStatus"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserControllerRegister      userRegister.UserControllers
	UserControllerStatus        userStatus.UserControllers
	UserControllerDeteilStattus deteilStatus.UserControllers
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("user/login", cl.UserControllerRegister.Login)
	e.POST("user/register", cl.UserControllerRegister.Register)
	e.POST("user/status/:id", cl.UserControllerStatus.StatusUser)
	e.GET("user/deteilUser/:id", cl.UserControllerRegister.DeteilRegister)
	e.POST("user/Comment/:id/:idstatus", cl.UserControllerDeteilStattus.MakingComment)
	e.POST("user/Comment/Like/:id/:idstatus", cl.UserControllerDeteilStattus.MakingLike)
	e.GET("user/Status/:idstatus/comment", cl.UserControllerStatus.SeeComment)
	e.GET("user/friend", cl.UserControllerRegister.GetFriend)
	// e.POST("user/addfriend/:id")
	// e.GET("user/profiledetile/:id",cl.UserController.)
}
