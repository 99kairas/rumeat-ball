package controllers

import (
	"net/http"
	"rumeat-ball/dto"
	m "rumeat-ball/middlewares"
	"rumeat-ball/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateCategoryController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	var categoryReq = dto.CreateCategoryRequest{}
	errBind := c.Bind(&categoryReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}

	categoryData := dto.ConvertToCreateCategoryModel(categoryReq)
	responseData, err := repositories.CreateCategory(categoryData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error create category",
			Response: err.Error(),
		})
	}

	categoryResponse := dto.ConvertToCreateCategoryResponse(responseData)

	return c.JSON(http.StatusCreated, dto.Response{
		Message:  "success create category",
		Response: categoryResponse,
	})
}

func GetCategoriesController(c echo.Context) error {
	data, err := repositories.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error get categories",
			Response: err.Error(),
		})
	}

	categoryData := []dto.GetCategoryResponse{}
	for _, category := range data {
		categoryData = append(categoryData, dto.ConvertToGetCategoryResponse(category))
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success get categories",
		Response: categoryData,
	})
}

func UpdateCategoryController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	var categoryReq = dto.UpdateCategoryRequest{}
	errBind := c.Bind(&categoryReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}

	categoryData := dto.ConvertToUpdateCategoryModel(categoryReq)
	responseData, err := repositories.UpdateCategory(categoryID, categoryData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error update category",
			Response: err.Error(),
		})
	}

	categoryResponse := dto.ConvertToUpdateCategoryResponse(responseData, categoryID)

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success update category",
		Response: categoryResponse,
	})
}

func DeleteCategoryController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	err = repositories.DeleteCategory(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error delete category",
			Response: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "success delete category",
	})
}

func GetCategoryByIDController(c echo.Context) error {
	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	data, err := repositories.GetCategoryByID(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error get category",
			Response: err.Error(),
		})
	}

	categoryData := dto.ConvertToGetCategoryResponse(data)
	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success get category",
		Response: categoryData,
	})
}
