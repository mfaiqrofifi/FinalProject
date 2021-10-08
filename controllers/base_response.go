package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseeResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	Data interface{} `json:"data"`
}

func NewSuccesResponse(c echo.Context, data interface{}) error {
	response := BaseeResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "succes"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func NewFailResponse(c echo.Context, Status int, err error) error {
	response := BaseeResponse{}
	response.Meta.Status = Status
	response.Meta.Message = err.Error()
	response.Data = nil
	return c.JSON(Status, response)
}
func SuccesLike(c echo.Context, data interface{}) error {
	response := BaseeResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "berhasil like"
	response.Data = nil
	return c.JSON(http.StatusOK, response)
}
