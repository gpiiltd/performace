package controllers

import (
	"encoding/json"
	"performance/models"
	"strconv"

	"github.com/astaxie/beego"
)

//ObjectiveController controls all strategic objectives
type ObjectiveController struct {
	beego.Controller
}

//CreateObjective creates a strategic objective for a user
// @Title CreateObjective
// @Description creates a stategic Objective
// @Param	object		KPIObject 	models.KPI	true		"the KPI object"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router / [post]
func (o *ObjectiveController) CreateObjective() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(o.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		o.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		o.ServeJSON()
		return
	}
	var obj models.StrategicObjective
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &obj)
	if err != nil {
		o.Data["json"] = models.ErrorResponse(405, err.Error())
		o.ServeJSON()
		return
	}
	o.Data["json"] = models.CreateStrategicObjective(obj, user)
	o.ServeJSON()
}

//GetTeamStrategiveObjectives gets the teams strategic objective
// @Title GetStrategiveObjectives
// @Description gets a team strategic objectives
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /lead/ [get]
func (o *ObjectiveController) GetTeamStrategiveObjectives() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(o.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		o.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		o.ServeJSON()
		return
	}
	strategicObjective, err := models.GetTeamLeadStrategicObj(user)
	if err != nil {
		o.Data["json"] = models.ErrorResponse(403, err.Error())
		o.ServeJSON()
		return
	}
	o.Data["json"] = models.ValidResponse(200, strategicObjective, "success")
	o.ServeJSON()
}

//GetMemberStrategiveObjectives gets the strategic objectives of the team they belong to
// @Title GetMemberStrategiveObjectives
// @Description gets the strategic objectives of the team they belong to.
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /member/ [get]
func (o *ObjectiveController) GetMemberStrategiveObjectives() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(o.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		o.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		o.ServeJSON()
		return
	}
	strategicObjective, err := models.GetTeamMemberStrategicObj(user)
	if err != nil {
		o.Data["json"] = models.ErrorResponse(403, err.Error())
		o.ServeJSON()
		return
	}
	o.Data["json"] = models.ValidResponse(200, strategicObjective, "success")
	o.ServeJSON()
}

//DeleteStrategicObj deletes a team's strategic objective
// @Title Delete
// @Description deletes a team strategic objective
// @Param	teamid		path 	string	true		"the id of the objective you want to delete"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /:objid [delete]
func (o *ObjectiveController) DeleteStrategicObj() {
	objectiveIDstring := o.GetString(":objid")
	var teamLead models.User
	resCode, teamLead := models.GetUserFromTokenString(o.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		o.Data["json"] = models.ErrorResponse(403, "Unable to get user from token string")
		o.ServeJSON()
		return
	}
	objectiveID, err := strconv.Atoi(objectiveIDstring)
	if err != nil {
		o.Data["json"] = models.ErrorResponse(403, "Invalid Strategic Objective ID.")
		o.ServeJSON()
		return
	}
	o.Data["json"] = models.DeleteStrategicObjective(teamLead, objectiveID)
	o.ServeJSON()
}

//MarkObjectiveComplete marks a strategic objective as completed
// @Title MarkObjectiveComplete
// @Description mark strategic objective as complete
// @Param	body		body 	models.Visit	true		"The Objective ID"
// @Success 200 {string} "Success"
// @Failure 403 body is empty
// @router /complete/:objectiveid [GET]
func (o *ObjectiveController) MarkObjectiveComplete() {
	objectiveID := o.GetString(":objectiveid")
	objectiveIDint, err := strconv.Atoi(objectiveID)
	if err != nil {
		o.Data["json"] = models.ErrorResponse(403, err.Error())
		o.ServeJSON()
		return
	}
	var user models.User
	code, user := models.GetUserFromTokenString(o.Ctx.Input.Header("authorization"))
	if code != 200 {
		o.Data["json"] = models.ErrorResponse(403, "Invalid token string")
		o.ServeJSON()
		return
	}
	o.Data["json"] = models.MarkObjComplete(user, objectiveIDint)
	o.ServeJSON()
}
