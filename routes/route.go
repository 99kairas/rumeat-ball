package routes

import (
	"rumeat-ball/configs"
	"rumeat-ball/controllers"
	"rumeat-ball/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Trailing Slash for slashing in endpoint
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	// USERS
	e.POST("/users/signup", controllers.SignUpUserController)
	e.PUT("/users/verify", controllers.ValidateOTP)
	e.POST("/users/login", controllers.LoginUserController)

	// ADMIN
	e.POST("/admin/login", controllers.AdminLoginController)
	admin := e.Group("/admins")
	admin.Use(middleware.JWT([]byte(configs.JWT_KEY)))
	admin.POST("/signup", controllers.AdminSignUpController, middlewares.CheckRole(configs.ROLE_ADMIN))

	return e
}
