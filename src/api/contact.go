package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"seekandyouwillbefound.org/db"
)

func PerformContact(c *gin.Context) {
	// obviously this isn't the most secure, but with our time requirements and threat
	// models, this is an acceptable risk for right now
	// if we could use HTTPS, it would be way better, but without a CA, no dice
	name := c.PostForm("inputName")
	email := c.PostForm("inputEmail")
	phone := c.PostForm("inputNumber")
	form := c.PostForm("formFile")
	fmt.Println(form)

	db.DBPut(db.NewEntry(name, email, phone, form, []byte(form)))
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
