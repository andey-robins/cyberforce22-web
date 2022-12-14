package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	}
}

func EnsureAdminLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err != nil || token != adminToken() {
			FourOhOne(c)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			FourOhOne(c)
		}
	}
}

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
			if token == adminToken() {
				c.Set("is_admin", true)
			} else {
				c.Set("is_admin", false)
			}
		} else {
			c.Set("is_logged_in", false)
			c.Set("is_admin", false)
		}
	}
}
