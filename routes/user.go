package routes

import (
	"projectpenyewaanlapangan/auth"
	"projectpenyewaanlapangan/handler"
	"projectpenyewaanlapangan/user"

	"github.com/gin-gonic/gin"
)

var (
	// DB             *gorm.DB = config.Connect()
	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository, userDetailRepository, userProfileRepository)
	authService    = auth.NewService()
	userHandler    = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowUserHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.GET("/users/:user_id", handler.Middleware(userService, authService), userHandler.GetUserByIDHandler)
	r.DELETE("/users/:user_id", handler.Middleware(userService, authService), userHandler.DeleteUserByIDHandler)
	r.PUT("/users/:user_id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
}
