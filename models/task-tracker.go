package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//CreateTaskTrack updates a new task to be tracked
func CreateTaskTrack(user User, task TaskTracker) interface{} {
	taskStatus := "pending"

	err := ValidateTaskObject(task)
	if err != nil {
		return ValidResponse(403, err.Error(), "error")
	}

	layout := "2006-01-02 15:04:05"
	str := "0000-00-00 00:00:01"
	t, _ := time.Parse(layout, str)

	task.StartTime = time.Time{}
	task.EndTime = time.Time{}
	task.Status = taskStatus
	task.UserID = user.ID
	task.DepartmentID = user.DepartmentID

	if createTask := Conn.Create(&task); createTask.Error != nil {
		LogError(createTask.Error)
		return ValidResponse(403, err.Error(), "error")
	}

	return ValidResponse(200, task, "success")
}

//ValidateTaskObject checks if a task struct is valid
func ValidateTaskObject(task TaskTracker) error {
	if task.Day == 0 || task.Month == 0 || task.Year == 0 {
		return errors.New("Empty Date Object")
	}

	if task.Task == "" {
		return errors.New("Empty task")
	}

	return nil
}

//ValidateTeamTaskObject checks if a task struct is valid
func ValidateTeamTaskObject(task TaskTracker) error {
	if task.Day == 0 || task.Month == 0 || task.Year == 0 {
		LogError(errors.New("Empty Date Object"))
		return errors.New("Empty Date Object")
	}

	if task.UserID == 0 {
		LogError(errors.New("Empty user ID"))
		return errors.New("Empty user ID")
	}

	return nil
}

//ConvertDayMonthYear converts the day, month, year in string to integers {}
func ConvertDayMonthYear(day string, month string, year string) (uint64, uint64, uint64, error) {
	var dayInt uint64
	var monthInt uint64
	var yearInt uint64

	dayInt, err := strconv.ParseUint(day, 10, 64)
	if err != nil {
		return dayInt, 0, 0, errors.New("Unable to convert day string")
	}

	monthInt, err = strconv.ParseUint(month, 10, 64)
	if err != nil {
		return dayInt, monthInt, 0, errors.New("Unable to convert month integer")
	}

	yearInt, err = strconv.ParseUint(year, 10, 64)
	if err != nil {
		return dayInt, monthInt, yearInt, errors.New("Unable to convert year integer")
	}

	return dayInt, monthInt, yearInt, nil

}

//GetAllUncompleteTasks gets all uncompleted task of a particular user
func GetAllUncompleteTasks(user User) ([]TaskTracker, error) {
	taskStatus := "in progress"
	var allUncompletedTask []TaskTracker
	if findTasks := Conn.Where("status = ?", taskStatus).Find(&allUncompletedTask); findTasks.Error != nil {
		return []TaskTracker{}, findTasks.Error
	}
	var sortedTask []TaskTracker
	sortedTask = RemoveTodayFromUncompletedTask(allUncompletedTask)

	return sortedTask, nil
}

//GetAllUserUncompleteTasks gets all uncompleted task of a particular user
func GetAllUserUncompleteTasks(user User) ([]TaskTracker, error) {
	taskStatus := "in progress"
	var allUncompletedTask []TaskTracker
	if findTasks := Conn.Where("status = ? AND user_id = ?", taskStatus, user.ID).Find(&allUncompletedTask); findTasks.Error != nil {
		return allUncompletedTask, findTasks.Error
	}

	var sortedTask []TaskTracker
	sortedTask = RemoveTodayFromUncompletedTask(allUncompletedTask)

	return sortedTask, nil
}

//RemoveTodayFromUncompletedTask removes today's date from uncompleted tasks
func RemoveTodayFromUncompletedTask(taskTracked []TaskTracker) []TaskTracker {
	dt := time.Now().Format("01-02-2006")
	split := strings.Split(dt, "-")
	day := split[0]
	month := split[1]
	year := split[2]

	var dayInt, monthInt, yearInt uint64
	dayInt, monthInt, yearInt, _ = ConvertDayMonthYear(day, month, year)

	var sortedTask []TaskTracker

	for _, task := range taskTracked {
		if task.Day != dayInt && task.Month != monthInt && task.Year != yearInt {
			sortedTask = append(sortedTask, task)
		}
	}

	return sortedTask

}

//GetTask retrieves the task that needs to be tracked speciifed by task Day, month, and year. User Data must be passed too.
func GetTask(user User, task TaskTracker) interface{} {
	if task.Month == 0 || task.Day == 0 || task.Year == 0 {
		return ValidResponse(403, "Empty Date Range", "error")
	}
	task.UserID = user.ID

	var retrievedTask []TaskTracker
	if findTasks := Conn.Where("day = ? AND year = ? AND month = ? AND user_id = ?", task.Day, task.Year, task.Month, task.UserID).Find(&retrievedTask); findTasks.Error != nil {
		return ValidResponse(403, "Unable to get task from specified date", findTasks.Error.Error())
	}

	return ValidResponse(200, retrievedTask, "success")
}

//GetAllTodayTask retrieves the task that needs to be tracked speciifed by task Day, month, and year. This gets data for all users
func GetAllTodayTask(task TaskInfo) (interface{}, error) {
	var allUsers []User
	Conn.Find(&allUsers)

	var retrievedTask []TaskTracker
	if findTasks := Conn.Where("day = ? AND year = ? AND month = ?", task.Day, task.Year, task.Month).Find(&retrievedTask); findTasks.Error != nil {
		return []TaskTracker{}, findTasks.Error
	}

	type taskResponse struct {
		Task  TaskTracker `json:"task"`
		Staff User        `json:"staff"`
	}

	var responseObject taskResponse
	var responseObjectArray []taskResponse

	for _, thisTask := range retrievedTask {
		for _, thisUser := range allUsers {
			if thisTask.UserID == thisUser.ID {
				responseObject.Task = thisTask
				responseObject.Staff = thisUser

				responseObjectArray = append(responseObjectArray, responseObject)
			}
		}
	}

	return responseObjectArray, nil
}

//GetAllTaskHistory retrieves the task history of every user
func GetAllTaskHistory() (interface{}, error) {
	var allUsers []User
	Conn.Find(&allUsers)

	var retrievedTask []TaskTracker
	if findTasks := Conn.Find(&retrievedTask); findTasks.Error != nil {
		return []TaskTracker{}, findTasks.Error
	}

	type taskResponse struct {
		Task  TaskTracker `json:"task"`
		Staff User        `json:"staff"`
	}

	var responseObject taskResponse
	var responseObjectArray []taskResponse

	for _, thisTask := range retrievedTask {
		for _, thisUser := range allUsers {
			if thisTask.UserID == thisUser.ID {
				responseObject.Task = thisTask
				responseObject.Staff = thisUser

				responseObjectArray = append(responseObjectArray, responseObject)
			}
		}
	}

	return responseObjectArray, nil
}

//GetTeamMemberTodayTask retrieves the task that needs to be tracked speciifed by task Day, month, and year. This gets data for all users
func GetTeamMemberTodayTask(task TaskTracker, teamLead User) ([]TaskTracker, error) {
	err := ValidateTeamTaskObject(task)
	if err != nil {
		return []TaskTracker{}, err
	}

	var team Members
	if isMyTeamLead := Conn.Where("member_id = ? AND team_lead_id = ?", task.UserID, teamLead.ID).Find(&team); isMyTeamLead.Error != nil {
		return []TaskTracker{}, errors.New("Team member doesn't belong to your team")
	}

	var retrievedTask []TaskTracker
	if findTasks := Conn.Where("day = ? AND year = ? AND month = ? AND user_id = ?", task.Day, task.Year, task.Month, task.UserID).Find(&retrievedTask); findTasks.Error != nil {
		return []TaskTracker{}, findTasks.Error
	}

	return retrievedTask, nil
}

//GetTrackedTaskFromID retrieved a daily task record with ID
func GetTrackedTaskFromID(taskID string) (TaskTracker, error) {
	var taskTracked TaskTracker
	if findTask := Conn.Where("id = ?", taskID).Find(&taskTracked); findTask.Error != nil {
		return taskTracked, findTask.Error
	}

	return taskTracked, nil
}

//GetUserTrackedTasks retrieves user task information from user id
func GetUserTrackedTasks(userid string) ([]TaskTracker, error) {
	var retrievedTask []TaskTracker
	if findTasks := Conn.Where("user_id = ?", userid).Find(&retrievedTask); findTasks.Error != nil {
		LogError(findTasks.Error)
		return retrievedTask, findTasks.Error
	}

	return retrievedTask, nil
}

//GetTaskUpdatesFromID gets task updates
func GetTaskUpdatesFromID(taskID string) ([]TaskTrackerUpdates, error) {
	var taskUpdates []TaskTrackerUpdates
	if findUpdates := Conn.Where("task_id = ?", taskID).Find(&taskUpdates); findUpdates.Error != nil {
		return taskUpdates, findUpdates.Error
	}

	return taskUpdates, nil
}

//RemoveTodayFromTasks removes today data from task array
func RemoveTodayFromTasks(taskArray []TaskTracker) []TaskTracker {
	todayDate := time.Now().Format("01-02-2006")
	splitedString := strings.Split(todayDate, "-")

	day := splitedString[1]
	month := splitedString[0]
	year := splitedString[2]

	dayUint, _ := strconv.ParseUint(day, 10, 64)
	monthUint, _ := strconv.ParseUint(month, 10, 64)
	yearUint, _ := strconv.ParseUint(year, 10, 64)

	var sortedTaskArray []TaskTracker
	for _, tasks := range taskArray {
		if tasks.Day == dayUint && tasks.Month == monthUint && tasks.Year == yearUint {
			break
		}
		sortedTaskArray = append(sortedTaskArray, tasks)
	}

	return sortedTaskArray
}

//StartTrackingTask marks a task has started.
func StartTrackingTask(user User, task TaskTracker) interface{} {
	var myTask TaskTracker
	if findThisTask := Conn.Where("user_id = ? AND id = ?", user.ID, task.ID).Find(&myTask); findThisTask.Error != nil {
		return ValidResponse(403, "User not authorized to start task", findThisTask.Error.Error())
	}

	taskStatus := "pending"
	if myTask.Status != taskStatus {
		return ValidResponse(403, "User can only start pending tasks", "error")
	}

	// layout := "15:04:05"
	// str := "00:00:00"
	// t, _ := time.Parse(layout, str)
	// nowTime := "0000-00-00"
	t := time.Now()

	//change task status
	taskStatus = "in progress"
	if startTask := Conn.Model(&myTask).Where("id = ?", task.ID).Updates(TaskTracker{Status: taskStatus, StartTime: t}); startTask.Error != nil {
		LogError(startTask.Error)
		return ValidResponse(403, startTask.Error.Error(), "error")
	}
	return ValidResponse(200, "Task Started", "success")
}

//CompleteTrackingTask marks a task has started.
func CompleteTrackingTask(user User, taskUpdate TaskTracker) interface{} {
	if taskUpdate.ID == 0 {
		return ValidResponse(403, "Empty Task ID", "error")
	}

	if taskUpdate.Comments == "" {
		return ValidResponse(403, "Please enter comments about task", "error")
	}
	taskID := fmt.Sprint(taskUpdate.ID)
	var task TaskTracker
	task, err := GetTrackedTaskFromID(taskID)
	if err != nil {
		return ValidResponse(403, "Unable to get task from task_id", err.Error())
	}

	taskStatus := "in progress"
	if task.Status != taskStatus {
		return ValidResponse(403, "Plese start task", "error")
	}

	if task.UserID != user.ID {
		return ValidResponse(403, "User not authorized to update task progress", "error")
	}
	taskStatus = "completed"
	endTIme := time.Now()

	Conn.Model(&task).Where("id = ?", task.ID).Updates(TaskTracker{Comments: task.Comments, EndTime: endTIme, Status: taskStatus})

	return ValidResponse(200, taskUpdate, "success")
}

//DeleteTrackedTask deletes a task has not been tracked.
func DeleteTrackedTask(user User, task TaskTracker) interface{} {
	taskStatus := "Pending"
	if findTask := Conn.Where("id = ? AND user_id = ? AND status = ?", task.ID, user.ID, taskStatus).Find(&TaskTracker{}); findTask.Error != nil {
		return ValidResponse(403, "User not authorized to delete task", findTask.Error.Error())
	}

	if deleteTask := Conn.Where("id = ?", task.ID).Delete(TaskTracker{}); deleteTask.Error != nil {
		return ValidResponse(403, "Unable to delete Task", deleteTask.Error.Error())
	}
	return ValidResponse(200, "Delete Successfully", "success")
}

//StructureTaskAndUpdates structures the task and updates
func StructureTaskAndUpdates(task TaskTracker, updates []TaskTrackerUpdates) TaskUpdateResponseBody {
	var response TaskUpdateResponseBody
	response.Task = task
	response.Updates = updates

	return response
}

//UpdateTaskProgress update task progress
func UpdateTaskProgress(user User, taskUpdate TaskTrackerUpdates) interface{} {
	if taskUpdate.TaskID == 0 {
		return ValidResponse(403, "Empty Task ID", "error")
	}
	taskID := fmt.Sprint(taskUpdate.TaskID)
	var task TaskTracker
	task, err := GetTrackedTaskFromID(taskID)
	if err != nil {
		return ValidResponse(403, "Unable to get task from task_id", err.Error())
	}

	taskStatus := "in progress"
	if task.Status != taskStatus {
		return ValidResponse(403, "Only tasks in progress can be updated", "error")
	}

	if task.UserID != user.ID {
		return ValidResponse(403, "User not authorized to update task progress", "error")
	}

	if taskUpdate.ID == 0 && taskUpdate.Update == " " {
		return ValidResponse(403, "Empty task update", "error")
	}

	Conn.Create(&taskUpdate)

	return ValidResponse(200, taskUpdate, "success")
}
