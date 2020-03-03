package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

//ValidResponseData holds exported data for valid response function.
type ValidResponseData struct {
	Code    int         `json:"code"`
	Body    interface{} `json:"body"`
	Message string      `json:"message"`
}

//ValidResponse structures the data for all API response.
func ValidResponse(code int, body interface{}, message string) ValidResponseData {
	var response ValidResponseData
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
		return false, "Unable to call POST API because: " + err.Error()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "Unable to read response body"
	}

	return true, string(body)
}

//ValidationResponseData holds data that needs a true or false response
type ValidationResponseData struct {
	Code int  `json:"code"`
	Body bool `json:"body"`
}

//ValidationResponse returns responses for validation purposes
func ValidationResponse(code int, body bool) ValidationResponseData {
	var response ValidationResponseData
	response.Code = code
	response.Body = body

	return response
}

//GetUserDataFromID gets a user information from the id provided
func GetUserDataFromID(userID int) (User, error) {
	var userData User
	if findUser := Conn.Where("id = ?", userID).Find(&userData); findUser.Error != nil {
		return userData, findUser.Error
	}
	return userData, nil
}

//ConvertStringToUint64 converts an integer to usigned integer
func ConvertStringToUint64(data string) (uint64, error) {
	dataUInt, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		LogError(err)
		return dataUInt, err
	}

	return dataUInt, nil
}

//LogError writes all application error into the error log file
func LogError(funcError error) {
	f, _ := os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	logger := log.New(f, "PAS: ", log.LstdFlags)
	logger.Println(funcError.Error())
}
