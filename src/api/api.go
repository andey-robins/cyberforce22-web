package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"seekandyouwillbefound.org/db"
)

func Home(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home",
		gin.H{},
	)
}

func Login(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login",
		gin.H{},
	)
}

func Manufacturing(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"manufacturing",
		gin.H{},
	)
}

func SolarGeneration(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"solar-generation",
		gin.H{},
	)
}

func ContactUs(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"contact-us",
		gin.H{},
	)
}

func Admin(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"admin",
		gin.H{
			"filePayload":  db.DBGetFiles(),
			"emailPayload": db.DBGetEmails(),
		},
	)
}

func FOF(c *gin.Context) {
	c.HTML(
		http.StatusNotFound,
		"fof",
		gin.H{},
	)
}

func FourOhOne(c *gin.Context) {
	c.HTML(
		http.StatusUnauthorized,
		"unauthorized",
		gin.H{},
	)
}
