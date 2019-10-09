package controllers

import (
	"performance/models"

	"github.com/astaxie/beego"
)

//TeamLeadController handles all team lead functionalities
type TeamLeadController struct {
	beego.Controller
}

//AddNewMember add a new team member to the system
// @Title AddNewMember
// @Description adds a new team member using the user ID
// @Param	visitid		path 	string	true		"the id of the user you want to make a front desk officer"
// @Success 200 {string} id of the user
// @Failure 403 body is empty
// @router /member/:id [post]
func (t *TeamLeadController) AddNewMember() {
	var member models.User
	memberID := t.GetString(":id")
	member, err := models.GetDataFromIDString(memberID)
	if err != nil {
		t.Data["json"] = models.ErrorResponse(404, "Member data does not exist on the system")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.AddNewTeamMember(member)
	t.ServeJSON()
}
