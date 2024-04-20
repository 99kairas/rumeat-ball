package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"rumeat-ball/configs"
	"rumeat-ball/dto"
	"rumeat-ball/repositories"
	"rumeat-ball/templates"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func isStrongPassword(password string) bool {
	var (
		hasUppercase bool
		hasSymbol    bool
		hasMinLength bool
	)

	// Cek minimal panjang password
	if len(password) >= 8 {
		hasMinLength = true
	}

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUppercase = true
		}

		symbols := "!@#$%^&*()-_=+{};:,<.>/?"
		if strings.ContainsRune(symbols, char) {
			hasSymbol = true
		}

		if hasMinLength && hasUppercase && hasSymbol {
			return true
		}
	}

	return false
}

func SignUpUserController(c echo.Context) error {
	var payloads = dto.UserRequest{}
	errBind := c.Bind(&payloads)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}

	validEmail := govalidator.IsEmail(payloads.Email)
	if !validEmail {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "fail sign up",
			Response: "email is not valid",
		})
	}

	// Pengecekan apakah password memenuhi kriteria
	if !isStrongPassword(payloads.Password) {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "fail sign up",
			Response: "password must be at least 8 characters and contain at least 1 uppercase letter and 1 symbol",
		})
	}

	emailExist := repositories.CheckUserEmail(payloads.Email)
	if emailExist {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "fail sign up",
			Response: "email already exist",
		})
	}

	signUpData := dto.ConvertToUserModel(payloads)

	data, err := repositories.CreateUser(signUpData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "fail sign up",
			Response: err.Error(),
		})
	}

	// SEND OTP
	otp := fmt.Sprint(rand.Intn(999999-100000) + 100000)

	err = repositories.SetVerificationOTP(payloads.Email, otp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "failed set otp",
			Response: err.Error(),
		})
	}

	emailBody, err := templates.RenderOTPTemplate(otp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "failed render otp template",
			Response: err.Error(),
		})
	}

	err = configs.SendMail(payloads.Email, "MySPP OTP", emailBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "failed send email",
			Response: err.Error(),
		})
	}

	response := dto.ConvertToUserResponse(data)

	return c.JSON(http.StatusCreated, dto.Response{
		Message:  "success sign up",
		Response: response,
	})
}
