package helper

import (
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"strconv"
	"time"
)

// var Secret = []byte("lalalalaaa")
var secretstr = "lalalalaaa"

func GenerateOTP() (int, error) {
	passcode := GeneratePassCode(secretstr)
	otpInt, err := strconv.Atoi(passcode)
	if err != nil {
		return 0, err
	}

	return otpInt, nil
}

//func ValidateOTP(inputOTP int) bool {
//	otpStr := strconv.Itoa(inputOTP)
//	valid := totp.Validate(otpStr, secretstr)
//	return valid
//}

func GeneratePassCode(utf8string string) string {
	secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}
