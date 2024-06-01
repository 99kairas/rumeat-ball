package controllers

import (
	"net/http"
	"rumeat-ball/dto"
	m "rumeat-ball/middlewares"
	"rumeat-ball/repositories"
	"rumeat-ball/util"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateMenuController(c echo.Context) error {
	menu := m.ExtractTokenUserId(c)
	if menu == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permision denied: user is not valid",
		})
	}

	var menuReq = dto.CreateMenuRequest{}
	errBind := c.Bind(&menuReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})

	}

	menuData := dto.ConvertToCreateMenuModel(menuReq)
	menuImage, err := c.FormFile("image")
	if err != http.ErrMissingFile {
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.Response{
				Message:  "error upload menu image",
				Response: err.Error(),
			})

		}

		menuURL, err := util.UploadToCloudinary(menuImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.Response{
				Message:  "error upload menu image to Cloudinary",
				Response: err.Error(),
			})

		}
		menuData.Image = menuURL
	}

	responseData, err := repositories.CreateMenu(menuData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error create menu",
			Response: err.Error(),
		})

	}

	menuResponse := dto.ConvertToCreateMenuResponse(responseData)

	return c.JSON(http.StatusOK, dto.Response{

		Message:  "success create menu",
		Response: menuResponse,
	})

}
