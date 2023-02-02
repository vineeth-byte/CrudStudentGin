package service

import (
	"log"
	"net/http"
	db "task/db"
	dto "task/dto"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(c *gin.Context) {
	var err error
	var user dto.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	password := user.Password
	user.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return
	}
	password = ""
	err = db.Mongo.InsertOne(user)
	if err != nil {
		log.Println("ERROR ON INSERT DOCUMENT")
		return
	}
	c.IndentedJSON(http.StatusCreated, "USER ADDED SUCCESSFULLY")
}
