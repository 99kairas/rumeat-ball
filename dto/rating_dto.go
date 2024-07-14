package dto

import (
	"rumeat-ball/models"
	"time"

	"github.com/google/uuid"
)

type CreateRatingMenuRequest struct {
	UserID  uuid.UUID `json:"user_id" form:"user_id"`
	MenuID  uuid.UUID `json:"menu_id" form:"menu_id"`
	Comment string    `json:"comment" form:"comment"`
	Rating  float64   `json:"rating" form:"rating"`
	Date    time.Time `json:"date" form:"date"`
}

type RatingResponse struct {
	ID      uuid.UUID `json:"id" form:"id"`
	MenuID  uuid.UUID `json:"menu_id" form:"menu_id"`
	Menu    GetMenuResponse
	UserID  uuid.UUID `json:"user_id" form:"user_id"`
	User    UserProfileResponse
	Comment string  `json:"comment" form:"comment"`
	Rating  float64 `json:"rating" form:"rating"`
	Date    string  `json:"date" form:"date"`
}

type CreateRatingMenuResponse struct {
	ID      uuid.UUID `json:"id" form:"id"`
	MenuID  uuid.UUID `json:"menu_id" form:"menu_id"`
	UserID  uuid.UUID `json:"user_id" form:"user_id"`
	Comment string    `json:"comment" form:"comment"`
	Rating  float64   `json:"rating" form:"rating"`
	Date    string    `json:"date" form:"date"`
}

type UpdateRatingMenuRequest struct {
	ID      uuid.UUID `json:"id" form:"id"`
	Comment string    `json:"comment" form:"comment"`
	Rating  float64   `json:"rating" form:"rating"`
	Date    time.Time `json:"date" form:"date"`
}

type UpdateRatingMenuResponse struct {
	ID      uuid.UUID `json:"id" form:"id"`
	UserID  uuid.UUID `json:"user_id" form:"user_id"`
	MenuID  uuid.UUID `json:"menu_id" form:"menu_id"`
	User    UserProfileResponse
	Menu    GetMenuResponse
	Comment string  `json:"comment" form:"comment"`
	Rating  float64 `json:"rating" form:"rating"`
	Date    string  `json:"date" form:"date"`
}

func ConvertToCreateRatingModel(ratingReq CreateRatingMenuRequest, userID uuid.UUID) models.Rating {
	return models.Rating{
		UserID:  userID,
		MenuID:  ratingReq.MenuID,
		Comment: ratingReq.Comment,
		Rating:  ratingReq.Rating,
		Date:    time.Now(),
	}
}

func ConvertToCreateRatingResponse(rating models.Rating) CreateRatingMenuResponse {
	dateFormat := rating.Date.Format("02 January 2006 15:04")
	return CreateRatingMenuResponse{
		ID:      rating.ID,
		MenuID:  rating.MenuID,
		UserID:  rating.UserID,
		Comment: rating.Comment,
		Rating:  rating.Rating,
		Date:    dateFormat,
	}
}

func ConvertToGetAllRatingsResponse(ratings []models.Rating) []RatingResponse {
	var ratingResponses []RatingResponse
	for _, rating := range ratings {
		dateFormat := rating.Date.Format("02 January 2006 15:04")
		ratingResponses = append(ratingResponses, RatingResponse{
			ID:      rating.ID,
			MenuID:  rating.MenuID,
			Menu:    ConvertToGetMenuResponse(rating.Menu),
			UserID:  rating.UserID,
			User:    ConvertToUserProfileResponse(rating.User),
			Comment: rating.Comment,
			Rating:  rating.Rating,
			Date:    dateFormat,
		})
	}
	return ratingResponses
}

func ConvertToUpdateRatingModel(ratingReq UpdateRatingMenuRequest, userID uuid.UUID) models.Rating {
	return models.Rating{
		ID:      userID,
		Comment: ratingReq.Comment,
		Rating:  ratingReq.Rating,
		Date:    time.Now(),
	}
}

func ConvertToUpdateRatingResponse(rating models.Rating) UpdateRatingMenuResponse {
	dateFormat := rating.Date.Format("02 January 2006 15:04")
	return UpdateRatingMenuResponse{
		ID:      rating.ID,
		UserID:  rating.UserID,
		MenuID:  rating.MenuID,
		User:    ConvertToUserProfileResponse(rating.User),
		Menu:    ConvertToGetMenuResponse(rating.Menu),
		Comment: rating.Comment,
		Rating:  rating.Rating,
		Date:    dateFormat,
	}
}
