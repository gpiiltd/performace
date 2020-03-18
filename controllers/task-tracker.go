package controllers

import (
	"encoding/json"
	"fmt"
	"performance/models"
	"strconv"

	"github.com/astaxie/beego"
)

//TTController controls all task tracking related activities
type TTController struct {
	beego.Controller
}

//CreateNewTask creates a new task to be tracked
// @Title CreateNewTask
// @Description creates a stategic Objective
// @Param	object		KPIObject 	models.TaskTracker	true		"the task object needed"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router / [post]
func (tt *TTController) CreateNewTask() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		tt.Data["json"] = models.ValidResponse(403, "Unable to get token string", "error")
		tt.ServeJSON()
		return
	}
	var task models.TaskTracker
	err := json.Unmarshal(tt.Ctx.Input.RequestBody, &task)
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, err.Error(), "error")
		tt.ServeJSON()
		return
	}
	tt.Data["json"] = models.CreateTaskTrack(user, task)
	tt.ServeJSON()
	return
}

//GetMyTasks gets user tasks for the specified day, month, and year
// @Title GetUserTasks
// @Description gets a team user task information for a particular day, time and year
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @Param	object		TaskObject 	models.TaskTracker	true		"the task object containing the day, month and year"
// @router /:day/:month/:year [GET]
func (tt *TTController) GetMyTasks() {
	day := tt.GetString(":day")
	month := tt.GetString(":month")
	year := tt.GetString(":year")

	var dayInt, monthInt, yearInt uint64
	dayInt, monthInt, yearInt, err := models.ConvertDayMonthYear(day, month, year)
	if err != nil {
		tt.Data["json"] = models.ErrorResponse(403, err.Error())
		tt.ServeJSON()
		return
	}

	var taskTrackerObject models.TaskTracker
	taskTrackerObject.Day = dayInt
	taskTrackerObject.Month = monthInt
	taskTrackerObject.Year = yearInt

	var user models.User
	resCode, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		tt.Data["json"] = models.ValidResponse(403, "Unable to get token string", "error")
		tt.ServeJSON()
		return
	}
	tt.Data["json"] = models.GetTask(user, taskTrackerObject)
	tt.ServeJSON()
	return
}

//GetAllUsersTasks gets user all tasks for the specified day, month, and year
// @Title GetUserTasks
// @Description gets a team user task information for a particular day, time and year
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @Param	object		TaskObject 	models.TaskTracker	true		"the task object containing the day, month and year"
// @router /alltask/ [POST]
func (tt *TTController) GetAllUsersTasks() {
	tasksInfo := models.TaskInfo{}
	err := json.Unmarshal(tt.Ctx.Input.RequestBody, &tasksInfo)
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, err.Error(), "error")
		tt.ServeJSON()
		return
	}

	var dayInt, monthInt, yearInt uint64
	dayInt, monthInt, yearInt, err = models.ConvertDayMonthYear(tasksInfo.Day, tasksInfo.Month, tasksInfo.Year)
	if err != nil {
		tt.Data["json"] = models.ErrorResponse(403, err.Error())
		tt.ServeJSON()
		return
	}

	var taskTrackerObject models.TaskTracker
	taskTrackerObject.Day = dayInt
	taskTrackerObject.Month = monthInt
	taskTrackerObject.Year = yearInt

	allUserTask, err := models.GetAllTodayTask(tasksInfo)
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, "Error Getting all user tasks for today", err.Error())
		tt.ServeJSON()
		return
	}

	tt.Data["json"] = models.ValidResponse(200, allUserTask, "success")
	tt.ServeJSON()
	return
}

//GetAllUsersTasksHistory gets user all task history
// @Title GetAllUsersTasksHistory
// @Description gets a team user task information for a particular day, time and year
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @Param	object		TaskObject 	models.TaskTracker	true		"the task object containing the day, month and year"
// @router /alltask/all [GET]
func (tt *TTController) GetAllUsersTasksHistory() {
	allUserTask, err := models.GetAllTaskHistory()
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, "Error Getting all user tasks for today", err.Error())
		tt.ServeJSON()
		return
	}

	tt.Data["json"] = models.ValidResponse(200, allUserTask, "success")
	tt.ServeJSON()
	return
}

//GetAllUncompletedTasks gets user uncompleted taskp0
// @Title GetAllUncompletedTasks
// @Description gets a team user task information for a particular day, time and year
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @Param	object		TaskObject 	models.TaskTracker	true		"the task object containing the day, month and year"
// @router /uncomplete/ [GET]
func (tt *TTController) GetAllUncompletedTasks() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		tt.Data["json"] = models.ValidResponse(403, "Unable to get token string", "error")
		tt.ServeJSON()
		return
	}

	var uncompletedTasks []models.TaskTracker
	uncompletedTasks, err := models.GetAllUncompleteTasks(user)
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, "Unable to get uncompleted tasks", err.Error())
		tt.ServeJSON()
		return
	}

	tt.Data["json"] = models.ValidResponse(200, uncompletedTasks, "success")
	tt.ServeJSON()
	return
}

//GetAllUserUncompletedTasks gets a particular user uncompleted task
// @Title GetAllUserUncompletedTask
// @Description gets a team user task information for a particular day, time and year
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @Param	object		TaskObject 	models.TaskTracker	true		"the task object containing the day, month and year"
// @router /uncomplete/:userid [GET]
func (tt *TTController) GetAllUserUncompletedTasks() {
	userID := tt.GetString(":userid")
	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		tt.Data["json"] = models.ErrorResponse(403, "Invalid User ID")
		tt.ServeJSON()
		return
	}
	var taskOwner models.User

	var user models.User
	resCode, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		tt.Data["json"] = models.ValidResponse(403, "Unable to get token string", "error")
		tt.ServeJSON()
		return
	}

	//convert userID to uint64
	userIDuint, _ := strconv.ParseUint(userID, 10, 64)

	if user.ID == userIDuint {
		taskOwner = user
	} else {
		taskOwner, err = models.GetUserDataFromID(userIDint)
		if err != nil {
			tt.Data["json"] = models.ValidResponse(403, "Unable to get user from user ID", "error")
			tt.ServeJSON()
			return
		}
	}

	var uncompletedTasks []models.TaskTracker
	uncompletedTasks, err = models.GetAllUserUncompleteTasks(taskOwner)
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, "Unable to get uncompleted tasks", err.Error())
		tt.ServeJSON()
		return
	}

	response := models.ValidResponse(200, uncompletedTasks, "success")

	tt.Data["json"] = response
	tt.ServeJSON()
	return
}

//StartTaskTracking indicates that a user has started a task
// @Title StartTaskTracking
// @Description starts a specified task
// @Success 200 {string} "Success"
// @Failure 403 body is empty
// @router /start/:tid [GET]
func (tt *TTController) StartTaskTracking() {
	taskID := tt.GetString(":tid")
	var task models.TaskTracker
	task, err := models.GetTrackedTaskFromID(taskID)
	if err != nil {
		response := models.ErrorResponse(405, "Unable to get task from task_id")
		tt.Data["json"] = response
		tt.ServeJSON()
		return
	}
	var user models.User
	code, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if code != 200 {
		tt.Data["json"] = models.ErrorResponse(403, "Invalid token string")
		tt.ServeJSON()
		return
	}
	tt.Data["json"] = models.StartTrackingTask(user, task)
	tt.ServeJSON()
	return
}

//CompleteTaskTracking indicates that a user has completed a task
// @Title StartTaskTracking
// @Description completes a specified task
// @Param	object		KPIObject 	models.TaskTracker	true		"the task update data object needed"
// @Success 200 {string} "Success"
// @Failure 403 body is empty
// @router /complete/ [POST]
func (tt *TTController) CompleteTaskTracking() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		tt.Data["json"] = models.ValidResponse(403, "Unable to get token string", "error")
		tt.ServeJSON()
		return
	}
	var taskUpdate models.TaskTracker
	err := json.Unmarshal(tt.Ctx.Input.RequestBody, &taskUpdate)
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, err.Error(), "error")
		tt.ServeJSON()
		return
	}
	tt.Data["json"] = models.CompleteTrackingTask(user, taskUpdate)
	tt.ServeJSON()
	return
}

//DeleteTrackedTask removes a tracked task from the system
// @Title DeleteTaskTracked
// @Description delete a tracked task
// @Param	uid		path 	string	true		"The id of Task to be deleted"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:tid [delete]
func (tt *TTController) DeleteTrackedTask() {
	taskID := tt.GetString(":tid")
	var user models.User
	code, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if code != 200 {
		tt.Data["json"] = models.ErrorResponse(403, "Invalid token string")
		tt.ServeJSON()
		return
	}
	var task models.TaskTracker
	task, err := models.GetTrackedTaskFromID(taskID)
	if err != nil {
		response := models.ErrorResponse(405, "Unable to get task from task_id")
		tt.Data["json"] = response
		tt.ServeJSON()
		return
	}
	tt.Data["json"] = models.DeleteTrackedTask(user, task)
	tt.ServeJSON()
}

//GetTrackedTask retrieves a task from ID
// @Title GetTrackedTask
// @Description gets a task from ID
// @Success 200 {string} "Success"
// @Failure 403 body is empty
// @router /:tid [GET]
func (tt *TTController) GetTrackedTask() {
	taskID := tt.GetString(":tid")
	var task models.TaskTracker
	task, err := models.GetTrackedTaskFromID(taskID)
	if err != nil {
		response := models.ErrorResponse(405, "Unable to get task from task_id")
		tt.Data["json"] = response
		tt.ServeJSON()
		return
	}
	var taskUpdates []models.TaskTrackerUpdates
	taskUpdates, _ = models.GetTaskUpdatesFromID(taskID)
	structuredUpdates := models.StructureTaskAndUpdates(task, taskUpdates)
	tt.Data["json"] = models.ValidResponse(200, structuredUpdates, "success")
	tt.ServeJSON()
	return
}

//GetTeamMemberTask gets user all tasks for the specified day, month, and year
// @Title GetUserTasks
// @Description gets a team user task information for a particular day, time and year
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @Param	object		TaskObject 	models.TaskTracker	true		"the task object containing the day, month and year"
// @router /:day/:month/:year/:memberid [GET]
func (tt *TTController) GetTeamMemberTask() {
	day := tt.GetString(":day")
	month := tt.GetString(":month")
	year := tt.GetString(":year")
	memberID := tt.GetString(":memberid")
	memberUint, _ := models.ConvertStringToUint64(memberID)

	var dayInt, monthInt, yearInt uint64
	dayInt, monthInt, yearInt, err := models.ConvertDayMonthYear(day, month, year)
	if err != nil {
		tt.Data["json"] = models.ErrorResponse(403, err.Error())
		tt.ServeJSON()
		return
	}

	var user models.User
	code, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if code != 200 {
		tt.Data["json"] = models.ErrorResponse(403, "Invalid token string")
		tt.ServeJSON()
		return
	}

	var taskTrackerObject models.TaskTracker
	taskTrackerObject.Day = dayInt
	taskTrackerObject.Month = monthInt
	taskTrackerObject.Year = yearInt
	taskTrackerObject.UserID = memberUint

	allUserTask, err := models.GetTeamMemberTodayTask(taskTrackerObject, user)
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, "Error Getting all user tasks for today", err.Error())
		tt.ServeJSON()
		return
	}

	tt.Data["json"] = models.ValidResponse(200, allUserTask, "success")
	tt.ServeJSON()
	return
}

//AddNewUpdate add a new update
// @Title AddNewUpdate
// @Description adds a new update to the task
// @Param	object		KPIObject 	models.TaskTrackerUpdates	true		"the task update data object needed"
// @Success 200 {object} models.ValidResponse
// @Failure 403 body is empty
// @router /update/ [post]
func (tt *TTController) AddNewUpdate() {
	var user models.User
	resCode, user := models.GetUserFromTokenString(tt.Ctx.Input.Header("authorization"))
	if resCode != 200 {
		tt.Data["json"] = models.ValidResponse(403, "Unable to get token string", "error")
		tt.ServeJSON()
		return
	}
	var taskUpdate models.TaskTrackerUpdates
	err := json.Unmarshal(tt.Ctx.Input.RequestBody, &taskUpdate)
	if err != nil {
		tt.Data["json"] = models.ValidResponse(403, err.Error(), "error")
		tt.ServeJSON()
		return
	}
	tt.Data["json"] = models.UpdateTaskProgress(user, taskUpdate)
	tt.ServeJSON()
	return
}

//GetUserTrackedTask retrieves a tracked task from a user
// @Title GetUserTrackedTask
// @Description gets a tracked task from user ID
// @Success 200 {string} "success"
// @Failure 403 body is empty
// @router /user/:userid [GET]
func (tt *TTController) GetUserTrackedTask() {
	userID := tt.GetString(":userid")
	var task []models.TaskTracker
	task, err := models.GetUserTrackedTasks(userID)
	if err != nil {
		tt.Data["json"] = models.ErrorResponse(405, "Unable to get task from user_id")
		tt.ServeJSON()
		return
	}
	var sortedTasks []models.TaskTracker
	sortedTasks = models.RemoveTodayFromTasks(task)
	var taskUpdates []models.TaskTrackerUpdates
	var structuredTaskAndUpdates []models.TaskUpdateResponseBody
	var structuredUpdates models.TaskUpdateResponseBody

	for _, individualTasks := range sortedTasks {
		taskUpdates, _ = models.GetTaskUpdatesFromID(fmt.Sprint(individualTasks.ID))
		structuredUpdates = models.StructureTaskAndUpdates(individualTasks, taskUpdates)
		structuredTaskAndUpdates = append(structuredTaskAndUpdates, structuredUpdates)
	}

	tt.Data["json"] = models.ValidResponse(200, structuredTaskAndUpdates, "success")
	tt.ServeJSON()
	return
}
