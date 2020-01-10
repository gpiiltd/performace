package controllers

import (
	"encoding/json"
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
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	var member models.User
	memberID := t.GetString(":id")
	member, err := models.GetDataFromIDString(memberID)
	if err != nil {
		t.Data["json"] = models.ErrorResponse(404, "Member data does not exist on the system")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.AddNewTeamMember(member, user)
	t.ServeJSON()
}

//CreateTeam creates a new team on the system
// @Title CreateTeam
// @Description creates a new team using the user ID
// @Param	teamid		path 	models.Team	true		"the team object"
// @Success 200 {string} id of the user
// @Failure 403 body is empty
// @router /team/ [post]
func (t *TeamLeadController) CreateTeam() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	var team models.Team
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &team)
	if err != nil {
		t.Data["json"] = models.ErrorResponse(405, "Method not allowed")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.AddNewTeam(user, team)
	t.ServeJSON()
}

//GetMyTeamInfo gets team details of authenticated user team
// @Title GetMyTeamInfo
// @Description gets a team information
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /myteam/ [get]
func (t *TeamLeadController) GetMyTeamInfo() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.GetTeamInfo(user)
	t.ServeJSON()
}

//GetMyPendingTeam gets pending team details of authenticated user team
// @Title GetMyPendingTeam
// @Description gets a team's pending member information
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /pending [get]
func (t *TeamLeadController) GetMyPendingTeam() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.MyPendingTeamInfo(user)
	t.ServeJSON()
}

//DeletePendingInvitation deletes a pending team invitation
// @Title Delete
// @Description deletes a team invitation using the user ID
// @Success 200 {string} id of the user
// @Failure 403 body is empty
// @router /pending/:id [delete]
func (t *TeamLeadController) DeletePendingInvitation() {
	teamMemberID := t.GetString(":id")
	t.Data["json"] = models.DeletePendingTeamMember(teamMemberID)
	t.ServeJSON()
}

//VerifiHasTeam checks if a teamlead has created a team
// @Title VerifiHasTeam
// @Description checks if a user has created a team
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /validate [get]
func (t *TeamLeadController) VerifiHasTeam() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	verifiLead, teamName := models.LeadHasTeam(user)
	if verifiLead != true {
		t.Data["json"] = models.ValidResponse(404, teamName, "false")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.ValidResponse(200, teamName, "true")
	t.ServeJSON()
}
