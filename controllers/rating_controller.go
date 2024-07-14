package controllers

import (
	"net/http"
	"rumeat-ball/dto"
	m "rumeat-ball/middlewares"
	"rumeat-ball/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateRatingMenuController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	var ratingReq = dto.CreateRatingMenuRequest{}
	errBind := c.Bind(&ratingReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}

	ratingData := dto.ConvertToCreateRatingModel(ratingReq, userID)
	ratingData.ID = userID

	rating, err := repositories.CreateRating(ratingData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error create rating",
			Response: err.Error(),
		})
	}

	ratingResponse := dto.ConvertToCreateRatingResponse(rating)

	return c.JSON(http.StatusCreated, dto.Response{
		Message:  "success create rating",
		Response: ratingResponse,
	})

}

func GetAllRatingsController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	ratings, err := repositories.GetAllRatings(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error get all ratings",
			Response: err.Error(),
		})
	}

	ratingResponse := dto.ConvertToGetAllRatingsResponse(ratings)

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success get all ratings",
		Response: ratingResponse,
	})
}

func UpdateRatingMenuController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	ratingID, _ := uuid.Parse(c.Param("id"))
	if ratingID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error parse id",
			Response: "error parse id",
		})
	}

	var ratingReq = dto.UpdateRatingMenuRequest{}
	errBind := c.Bind(&ratingReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}

	ratingData := dto.ConvertToUpdateRatingModel(ratingReq, userID)
	ratingData.ID = userID

	rating, err := repositories.UpdateRating(ratingData, ratingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error update rating",
			Response: err.Error(),
		})
	}

	ratingResponse := dto.ConvertToUpdateRatingResponse(rating)

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success update rating",
		Response: ratingResponse,
	})
}

func DeleteRatingMenuController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	ratingID, _ := uuid.Parse(c.Param("id"))
	if ratingID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error parse id",
			Response: "error parse id",
		})
	}

	err := repositories.DeleteRating(ratingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error delete rating",
			Response: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success delete rating",
		Response: nil,
	})
}
