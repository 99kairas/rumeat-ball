package controllers

func CreateOrderController(c echo.Context) error {
	var payloads = dto.OrderRequest{}
	errBind := c.Bind(&payloads)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}
	data, err := repositories.CreateOrder(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "failed create order",
			Response: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success create order",
		Response: data,
	})
}