package controllers

import (
	"encoding/json"
	"performance/models"
	"strconv"

	"github.com/astaxie/beego"
)

//TeamController handles all team RELATED MATTERS
type TeamController struct {
	beego.Controller
}

//GetMyTeamInformation gets team details of authenticated user team
// @Title GetMyTeamInfo
// @Description gets a team information
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /myteam [get]
func (t *TeamController) GetMyTeamInformation() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.GetMyTeamInformations(user)
	t.ServeJSON()
}

//GetMyPendingTeam gets pending team information
// @Title GetMyPendingTeam
// @Description gets a team pending team requests
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /invitations/pending [get]
func (t *TeamController) GetMyPendingTeam() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.GetPendingTeamRequests(user)
	t.ServeJSON()
}

//AcceptTeamInvitation accepts a new team invitation
// @Title AcceptTeamInvitation
// @Description accepts a new team invitation
// @Param	teamid		path 	string	true		"the team object"
// @Success 200 {string} id of the user
// @Failure 403 body is empty
// @router /accept/:teamid [post]
func (t *TeamController) AcceptTeamInvitation() {
	teamID := t.GetString(":teamid")
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.AcceptInvitation(user, teamID)
	t.ServeJSON()
}

//VerifyHasTeam checks if a user has a team
// @Title VerifyHasTeam
// @Description checks if a user has a team
// @Param	teamid		path 	string	true		"the user id"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /verifi/:userid [get]
func (t *TeamController) VerifyHasTeam() {
	teamID := t.GetString(":userid")
	teamIDint, err := strconv.Atoi(teamID)
	if err != nil {
		t.Data["json"] = models.ErrorResponse(403, err.Error())
		t.ServeJSON()
		return
	}
	var teamLead models.User
	teamLead, err = models.GetUserDataFromID(teamIDint)
	if err != nil {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get user Data")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.TeamLeadHasTeam(teamLead)
	t.ServeJSON()
}

//TakeBehaviourTest takes a team behavioural tests
// @Title TakeBehaviourTest
// @Description takes a team behavioural tests
// @Param	models.BehaviourTest		path 	object	true		"the test result"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /behaviour/ [post]
func (t *TeamController) TakeBehaviourTest() {
	var behaviour models.BehaviourTest
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &behaviour)
	if err != nil {
		t.Data["json"] = models.ErrorResponse(405, err.Error())
		t.ServeJSON()
		return
	}
	var teamLead models.User
	resCode, teamLead := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.BehaviourTestResults(teamLead, behaviour)
	t.ServeJSON()
}

//GetTeamReport gets all team report for the head of HR
// @Title GetTeamReport
// @Description gets a team list report
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /report [get]
func (t *TeamController) GetTeamReport() {
	allTeam := models.GetTeamReport()
	t.Data["json"] = models.ValidResponse(200, allTeam, "success")
	t.ServeJSON()
}

//DeleteTeam deletes a team
// @Title Delete
// @Description deletes a team using the team id
// @Param	teamid		path 	string	true		"the id of the team you want to delete"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router / [delete]
func (t *TeamController) DeleteTeam() {
	var teamLead models.User
	resCode, teamLead := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get user from token string")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.DeleteTeamFunc(teamLead)
	t.ServeJSON()
}

//DeleteTeamMember deletes a team member
// @Title DeleteTeamMember
// @Description deletes a team member using the team id
// @Param	teamid		path 	string	true		"the id of the team member you want to delete"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /member/:id [delete]
func (t *TeamController) DeleteTeamMember() {
	var teamLead models.User
	resCode, teamLead := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get user from token string")
		t.ServeJSON()
		return
	}
	memberIDstring := t.GetString(":id")
	memberIDint, err := strconv.Atoi(memberIDstring)
	if err != nil {
		t.Data["json"] = models.ValidResponse(403, "Invalid Member ID", err.Error())
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.DeleteTeamMemberFunc(teamLead, memberIDint)
	t.ServeJSON()
}

//DeleteTeamPendingInvitation deletes a team's pending invitation
// @Title Delete
// @Description deletes a team pending invtation.
// @Param	teamid		path 	string	true		"the id of the invitation you want to delete"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /invitations/:invitationid [delete]
func (t *TeamController) DeleteTeamPendingInvitation() {
	invitationID := t.GetString(":invitationid")
	var teamLead models.User
	resCode, teamLead := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get user from token string")
		t.ServeJSON()
		return
	}
	invitationINT, err := strconv.Atoi(invitationID)
	if err != nil {
		t.Data["json"] = models.ErrorResponse(403, "Invalid invitationID.")
		t.ServeJSON()
		return
	}
	t.Data["json"] = models.DeletePendingTeamInvitation(teamLead, invitationINT)
	t.ServeJSON()
}

//GetNonMembers gets pending team information
// @Title GetMyPendingTeam
// @Description gets a team pending team requests
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /non/ [get]
func (t *TeamController) GetNonMembers() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(t.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		t.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		t.ServeJSON()
		return
	}

	var nonMembers []models.User
	nonMembers = models.GetNonMembers(user)

	t.Data["json"] = models.ValidResponse(200, nonMembers, "All Non-members of your team")
	t.ServeJSON()
}
