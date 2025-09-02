package router

import (
	"auth-api/internal/users"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRoutes(r *gin.Engine, uc users.Usecase) {
	r.POST("/login", users.Login(uc))

	r.GET("/logout", users.Logout(uc))

	r.POST("/register", users.Register(uc))
}
