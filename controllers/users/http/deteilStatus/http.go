package deteilStatus

import (
	"net/http"
	"social_media/business/users/statusdesc"
	"social_media/controllers"
	"social_media/controllers/users/request"
	"social_media/controllers/users/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllers struct {
	UserUsecase statusdesc.Usecase
}

func NewController(UserUsecase statusdesc.Usecase) *UserControllers {
	return &UserControllers{
		UserUsecase: UserUsecase,
	}
}

func (UserController UserControllers) MakingComment(c echo.Context) error {
	userComent := request.RequestComent{}
	id, _ := strconv.Atoi(c.Param("id"))
	idstatus, _ := strconv.Atoi(c.Param("idstatus"))
	c.Bind(&userComent)
	ctx := c.Request().Context()
	result, err := UserController.UserUsecase.MakingComment(ctx, userComent.Coment, id, idstatus)
	if err != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomainComment(result))
}

func (UserController UserControllers) MakingLike(c echo.Context) error {
	userLike := request.RequestComent{}
	id, _ := strconv.Atoi(c.Param("id"))
	idstatus, _ := strconv.Atoi(c.Param("idstatus"))
	c.Bind(&userLike)
	ctx := c.Request().Context()
	result, err := UserController.UserUsecase.MakingLike(ctx, id, idstatus)
	if err != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.SuccesLike(c, responses.FromDomainComment(result))
}
