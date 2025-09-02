package router

import (
	"auth-api/internal/users"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRoutes(r *gin.Engine, uc users.Usecase) {
	r.POST("/api/login", users.Login(uc))

	r.GET("/api/logout", users.Logout(uc))

	r.POST("/api/register", users.Register(uc))
}
