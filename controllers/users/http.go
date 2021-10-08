package users

// import (
// 	"net/http"
// 	"social_media/business/users"
// 	"social_media/controllers"
// 	"social_media/controllers/users/request"
// 	"social_media/controllers/users/responses"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type UserControllers struct {
// 	UserUsecase users.Usecase
// }

// func NewController(UserUsecase users.Usecase) *UserControllers {
// 	return &UserControllers{
// 		UserUsecase: UserUsecase,
// 	}
// }

// func (UserController UserControllers) Login(c echo.Context) error {
// 	userLogin := request.Login{}
// 	c.Bind(&userLogin)
// 	ctx := c.Request().Context()
// 	result, err := UserController.UserUsecase.Login(ctx, userLogin.Email, userLogin.Password)
// 	if err != nil {
// 		if err.Error() == "password wrong" {
// 			return controllers.NewFailResponse(c, http.StatusForbidden, err)
// 		} else if err.Error() == "email empety" || err.Error() == "password empety" {
// 			return controllers.NewFailResponse(c, http.StatusBadRequest, err)
// 		}
// 		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controllers.NewSuccesResponse(c, responses.FromLoginResponse(result))

// }

// func (UserController UserControllers) Register(c echo.Context) error {
// 	userRegister := request.UserRegister{}
// 	c.Bind(&userRegister)
// 	ctx := c.Request().Context()
// 	result, err := UserController.UserUsecase.Register(ctx, userRegister.Name, userRegister.Email, userRegister.Address, userRegister.Password)
// 	if err != nil {
// 		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controllers.NewSuccesResponse(c, responses.FromDomain(result))

// }
// func (UserController UserControllers) StatusUser(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	userStatus := request.StatusController{}
// 	c.Bind(&userStatus)
// 	ctx := c.Request().Context()
// 	result, err := UserController.UserUsecase.CreateStatus(ctx, userStatus.Status, id)
// 	if err != nil {
// 		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controllers.NewSuccesResponse(c, responses.FromDomainStatus(result))
// }

// func (UserController UserControllers) UserDeteil(c echo.Context) error {
// 	userDeteil := request.UserDeteil{}

// 	ctx := c.Request().Context()
// 	result, err := UserController.UserUsecase.UserDeteil(ctx,userDeteil.Name,userDeteil.Address, userDeteil.Status)
// 	if err != nil {
// 		return controllers.NewFailResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controllers.NewSuccesResponse(c, responses.FromDomain(result))

// }
