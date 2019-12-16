package controllers

import (
	"encoding/json"
	"performance/models"
	"strconv"

	"github.com/astaxie/beego"
)

//KPIController controls all KPI related activities
type KPIController struct {
	beego.Controller
}

//AssignKPI assigns KPI to a member of the team
// @Title AssignKPI
// @Description assigns new KPI information to a member of the team
// @Param	object		KPIObject 	models.KPI	true		"the KPI object"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /assign/ [post]
func (kpi *KPIController) AssignKPI() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(kpi.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		kpi.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		kpi.ServeJSON()
		return
	}
	var kpiInfo models.KPI
	err := json.Unmarshal(kpi.Ctx.Input.RequestBody, &kpiInfo)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(405, err.Error())
		kpi.ServeJSON()
		return
	}
	kpi.Data["json"] = models.AssignNewKPI(kpiInfo, user)
	kpi.ServeJSON()
}

//GetMemberKPIReport gets the kpi information for a user month
// @Title GetMemberKPIReport
// @Description gets a team member kpi information for the month
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /report/:userid/:month [get]
func (kpi *KPIController) GetMemberKPIReport() {
	userID := kpi.GetString(":userid")
	var teamMember models.User
	teamMember, err := models.GetDataFromIDString(userID)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(403, err.Error())
		kpi.ServeJSON()
		return
	}
	monthString := kpi.GetString(":month")
	month, err := strconv.Atoi(monthString)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(403, "Invalid month string.")
		kpi.ServeJSON()
		return
	}
	var teamLead models.User
	resCode, teamLead := models.GetUserFromTokenString(kpi.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		kpi.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
		kpi.ServeJSON()
		return
	}
	kpi.Data["json"] = models.TeamMemberKPIReport(teamLead, month, teamMember)
	kpi.ServeJSON()
}

//DeleteKPI removes a user kpi from the system
// @Title DeleteKPI
// @Description delete a user kpi information
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:kpiid [delete]
func (kpi *KPIController) DeleteKPI() {
	kpiID := kpi.GetString(":kpiid")
	var teamLead models.User
	code, teamLead := models.GetUserFromTokenString(kpi.Ctx.Input.Header("authorization"))
	if code != 200 {
		kpi.Data["json"] = models.ErrorResponse(403, "Invalid token string")
		kpi.ServeJSON()
		return
	}
	var kpi models.KPI
	kpi, err := models.GetKPIFromIDString(kpiID)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(403, "Error Retrieving kpi from kpiID")
		kpi.ServeJSON()
		return
	}
	deleteKPI := 
	kpi.Data["json"] = ideas.DeleteIdea(uid, mentee)
	kpi.ServeJSON()
}

//GetMonthKPI get the kpi of the specified user id and month
// @Title GetMonthKPI
// @Description gets a kpi for the month
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /kpi/ [get]
// func (kpi *KPIController) GetMonthKPI() {
// 	var user models.User
// 	resCode, user := models.GetUserFromTokenString(kpi.Ctx.Input.Header("authorization"))
// 	if resCode != 200 {
// 		kpi.Data["json"] = models.ErrorResponse(403, "Unable to get token string")
// 		kpi.ServeJSON()
// 		return
// 	}
// 	var kpiRequest models.KPIRequest
// 	err := json.Unmarshal(kpi.Ctx.Input.RequestBody, &kpiRequest)
// 	if err != nil {
// 		kpi.Data["json"] = models.ErrorResponse(405, err.Error())
// 		kpi.ServeJSON()
// 		return
// 	}
// 	kpi.Data["json"] = models.GetKPIForTheMonth(kpiRequest, user)
// 	kpi.ServeJSON()
// }
