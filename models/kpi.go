package models

import (
	"strconv"
)

//AssignNewKPI assigns a new KPI to a user
func AssignNewKPI(kpi KPI, user User) interface{} {
	if findTeamMatch := Conn.Where("team_lead_id = ? AND member_id = ?", user.ID, kpi.EmployeeID).Find(&Members{}); findTeamMatch.Error != nil {
		return ErrorResponse(403, "Unauthorized. Not a member of your Team")
	}

	kpi.TeamLead = user.FullName
	kpi.TeamLeadID = user.ID
	kpi.Status = "pending"
	Conn.Create(&kpi)
	return ValidResponse(200, kpi, "Successfully assigned KPI")
}

//GetKPIFromIDString gets a user kpi from kpi id (string)
func GetKPIFromIDString(kpiID string) (KPI, error) {
	var kpi KPI
	if findKPI := Conn.Where("id = ?", kpiID).Find(&kpi); findKPI.Error != nil {
		return kpi, findKPI.Error
	}
	return kpi, nil
}

//DeleteKPI deletes a user KPi
func DeleteKPI(kpi KPI, teamLead User) interface{} {
	// if deleteKPI := Conn.Where("id = ? AND team_lead_id = ?", kpi.ID, teamLead.ID).Delete(&KPI{}); deleteKPI.Error != nil {
	// 	return ErrorResponse(403, "Unauthorized, User not authorized to delete KPI")
	// }
	return ValidResponse(200, kpi, "Success")
}

//TeamMemberKPIReport gets the kpi report of a particular user for a particular month
func TeamMemberKPIReport(teamLead User, month int, member User) interface{} {
	var kpiInfo []KPI
	if findKPI := Conn.Where("employee_id = ? AND start_date = ?", member.ID, month).Find(&kpiInfo); findKPI.Error != nil {
		return ErrorResponse(403, "No KPI information for that user.")
	}
	var behaviour BehaviourTest
	behaviour, _ = GetBehaviourFromMonth(month, member.ID)
	type kpiInfoResponse struct {
		KPI       []KPI         `json:"kpi"`
		Behaviour BehaviourTest `json:"behaviour_test"`
	}

	var kpiInfoResp kpiInfoResponse
	kpiInfoResp.KPI = kpiInfo
	kpiInfoResp.Behaviour = behaviour
	return ValidResponse(200, kpiInfoResp, "Success")
}

//GetKPIInfo gets all kpi Information
func GetKPIInfo(kpiID int) interface{} {
	var kpi KPI
	kpi, err := GetKPIInfoID(kpiID)
	if err != nil {
		return ErrorResponse(403, "Unable to get KPI from ID")
	}
	monthInt, _ := strconv.Atoi(kpi.StartDate)

	var behaviour BehaviourTest
	behaviour, _ = GetBehaviourFromMonth(monthInt, kpi.EmployeeID)

	type kpiInfoResponse struct {
		KPI       KPI           `json:"kpi"`
		Behaviour BehaviourTest `json:"behaviour_test"`
	}

	var kpiInfoResp kpiInfoResponse
	kpiInfoResp.KPI = kpi
	kpiInfoResp.Behaviour = behaviour

	return ValidResponse(200, kpiInfoResp, "success")
}

//GetBehaviourFromMonth gets behavouor from KPI
func GetBehaviourFromMonth(month int, userID uint64) (BehaviourTest, error) {
	var behaviour BehaviourTest
	if findBehaviour := Conn.Where("month = ? AND subordinate_id = ?", month, userID).Find(&behaviour); findBehaviour.Error != nil {
		return behaviour, findBehaviour.Error
	}
	return behaviour, nil
}

//GetKPIInfoID gets kpi information from id
func GetKPIInfoID(kpiID int) (KPI, error) {
	var kpi KPI
	if findKPI := Conn.Where("id = ?", kpiID).Find(&kpi); findKPI.Error != nil {
		return kpi, findKPI.Error
	}

	return kpi, nil
}

//GetKPITasks gets all tasks belonging to a KPI
func GetKPITasks(kpiID int) ([]Task, error) {
	var tasks []Task
	if findTask := Conn.Where("kpi_id = ?", kpiID).Find(&tasks); findTask.Error != nil {
		return tasks, findTask.Error
	}

	return tasks, nil
}

//GetAllTask gets all task from the system
func GetAllTask() []Task {
	var allTask []Task
	Conn.Find(&allTask)
	return allTask
}

//ValidTaskObject checks if task struct is valid
func ValidTaskObject(task Task) (bool, Task) {
	if task.KPIID == 0 {
		return false, task
	}
	if task.Task == "" {
		return false, task
	}
	return true, task
}

//CreateKPITask creates a new kpi tasks for a user
func CreateKPITask(task Task, user User) (Task, bool) {
	var validatedTask Task
	status, validatedTask := ValidTaskObject(task)
	if status != true {
		return task, status
	}
	validatedTask.User = user.FullName
	validatedTask.UserID = user.ID
	validatedTask.Status = "in progress"
	Conn.Create(&validatedTask)
	return validatedTask, true
}

//GetTaskFromID gets a task from the id
func GetTaskFromID(taskID int) (Task, error) {
	var task Task
	if findTask := Conn.Where("id = ?", taskID).Find(&task); findTask.Error != nil {
		return task, findTask.Error
	}
	return task, nil
}

//MarkTaskAsComplete marks a task complete by the user
func MarkTaskAsComplete(user User, task Task) interface{} {
	var thisTask Task
	if findTask := Conn.Where("id = ? AND user_id = ?", task.ID, task.UserID).Find(&thisTask); findTask.Error != nil {
		return ErrorResponse(403, "Not authorized to edit Task information.")
	}

	taskStatus := "completed"
	Conn.Model(&task).Where("id = ?", task.ID).Updates(Task{Status: taskStatus})
	return ValidResponse(200, task, "success")
}

//ScoreKPComment scores a kpi information with a supervisor comment
func ScoreKPComment(kpi KPI, teamLead User) interface{} {
	var kpiObject KPI
	if findKPI := Conn.Where("id = ?", kpi.ID).Find(&kpiObject); findKPI.Error != nil {
		return ErrorResponse(403, "No KPI with ID")
	}
	if kpiObject.TeamLeadScore != 0 {
		return ErrorResponse(403, "KPI has already been scored")
	}
	if validateKPI := Conn.Where("id = ? AND team_lead_id = ?", kpi.ID, teamLead.ID).Find(&kpiObject); validateKPI.Error != nil {
		return ErrorResponse(403, "User not subordinate's team lead")
	}
	status := "reviewed"
	if scoreKPI := Conn.Model(&kpi).Where("id = ?", kpi.ID).Updates(KPI{TeamLeadScore: kpi.TeamLeadScore, TeamLeadComment: kpi.TeamLeadComment, Status: status}); scoreKPI.Error != nil {
		return ErrorResponse(403, "Error when scoring KPI")
	}
	return ValidResponse(200, kpi, "success")
}
