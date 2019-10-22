package controllers

import (
	"encoding/json"
	"performance/models"

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
		kpi.Data["json"] = models.ErrorResponse(405, "Method not allowed")
		kpi.ServeJSON()
		return
	}
	kpi.Data["json"] = models.AssignNewKPI(kpiInfo, user)
	kpi.ServeJSON()
}
