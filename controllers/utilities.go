package controllers

import (
	"performance/models"
	"strings"
	"time"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

//TokenController handles all about token
type TokenController struct {
	beego.Controller
}

//GetTokenString gets a new token string
// @Title Self Sign-In
// @Description confirms a number have been registerd for today
// @Param	body		body 	string	true		"guest phone number"
// @Success 200 {object} models.Success
// @Failure 403 body is empty
// @router /token/:email [get]
func (t *TokenController) GetTokenString() {
	email := t.GetString(":email")
	tokenString := models.GetTokenString(email)
	t.Data["json"] = models.ErrorResponse(200, tokenString)
	t.ServeJSON()
}

//ValidateToken validates token
var ValidateToken = func(ctx *context.Context) {
	filter := Filter(ctx)
	if filter == true {
		return
	}
	type unAuthorized struct {
		Code int    `json:"code"`
		Body string `json:"body"`
	}

	token := ctx.Input.Header("authorization")
	validToken := ValidToken(token)
	if validToken != true {
		var res unAuthorized
		res.Code = 403
		res.Body = "Unauthorized Connection. Invalid Token prefix"
		ctx.Output.JSON(res, false, false)

		return
	}
	isNull := NullToken(token)
	if isNull == true {
		var res unAuthorized
		res.Code = 403
		res.Body = "Unauthorized Connection. Empty Token String"
		ctx.Output.JSON(res, false, false)

		return
	}
	if token == "" {
		var res unAuthorized
		res.Code = 403
		res.Body = "Unauthorized Connection. Empty Token"
		ctx.Output.JSON(res, false, false)

		return
	}
	isTokenExpired := TokenExpire(token)
	if isTokenExpired != true {
		var res unAuthorized
		res.Code = 401
		res.Body = "Token Expired, Kindly Login again."
		ctx.Output.JSON(res, false, false)
	}
	if strings.HasPrefix(ctx.Input.URL(), "/v1/user/validate") {
		return
	}
}

//NullToken checks if token is null
func NullToken(wholeToken string) bool {
	splitString := strings.Split(wholeToken, ",")
	if splitString[1] == "" {
		return true
	}

	return false
}

//TokenExpire checks if the user token is valid and hasn't expired
func TokenExpire(tokenS string) bool {
	wholeString := strings.Split(tokenS, ",")
	tokenString := wholeString[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})
	if err != nil {
		return false
	}
	var expireAt float64
	nowTime := time.Now().Add(time.Minute * 1).Unix()
	for key, val := range claims {
		if key == "expire" {
			expireAt = val.(float64)
		}
	}
	tm := float64(nowTime)
	diff := tm - expireAt
	if diff >= 360000 {
		return true
	}

	return true
}

//ValidToken checks if a token is valid
func ValidToken(wholeToken string) bool {
	splitString := strings.Split(wholeToken, ",")
	if splitString[0] != beego.AppConfig.String("tokenprefix") {
		return false
	}

	return true
}

//Filter checks if there are endpoint that shouldn't contain token string
func Filter(ctx *context.Context) bool {
	if strings.HasPrefix(ctx.Input.URL(), "/v1/contact/") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/token/token/") {
		return true
	}
	return false
}
