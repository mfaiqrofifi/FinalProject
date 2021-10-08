package userRegister

import (
	"net/http"
	"social_media/business/users/userRegister"
	"social_media/controllers"
	"social_media/controllers/users/request"
	"social_media/controllers/users/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllers struct {
	UserUsecase userRegister.Usecase
}

func NewController(UserUsecase userRegister.Usecase) *UserControllers {
	return &UserControllers{
		UserUsecase: UserUsecase,
	}
}

func (UserController UserControllers) Login(c echo.Context) error {
	userLogin := request.Login{}
	c.Bind(&userLogin)
	ctx := c.Request().Context()
	result, err := UserController.UserUsecase.Login(ctx, userLogin.Email, userLogin.Password)
	if err != nil {
		if err.Error() == "password wrong" {
			return controllers.NewFailResponse(c, http.StatusForbidden, err)
		} else if err.Error() == "email empety" || err.Error() == "password empety" {
			return controllers.NewFailResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromLoginResponse(result))

}

func (UserController UserControllers) Register(c echo.Context) error {
	userRegister := request.UserRegister{}
	c.Bind(&userRegister)
	ctx := c.Request().Context()
	result, err := UserController.UserUsecase.Register(ctx, userRegister.Name, userRegister.Email, userRegister.Address, userRegister.Password)
	if err != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(result))

}

func (userControllers UserControllers) DeteilRegister(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	user, error := userControllers.UserUsecase.DeteilRegister(ctx, id)
	if error != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, error)
	}
	return controllers.NewSuccesResponse(c, user)
}

func (userControllers UserControllers) GetFriend(c echo.Context) error {
	nama := c.QueryParam("nama")
	ctx := c.Request().Context()
	users, error := userControllers.UserUsecase.GetFriend(ctx, nama)
	if error != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, error)
	}
	response := make([]responses.Response, len(users))
	for i, user := range users {
		response[i] = responses.FromDomain(user)
	}
	return controllers.NewSuccesResponse(c, response)
}
