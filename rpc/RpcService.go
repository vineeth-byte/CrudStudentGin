package rpc

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"task/db"
	"task/dto"
	service "task/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Connect() {
	api := gin.Default()
	api.Use(Middleware())
	api.POST("/adduser", service.AddUser)
	api.POST("/addStudent", service.AddStudent)
	api.GET("/getStudent", service.GetStudent)
	api.PUT("/updateMob", service.UpdateMob)
	api.DELETE("/deleteStu", service.DeleteStudent)
	api.DELETE("/getStudentByEmail", service.GetStudentByEmail)
	api.Run(":5000")
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.FullPath()
		urlSplit := strings.Split(url, "/")
		fmt.Println(urlSplit)
		if urlSplit[1] != "adduser" {
			var user dto.User
			value := c.Request.Header["Authorization"]
			if len(value) >= 1 {
				valSplit := strings.Split(value[0], " ")
				auth, _ := base64.StdEncoding.DecodeString(valSplit[1])
				cred := strings.Split(string(auth), ":")
				emailId := cred[0]
				password := cred[1]
				err := db.Mongo.Find(&user, bson.M{"emailId": emailId})
				if err != nil {
					log.Println("Error on FIND DOCUMENT")
					return
				}
				passwordEncryption := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
				if passwordEncryption != nil {
					service.RespondWithError(c, 404, "Error on DECRYPTION")
					log.Println("Error on DECRYPTION")
					return
				}
			}
		}
		c.Next()
	}
}
