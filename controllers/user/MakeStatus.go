package user

import (
	"net/http"
	"social_media/configs"
	"social_media/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BuatStatus(e echo.Context) error {
	var buatstatus models.StatusUser
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&buatstatus)
	var statusDB models.StatusUserDB
	statusDB.Status = buatstatus.Status
	statusDB.IdUser = id
	buatstatus.IdUser = id
	result := configs.DB.Create(&statusDB)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.BaseeResponse{
			Code:    http.StatusInternalServerError,
			Message: "Maaf Sedang Error",
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, models.BaseeResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    buatstatus,
	})
}
