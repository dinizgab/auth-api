package users

import (
	"errors"
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
			return
		}

		err = uc.CreateUser(c, user)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error creating user"})
			return
		}

		c.Status(http.StatusCreated)
	}
}

func Login(uc Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var login struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := c.ShouldBindJSON(&login)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
			return
		}

		tokenPair, err := uc.Login(c, login.Email, login.Password)
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, errors.New("wrong password")) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error logging in"})
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		// Max age is set with the same age as the cookie
		// 15 min
		c.SetCookie("access_token", tokenPair.AccessToken, 900, "/", "", false, true)
		// 30 days
		c.SetCookie("refresh_token", tokenPair.RefreshToken, 2592000, "/", "", false, true)

		c.Status(http.StatusOK)
	}
}

func Logout(uc Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
