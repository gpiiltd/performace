package controllers

import (
	"encoding/json"
	"performance/models"

	"github.com/astaxie/beego"
)

//UserController Operations about Users
type UserController struct {
	beego.Controller
}

//UpdateProfile updates a user profile
// @Title UpdateProfile
// @Description update the user profile
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /update/ [post]
func (u *UserController) UpdateProfile() {
	token := u.Ctx.Input.Header("authorization")
	var update models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &update)
	if err != nil {
		u.Data["json"] = models.ErrorResponse(405, err.Error())
		u.ServeJSON()
		return
	}
	var user models.User
	resCode, user := models.GetUserFromTokenString(token)
	if resCode != 200 {
		if resCode == 404 {
			models.CreateUser(update)
			u.Data["json"] = models.ErrorResponse(200, "success")
			u.ServeJSON()
		}
		u.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		u.ServeJSON()
		return
	}
	u.Data["json"] = models.UpdateProfile(update, user)
	u.ServeJSON()
}
