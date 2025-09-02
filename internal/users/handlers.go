package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(uc Usecase) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		}

		err = uc.CreateUser(c, user)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error creating user"})
		}

		c.Status(http.StatusCreated)
	}
}

func Login(uc Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Logout(uc Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
