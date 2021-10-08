package user

import (
	"net/http"
	"social_media/configs"
	"social_media/models"

	"github.com/labstack/echo/v4"
	// "gorm.io/gorm"
)

func Register(c echo.Context) error {
	var UserRegister models.Register
	c.Bind(&UserRegister)
	var RegisterDB models.UserDatabase
	RegisterDB.Nama = UserRegister.Nama
	RegisterDB.Alamat = UserRegister.Alamat
	RegisterDB.Email = UserRegister.Email
	RegisterDB.Umur = UserRegister.Umur
	RegisterDB.Password = UserRegister.Password
	Result := configs.DB.Create(&RegisterDB)
	if Result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseeResponse{
			Code:    http.StatusInternalServerError,
			Message: "Server Sedang Error",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseeResponse{
		Code:    http.StatusOK,
		Message: "Succes Register",
		Data:    UserRegister,
	})
}
