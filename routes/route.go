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
	e.POST("/users/resend-otp", controllers.ResendOTPController)
	user := e.Group("/users")
	user.Use(middleware.JWT([]byte(configs.JWT_KEY)))

	// USERS ORDER
	user.POST("/order", controllers.CreateOrderController, middlewares.CheckRole(configs.ROLE_USER))
	user.GET("/order", controllers.GetAllOrdersController, middlewares.CheckRole(configs.ROLE_USER))
	user.GET("/order/:id", controllers.GetOrderByIDController, middlewares.CheckRole(configs.ROLE_USER))
	user.PUT("/order/cancel/:id", controllers.CancelOrderController, middlewares.CheckRole(configs.ROLE_USER))
	user.PUT("/order/update/:id", controllers.UpdateOrderController, middlewares.CheckRole(configs.ROLE_USER))
	user.PUT("/order/:id", controllers.UserUpdateOrderStatusController, middlewares.CheckRole(configs.ROLE_USER))
	// e.DELETE("/order/:id", controllers.DeleteOrderController)

	// USERS MENU
	e.GET("/menu", controllers.GetMenuController)
	e.GET("/menu/:id", controllers.GetMenuByIDController)

	// USERS CATEGORY MENU
	e.GET("/category", controllers.GetCategoriesController)
	e.GET("/category/:id", controllers.GetCategoryByIDController)

	// USERS TRANSACTION
	user.POST("/transaction", controllers.CreateTransactionController, middlewares.CheckRole(configs.ROLE_USER))
	e.POST("/midtrans/notification", controllers.HandleMidTransNotificationController)

	// USERS RATINGS
	user.POST("/ratings", controllers.CreateRatingMenuController, middlewares.CheckRole(configs.ROLE_USER))
	user.GET("/ratings", controllers.GetAllRatingsController, middlewares.CheckRole(configs.ROLE_USER))
	user.PUT("/ratings/:id", controllers.UpdateRatingMenuController, middlewares.CheckRole(configs.ROLE_USER))
	user.DELETE("/ratings/:id", controllers.DeleteRatingMenuController, middlewares.CheckRole(configs.ROLE_USER))

	// USERS PROFILE
	user.GET("/profile", controllers.GetUserProfileController, middlewares.CheckRole(configs.ROLE_USER))
	user.PUT("/profile", controllers.UpdateUserProfileController, middlewares.CheckRole(configs.ROLE_USER))
	user.DELETE("/profile", controllers.DeleteUserProfileController, middlewares.CheckRole(configs.ROLE_USER))
	user.PUT("/profile/change-password", controllers.ChangePasswordController, middlewares.CheckRole(configs.ROLE_USER))

	// ---------------------------------------------------------------------------------------------------------------------------------------------------------

	// ADMIN
	e.POST("/admin/login", controllers.AdminLoginController)
	admin := e.Group("/admin")
	admin.Use(middleware.JWT([]byte(configs.JWT_KEY)))
	admin.POST("/signup", controllers.AdminSignUpController, middlewares.CheckRole(configs.ROLE_ADMIN))

	// ADMIN MENU
	e.GET("/menu", controllers.GetMenuController)
	e.GET("/menu/:id", controllers.GetMenuByIDController)
	admin.POST("/menu", controllers.CreateMenuController, middlewares.CheckRole(configs.ROLE_ADMIN))
	admin.PUT("/menu/:id", controllers.UpdateMenuController, middlewares.CheckRole(configs.ROLE_ADMIN))
	admin.DELETE("/menu/:id", controllers.DeleteMenuController, middlewares.CheckRole(configs.ROLE_ADMIN))

	// ADMIN CATEGORY MENU
	e.GET("/category", controllers.GetCategoriesController)
	e.GET("/category/:id", controllers.GetCategoryByIDController)
	admin.POST("/category", controllers.CreateCategoryController, middlewares.CheckRole(configs.ROLE_ADMIN))
	admin.PUT("/category/:id", controllers.UpdateCategoryController, middlewares.CheckRole(configs.ROLE_ADMIN))
	admin.DELETE("/category/:id", controllers.DeleteCategoryController, middlewares.CheckRole(configs.ROLE_ADMIN))

	// ADMIN TRANSACTIONS
	admin.GET("/transaction", controllers.AdminGetAllTransactionsController, middlewares.CheckRole(configs.ROLE_ADMIN))

	// ADMIN RATINGS MANAGEMENT

	// ADMIN ORDER MANAGEMENT
	admin.GET("/order", controllers.AdminGetAllOrdersController, middlewares.CheckRole(configs.ROLE_ADMIN))
	admin.PUT("/order/:id", controllers.AdminUpdateOrderStatusController, middlewares.CheckRole(configs.ROLE_ADMIN))

	// ADMIN USER MANAGEMENT
	admin.GET("/all-user", controllers.AdminGetAllUserController, middlewares.CheckRole(configs.ROLE_ADMIN))

	return e
}
