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
	var kpiInfo models.KPI
	kpiInfo, err := models.GetKPIFromIDString(kpiID)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(403, "Error Retrieving kpi from kpiID")
		kpi.ServeJSON()
		return
	}
	kpi.Data["json"] = models.DeleteKPI(kpiInfo, teamLead)
	kpi.ServeJSON()
}

//GetKPIFromID gets the kpi information for a kpi ID
// @Title GetKPIFromID
// @Description gets a kpi information from ID
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /:kpiid [get]
func (kpi *KPIController) GetKPIFromID() {
	kpiID := kpi.GetString(":kpiid")
	kpiIDint, err := strconv.Atoi(kpiID)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(403, err.Error())
		kpi.ServeJSON()
		return
	}
	kpi.Data["json"] = models.GetKPIInfo(kpiIDint)
	kpi.ServeJSON()
}

//GetKPIRange gets the kpi information for a kpi ID
// @Title GetKPIFromID
// @Description gets a kpi information from ID
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /range/ [POST]
func (kpi *KPIController) GetKPIRange() {
	var requestInfo models.DateRange
	err := json.Unmarshal(kpi.Ctx.Input.RequestBody, &requestInfo)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(405, err.Error())
		kpi.ServeJSON()
		return
	}
	rangeKPI := models.GetKPIsFromRange(requestInfo)
	kpi.Data["json"] = models.ValidResponse(200, rangeKPI, "success")
	kpi.ServeJSON()
}

//GetAllTasks gets all tasks on the system
// @Title GetAllTasks
// @Description gets all tasks belonging to a KPI
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /task/ [get]
func (kpi *KPIController) GetAllTasks() {
	getTask := models.GetAllTask()
	kpi.Data["json"] = models.ValidResponse(200, getTask, "success")
	kpi.ServeJSON()
}

//GetAllKPITasks gets all tasks belonging to a KPI
// @Title GetAllKPITasks
// @Description gets all tasks belonging to a KPI
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /task/:kpiid [get]
func (kpi *KPIController) GetAllKPITasks() {
	kpiID := kpi.GetString(":kpiid")
	kpiInt, err := strconv.Atoi(kpiID)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(403, "Invalid kpi_id string.")
		kpi.ServeJSON()
		return
	}
	getTask, err := models.GetKPITasks(kpiInt)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(403, "Invalid kpi_id string when getting tasks.")
		kpi.ServeJSON()
		return
	}
	kpi.Data["json"] = models.ValidResponse(200, getTask, "success")
	kpi.ServeJSON()
}

//CreateTask creates a new task on the system
// @Title CreateTask
// @Description create object
// @Param	body		body 	models.Visit	true		"The task data"
// @Success 200 {string} "Success"
// @Failure 403 body is empty
// @router /task/ [post]
func (kpi *KPIController) CreateTask() {
	task := models.Task{}
	err := json.Unmarshal(kpi.Ctx.Input.RequestBody, &task)
	if err != nil {
		response := models.ErrorResponse(405, err.Error())
		kpi.Data["json"] = response
		kpi.ServeJSON()
	}
	var user models.User
	code, user := models.GetUserFromTokenString(kpi.Ctx.Input.Header("authorization"))
	if code != 200 {
		kpi.Data["json"] = models.ErrorResponse(403, "Invalid token string")
		kpi.ServeJSON()
		return
	}
	createTask, status := models.CreateKPITask(task, user)
	if status != true {
		response := models.ErrorResponse(405, "Unable to create KPI task")
		kpi.Data["json"] = response
		kpi.ServeJSON()
	}
	kpi.Data["json"] = models.ValidResponse(200, createTask, "success")
	kpi.ServeJSON()
}

//MarkTaskComplete marks a task as completed
// @Title MarkTaskComplete
// @Description mark task as complete
// @Param	body		body 	models.Visit	true		"The task data"
// @Success 200 {string} "Success"
// @Failure 403 body is empty
// @router /task/complete/:tid [GET]
func (kpi *KPIController) MarkTaskComplete() {
	taskID := kpi.GetString(":tid")
	taskIDint, err := strconv.Atoi(taskID)
	if err != nil {
		kpi.Data["json"] = models.ErrorResponse(403, err.Error())
		kpi.ServeJSON()
		return
	}
	var task models.Task
	task, err = models.GetTaskFromID(taskIDint)
	if err != nil {
		response := models.ErrorResponse(405, "Unable to get task from task_id")
		kpi.Data["json"] = response
		kpi.ServeJSON()
		return
	}
	var user models.User
	code, user := models.GetUserFromTokenString(kpi.Ctx.Input.Header("authorization"))
	if code != 200 {
		kpi.Data["json"] = models.ErrorResponse(403, "Invalid token string")
		kpi.ServeJSON()
		return
	}
	kpi.Data["json"] = models.MarkTaskAsComplete(user, task)
	kpi.ServeJSON()
}

//ScoreKPI scores a kpi
// @Title ScoreKPI
// @Description create score and supervisor comment for a KPI
// @Param	body		body 	models.KPI	true		"The kpi data"
// @Success 200 {string} "Success"
// @Failure 403 body is empty
// @router /score/ [post]
func (kpi *KPIController) ScoreKPI() {
	kpiObject := models.KPI{}
	err := json.Unmarshal(kpi.Ctx.Input.RequestBody, &kpiObject)
	if err != nil {
		response := models.ErrorResponse(405, err.Error())
		kpi.Data["json"] = response
		kpi.ServeJSON()
	}
	var user models.User
	code, user := models.GetUserFromTokenString(kpi.Ctx.Input.Header("authorization"))
	if code != 200 {
		kpi.Data["json"] = models.ErrorResponse(403, "Invalid token string")
		kpi.ServeJSON()
		return
	}
	kpi.Data["json"] = models.ScoreKPComment(kpiObject, user)
	kpi.ServeJSON()
}
