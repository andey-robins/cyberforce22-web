package main

import (
	"github.com/gin-gonic/gin"
	"seekandyouwillbefound.org/api"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(api.SetUserStatus())
	r.NoRoute(api.FOF)

	//Static Routes
	r.GET("/", api.EnsureLoggedIn(), api.Home)
	r.POST("/", api.EnsureLoggedIn(), api.Home)
	r.GET("/login", api.EnsureNotLoggedIn(), api.Login)
	r.POST("/login", api.EnsureNotLoggedIn(), api.PerformLogin)
	r.GET("/logout", api.EnsureLoggedIn(), api.PerformLogout)
	r.GET("/manufacturing", api.EnsureLoggedIn(), api.Manufacturing)
	r.GET("/solar-generation", api.EnsureLoggedIn(), api.SolarGeneration)
	r.GET("/contact-us", api.EnsureLoggedIn(), api.ContactUs)
	r.POST("/contact", api.EnsureLoggedIn(), api.PerformContact)
	r.GET("/admin", api.EnsureLoggedIn(), api.EnsureAdminLoggedIn(), api.Admin)

	//Load HTML Templates
	r.Static("/assets", "../website/assets")
	r.LoadHTMLGlob("../website/*.html")

	r.Run(":80")
}
