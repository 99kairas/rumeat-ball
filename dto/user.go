package dto

import (
	"rumeat-ball/models"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type OTPRequest struct {
	Email string `json:"email" form:"email"`
}

type ValidateOTPRequest struct {
	Email string `json:"email" form:"email"`
	OTP   string `json:"otp" form:"otp"`
}

// type ChangeUserPasswordRequest struct {
// 	ID       uuid.UUID `json:"id" form:"id"`
// 	Password string    `json:"password" form:"password"`
// }

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func ConvertToUserModel(user UserRequest) models.User {
	return models.User{
		ID:       uuid.New(),
		Email:    user.Email,
		Password: user.Password,
	}
}

type LoginResponse struct {
	ID    uuid.UUID `json:"user_id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

// func ConvertToChangeUserPasswordModel(user ChangeUserPasswordRequest) models.User {
// 	return models.User{
// 		ID:       user.ID,
// 		Password: user.Password,
// 	}
// }

func ConvertToUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}
