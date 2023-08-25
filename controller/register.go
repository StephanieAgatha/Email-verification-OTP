package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"golang.org/x/crypto/bcrypt"
	"log"
	"login-register-email/config"
	"login-register-email/model"
	"login-register-email/service"
)

func Register(c *gin.Context) {
	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Erorr": "Invalid JSON Format"})
		return
	}

	//email check
	var emailexist model.Student
	if err := config.DB.Where("email = ?", student.Email).First(&emailexist).Error; err == nil {
		c.AbortWithStatusJSON(401, gin.H{"Error": "Email already exists"})
		return
	}

	// username check
	var usernameexist model.Student
	if err := config.DB.Where("username = ?", student.Username).First(&usernameexist).Error; err == nil {
		c.AbortWithStatusJSON(401, gin.H{"Error": "Username already exists"})
		return
	}

	//generate bcrypt
	hashpw, err := bcrypt.GenerateFromPassword([]byte(student.Password), 10)
	if err != nil {
		log.Println("Failed to generate bcrypt ", err.Error())
		return
	}
	student.Password = string(hashpw)

	//save to db
	if err = config.DB.Create(&student).Error; err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		log.Println("Failed to register ", err.Error())
		return
	}

	slog.Infof("New student with email : %v", student.Email)
	c.JSON(200, gin.H{"Message": "Success"})

	//implement sendmail via api
	service.SendEmailRegister(student.Email, student.Firstname)
	slog.Infof("Success sending email to : %v", student.Email)
	return
}
