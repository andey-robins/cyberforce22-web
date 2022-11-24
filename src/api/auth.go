package api

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func adminToken() string {
	return "992db5e2595725ef68eb5066e014da8c7f6baca0afa96f1021b56b66782133f2"
}

func PerformLogin(c *gin.Context) {
	// obviously this isn't the most secure, but with our time requirements and threat
	// models, this is an acceptable risk for right now
	// if we could use HTTPS, it would be way better, but without a CA, no dice
	username := c.PostForm("username")
	password := c.PostForm("password")

	if isUserValid(username, password) {
		token := generateSessionToken(username)
		fmt.Print(token)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		c.Redirect(http.StatusUnauthorized, "/login")
	}
}

func PerformLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.SetCookie("is_logged_in", "", -1, "", "", false, true)
	c.SetCookie("is_admin", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func isUserValid(username, password string) bool {
	userList := []struct {
		Username string
		Password string
	}{
		{"plank", "5!ys!hhsds"},
		{"bob", "sjhd76eww!"},
		{"clem", "khsd54#h"},
		{"alicia", "jhsjhsd222!"},
		{"sue", "76shshs63!"},
	}

	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func generateSessionToken(u string) string {
	if u == "plank" {
		return adminToken()
	}
	hash := sha256.Sum256([]byte(u))
	return fmt.Sprintf("%v", hash)
}
