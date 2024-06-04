package repositories

import (
	"errors"
	"rumeat-ball/configs"
	"rumeat-ball/database"
	"rumeat-ball/middlewares"
	"rumeat-ball/models"

	"golang.org/x/crypto/bcrypt"
)

func CheckUser(email string, password string) (models.User, string, error) {
	var data models.User

	tx := database.DB.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return models.User{}, "", errors.New("invalid email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
	if err != nil {
		return models.User{}, "", errors.New("invalid email or password")
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(data.ID, configs.ROLE_USER, data.Email)
		if errToken != nil {
			return models.User{}, "", errToken
		}
	}
	return data, token, nil
}

func CheckAdmin(email string, password string) (models.User, string, error) {
	var data models.User

	tx := database.DB.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return models.User{}, "", errors.New("invalid email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
	if err != nil {
		return models.User{}, "", errors.New("invalid email or password")
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(data.ID, configs.ROLE_ADMIN, data.Email)
		if errToken != nil {
			return models.User{}, "", errToken
		}
	}
	return data, token, nil
}

func CreateUser(data models.User) (models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	data.Password = string(hashPassword)

	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return data, nil
}

func CheckUserEmail(email string) bool {
	var data models.User

	tx := database.DB.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return false
	}

	return true
}

func SetVerificationOTP(email, otp string) error {
	tx := database.DB.Model(&models.User{}).Where("email = ?", email).Update("OTP", otp)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func SetOTP(email, otp string) error {
	if !CheckUserEmail(email) {
		return errors.New("user email not found")
	}

	tx := database.DB.Model(&models.User{}).Where("email = ?", email).Update("OTP", otp)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func ValidateOTP(email, otp string) (models.User, error) {
	var data models.User

	tx := database.DB.Where("email = ? AND otp = ?", email, otp).First(&data)
	if tx.Error != nil {
		return models.User{}, errors.New("invalid Email or OTP")
	}

	database.DB.Model(&models.User{}).Where("email = ?", email).Update("OTP", nil).Update("Status", "verified")
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	return data, nil
}
