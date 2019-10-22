package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

//ValidateUserRoleAPI calls the api to check if a user role claim is valid
func ValidateUserRoleAPI(teamLeadID uint64, roleCode uint64) (bool, string) {
	requestBody, err := json.Marshal(map[string]uint64{
		"user_id":   teamLeadID,
		"role_code": roleCode,
	})
	if err != nil {
		return false, "Unable to get API data"
	}
	resp, err := http.Post(beego.AppConfig.String("validateroleendpoint"), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, "Unable to call POST API"
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "Unable to read response body"
	}

	return true, string(body)
}

//ValidationResponse returns responses for validation purposes
func ValidationResponse(code int, body bool) interface{} {
	type validationResponseData struct {
		Code int  `json:"code"`
		Body bool `json:"body"`
	}

	var response validationResponseData
	response.Code = code
	response.Body = body

	return response
}
