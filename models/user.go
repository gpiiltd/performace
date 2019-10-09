package models

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

//GetUserDataEmail gets user data from email
func GetUserDataEmail(email string) (User, error) {
	var u User
	data := Conn.Where("email = ?", email).Find(&u)
	if data != nil && data.Error != nil {
		return u, data.Error
	}

	return u, nil
}

//SplitToken splits token token and
func SplitToken(wholeToken string) (int, string) {
	splitString := strings.Split(wholeToken, ",")
	if splitString[0] != beego.AppConfig.String("tokenprefix") {
		return 403, splitString[1]
	}

	return 200, splitString[1]
}

//GetUserFromTokenString get full user information from string
func GetUserFromTokenString(token string) (int, User) {
	var user User
	code, tokenString := SplitToken(token)
	if code != 200 {
		return 401, user
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})
	if err != nil {
		return 401, user
	}
	var email string
	for key, val := range claims {
		if key == "email" {
			email = val.(string)
		}
	}
	if email != "" {
		user, err = GetUserDataEmail(email)
		if err != nil {
			return 404, user
		}
	}
	return 200, user
}

//CreateUser creates a new user
func CreateUser(user User) {
	Conn.Create(&user)
	return
}

//UpdateProfile updates a user profile
func UpdateProfile(update User, authenticatedUser User) interface{} {
	// log.Println(update)
	verifiObject := VerifiUpdateProfile(update)
	if verifiObject != true {
		return ErrorResponse(403, "User object is not valid")
	}
	if update.Email != authenticatedUser.Email {
		return ErrorResponse(403, "Unauthorized access")
	}
	var u User
	Conn.Model(&u).Where("id = ?", authenticatedUser.ID).Updates(User{JobTitle: update.JobTitle, Number: update.Number, Location: update.Location})
	return ValidResponse(200, u, "User profile changed succesfully")
}

//VerifiUpdateProfile verifis that the fullname and the email is not left blank
func VerifiUpdateProfile(u User) bool {
	if u.FullName == "" {
		return false
	}
	if u.Email == "" {
		return false
	}
	return true
}

//GetDataFromIDString retrieves user data from ID string
func GetDataFromIDString(id string) (User, error) {
	var u User
	if getUserData := Conn.Where("id = ?", id).Find(&u); getUserData.Error != nil {
		return u, getUserData.Error
	}
	return u, nil
}
