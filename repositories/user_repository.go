package repositories

import (
	"errors"
	"fmt"
	"rumeat-ball/configs"
	"rumeat-ball/database"
	"rumeat-ball/middlewares"
	"rumeat-ball/models"

	"github.com/google/uuid"
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

func ResetPassword(email, password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	password = string(hashPassword)

	tx := database.DB.Model(&models.User{}).Where("email = ?", email).Update("password", password)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetUserProfile(userID uuid.UUID) (models.User, error) {
	var user models.User
	tx := database.DB.Where("id = ?", userID).First(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}

func UpdateUserProfile(userID uuid.UUID, updatedData models.User) (models.User, error) {
	var user models.User
	// Ambil data pengguna yang ada
	tx := database.DB.Where("id = ?", userID).First(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	// Perbarui field hanya jika ada nilai baru yang diberikan
	if updatedData.Name != "" {
		user.Name = updatedData.Name
	}
	if updatedData.Password != "" {
		user.Password = updatedData.Password
	}
	if updatedData.Address != "" {
		user.Address = updatedData.Address
	}
	if updatedData.Phone != "" {
		user.Phone = updatedData.Phone
	}
	if updatedData.ProfileImage != "" {
		user.ProfileImage = updatedData.ProfileImage
	}

	// Simpan perubahan
	tx = database.DB.Save(&user)
	return user, tx.Error
}

func DeleteUserProfile(userID uuid.UUID) error {
	var user models.User
	tx := database.DB.Model(&user).Where("id = ?", userID).Delete(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ChangePassword(userID uuid.UUID, oldPassword, newPassword string) error {
	var user models.User
	tx := database.DB.Where("id = ?", userID).First(&user)
	if tx.Error != nil {
		return tx.Error
	}

	// Verify old password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return fmt.Errorf("old password is incorrect")
	}

	// Encrypt new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password in the database
	tx = database.DB.Model(&user).Update("password", string(hashedPassword))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
