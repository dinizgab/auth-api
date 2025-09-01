package users

import "github.com/gin-gonic/gin"

func Register(uc UsersUsecase) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func Login(uc UsersUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Logout(uc UsersUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
