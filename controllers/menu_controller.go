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
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
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

	// Ensure category exists
	if _, err := repositories.GetCategoryByID(menuReq.CategoryID); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "invalid category ID",
			Response: "category not found",
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

	return c.JSON(http.StatusCreated, dto.Response{
		Message:  "success create menu",
		Response: menuResponse,
	})
}

func GetMenuController(c echo.Context) error {
	menuName := c.FormValue("name")
	categoryIDStr := c.FormValue("category_id")
	var categoryID uuid.UUID
	if categoryIDStr != "" {
		var err error
		categoryID, err = uuid.Parse(categoryIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.Response{
				Message:  "invalid category ID",
				Response: err.Error(),
			})
		}
	}

	data, err := repositories.GetMenu(menuName, categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error get menu",
			Response: err.Error(),
		})
	}

	commentCounts, averageRatings, err := repositories.GetAllCommentCountsAndAverageRatings()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error get comment counts and average ratings",
			Response: err.Error(),
		})
	}

	menuData := dto.ConvertToGetAllMenuResponse(data, commentCounts, averageRatings)

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success get menu",
		Response: menuData,
	})
}

func GetMenuByIDController(c echo.Context) error {
	menuID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	data, err := repositories.GetMenuByID(menuID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error get menu",
			Response: err.Error(),
		})
	}

	commentCount, averageRating, err := repositories.GetCommentCountAndAverageRating(menuID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error get comment count and average rating",
			Response: err.Error(),
		})
	}

	comments, err := repositories.GetCommentsByMenuID(menuID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error get comments",
			Response: err.Error(),
		})
	}

	commentResponses := dto.ConvertToGetAllRatingsResponseDetailsMenu(comments)
	menuData := dto.ConvertToGetMenuResponse(data, commentCount, averageRating)
	menuData.Comments = commentResponses // Set data komentar

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success get menu",
		Response: menuData,
	})
}

func UpdateMenuController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	menuID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	var menuReq = dto.UpdateMenuRequest{}
	errBind := c.Bind(&menuReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}

	// Ensure category exists
	if _, err := repositories.GetCategoryByID(menuReq.CategoryID); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "invalid category ID",
			Response: "category not found",
		})
	}

	menuData := dto.ConvertToUpdateMenuModel(menuReq)
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
	responseData, err := repositories.UpdateMenu(menuData, menuID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error update menu",
			Response: err.Error(),
		})
	}
	menuResponse := dto.ConvertToUpdateMenuResponse(responseData, menuID)
	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success update menu",
		Response: menuResponse,
	})
}

func DeleteMenuController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	menuID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	err = repositories.DeleteMenu(menuID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error delete menu",
			Response: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Response{
		Message: "success delete menu",
	})
}
