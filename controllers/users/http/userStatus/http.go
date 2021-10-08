package userStatus

import (
	"net/http"
	"social_media/business/users/userStatus"
	"social_media/controllers"
	"social_media/controllers/users/request"
	"social_media/controllers/users/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllers struct {
	UserUsecase userStatus.Usecase
}

func NewController(UserUsecase userStatus.Usecase) *UserControllers {
	return &UserControllers{
		UserUsecase: UserUsecase,
	}
}

func (UserController UserControllers) StatusUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userStatus := request.StatusController{}
	c.Bind(&userStatus)
	ctx := c.Request().Context()
	result, err := UserController.UserUsecase.CreateStatus(ctx, userStatus.Status, id)
	if err != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomainStatus(result))
}
func (userControllers UserControllers) SeeComment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("idstatus"))
	ctx := c.Request().Context()
	user, error := userControllers.UserUsecase.SeeComment(ctx, id)
	if error != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, error)
	}
	return controllers.NewSuccesResponse(c, user)
}
