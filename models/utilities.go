package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

//GetTokenString generates and returns a string.
func GetTokenString(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"expire": time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("jwtkey")))
	if err != nil {
		panic(err)
	}

	return tokenString
}

//ValidResponse structures the data for all API response.
func ValidResponse(code int, body interface{}, message string) interface{} {
	type validResponseData struct {
		Code    int         `json:"code"`
		Body    interface{} `json:"body"`
		Message string      `json:"message"`
	}
	var response validResponseData
	response.Code = code
	response.Body = body
	response.Message = message

	return response
}

//ErrorResponse structures error messages
func ErrorResponse(code int, message string) interface{} {
	var response ErrorResponseData
	response.Code = code
	response.Status = message

	return response
}
