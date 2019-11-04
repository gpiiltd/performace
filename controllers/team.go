package controllers

import (
	"performance/models"

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
