package service

import (
	"log"
	"net/http"
	db "task/db"
	dto "task/dto"

	"github.com/gin-gonic/gin"
)

func AddStudent(c *gin.Context) {
	var err error
	var stu dto.Student
	if err := c.BindJSON(&stu); err != nil {
		return
	}
	err = db.Mongo.InsertOne(stu)
	if err != nil {
		log.Println("ERROR ON INSERT DOCUMENT")
		return
	}
	c.IndentedJSON(http.StatusCreated, "student Added successfulyy")
}

func GetStudent(c *gin.Context) {
	var err error
	var stu []dto.Student
	err = db.Mongo.FindAll(&stu, nil, nil)
	if err != nil {
		log.Println("ERROR ON FIND DOCUMENT", err)
		return
	}
	c.IndentedJSON(http.StatusCreated, stu)
}

func UpdateMob(c *gin.Context) {
	var stu dto.Student
	if err := c.BindJSON(&stu); err != nil {
		return
	}
	set := db.M{"$set": db.M{"mobile": stu.Mobile}}
	opts := new(db.Options).SetReturnDocument(db.After)
	opts.SetUpsert(true)
	err := db.Mongo.FindAndUpdate(&stu, db.M{"email": stu.Email}, set, opts)
	if err != nil {
		log.Println("ERROR ON UPDATE DOCUMENT", err)
		return
	}
	c.IndentedJSON(http.StatusCreated, stu)
}

func DeleteStudent(c *gin.Context) {
	var stu dto.Student
	if err := c.BindJSON(&stu); err != nil {
		return
	}
	err := db.Mongo.DeleteOne(&stu, db.M{"email": stu.Email})
	if err != nil {
		log.Println("ERROR ON DELETE DOCUMENT", err)
		return
	}
	c.IndentedJSON(http.StatusCreated, "DELETED THE DOCUMENT SUCCESSFULLY")
}

func GetStudentByEmail(c *gin.Context) {
	var stu dto.Student
	if err := c.BindJSON(&stu); err != nil {
		return
	}
	err := db.Mongo.Find(&stu, db.M{"email": stu.Email})
	if err != nil {
		log.Println("ERROR ON DELETE DOCUMENT", err)
		return
	}
	c.IndentedJSON(http.StatusCreated, stu)
}
