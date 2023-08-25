package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"login-register-email/config"
	"login-register-email/helper"
	"login-register-email/model"
	"login-register-email/service"
	"net/http"
	"strconv"
)

var otpMap = make(map[string]int)

func LoginStart(c *gin.Context) {
	type LoginRequest struct {
		Email string `json:"email"`
	}

	var loginreq LoginRequest
	if err := c.ShouldBindJSON(&loginreq); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Message": "Invalid JSON Format"})
		return
	}

	//validate email apakah ada di database atau tidak
	var existEmail model.Student
	if err := config.DB.Where("email = ?", loginreq.Email).First(&existEmail).Error; err != nil {
		c.JSON(401, gin.H{"Error": "Email Not Found"})
		return
	}

	//send email with otp
	otp, _ := helper.GenerateOTP()
	service.SendEmailWithOTP(loginreq.Email, strconv.Itoa(otp))
	otpMap[loginreq.Email] = otp // Simpan OTP sebagai int

	c.JSON(http.StatusOK, gin.H{"Message": "Email sent with OTP"})
	slog.Infof("Sending otp to %v", loginreq.Email)
}

func HandleLogin(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
		OTP   int    `json:"otp"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid JSON Format"})
		return
	}

	storedOTP, exists := otpMap[request.Email]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "OTP not found or expired"})
		return
	}

	//jika otp request nya sama dengan stored otp (otp yang dikirimkan ke email
	//send response ok dengan jwt dan delete otp nya
	if request.OTP == storedOTP {
		token, err := helper.GenerateJWT(request.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create token"})
			return
		}

		delete(otpMap, request.Email)

		// save token to cookie > exp dalam waktu 6 jam hehe
		c.SetCookie("token", token, 21600, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"Token": token})

		slog.Infof("User %v logged in !", request.Email)
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid OTP"})
	}
}
