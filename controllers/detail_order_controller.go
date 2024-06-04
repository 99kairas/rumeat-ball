package controllers

import (
	"net/http"
	"rumeat-ball/dto"
	m "rumeat-ball/middlewares"
	"rumeat-ball/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateDetailOrderController membuat detail order baru
func CreateDetailOrderController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permision denied: user is not valid",
		})
	}

	var detailOrderReq = dto.DetailOrderRequest{}
	errBind := c.Bind(&detailOrderReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}

	detailOrderData := dto.ConvertToDetailOrderModel(detailOrderReq)
	detailOrderData.ID = userID

	detailOrder, err := repositories.CreateDetailOrder(detailOrderData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error create detail order",
			Response: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success create detail order",
		Response: detailOrder,
	})
}
