package controllers

import (
	"performance/models"

	"github.com/astaxie/beego"
)

//ReportController gets the report of the system KPI
type ReportController struct {
	beego.Controller
}

//GetTeamReport gets a team report for the month
// @Title GetTeamReport
// @Description gets a team reports for the month
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @Param	object		TaskObject 	models.TaskTracker	true		"the monthly report"
// @router /kpi/:month [GET]
func (rep *ReportController) GetTeamReport() {
	month := rep.GetString(":month")
	rep.Data["json"] = models.GetReport(rep.Ctx.Input.Header("authorization"), month)
	rep.ServeJSON()
	return
}
