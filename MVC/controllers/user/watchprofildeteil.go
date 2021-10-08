package user

import (
	"net/http"
	"social_media/configs"
	"social_media/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func WatchDeteilProfil(e echo.Context) error {
	// id, _ := strconv.Atoi(e.Param("id"))
	user := []models.UserDatabase{}

	result := configs.DB.Find(&user)
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
		Data:    user,
	})

}
