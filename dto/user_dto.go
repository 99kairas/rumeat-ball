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
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}
type UserResponse struct {
	ID    uuid.UUID `json:"user_id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

type AdminRequest struct {
	Email    string `json:"email" form:"email"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}
type AdminRespose struct {
	ID    uuid.UUID `json:"user_id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

type UserProfileResponse struct {
	ID           uuid.UUID `json:"user_id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Status       string    `json:"status"`
	ProfileImage string    `json:"profile_image"`
}

func ConvertToUserModel(user UserRequest) models.User {
	return models.User{
		ID:       uuid.New(),
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
	}
}

func ConvertToAdminModel(admin AdminRequest) models.User {
	return models.User{
		ID:       uuid.New(),
		Email:    admin.Email,
		Name:     admin.Name,
		Password: admin.Password,
		Address:  admin.Address,
		Phone:    admin.Phone,
	}
}

type LoginResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
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

func ConvertToUserProfileResponse(user models.User) UserProfileResponse {
	return UserProfileResponse{
		ID:           user.ID,
		Email:        user.Email,
		Name:         user.Name,
		Phone:        user.Phone,
		Address:      user.Address,
		Status:       user.Status,
		ProfileImage: user.ProfileImage,
	}
}

// USER UPDATE PROFILE

type UserUpdateRequest struct {
	Address      string `json:"address" form:"address"`
	Phone        string `json:"phone" form:"phone"`
	ProfileImage string `json:"profile_image" form:"profile_image"`
}

func ConvertToUpdateUserProfileModel(user UserUpdateRequest) models.User {
	return models.User{
		Address:      user.Address,
		Phone:        user.Phone,
		ProfileImage: user.ProfileImage,
	}
}

func ConvertToUpdateUserProfileResponse(user models.User) UserProfileResponse {
	return UserProfileResponse{
		ID:           user.ID,
		Email:        user.Email,
		Name:         user.Name,
		Phone:        user.Phone,
		Address:      user.Address,
		Status:       user.Status,
		ProfileImage: user.ProfileImage,
	}
}
