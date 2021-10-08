package user

import (
	"net/http"
	"social_media/configs"
	"social_media/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginUser(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	afterLogins := []models.UserDatabase{}

	result := configs.DB.Select("nama", "alamat", "umur").Where("id=?", id).Find(&afterLogins)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return e.JSON(http.StatusInternalServerError, models.BaseeResponse{
				Code:    http.StatusInternalServerError,
				Message: "server maaf",
				Data:    nil,
			})
		}
	}
	return e.JSON(http.StatusOK, models.BaseeResponse{
		Code:    http.StatusOK,
		Message: "Berhasil Login",
		Data:    afterLogins,
	})
}
