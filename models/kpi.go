package models

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

//GetKPIForTheMonth gets the kpi information for a particular month
// func GetKPIForTheMonth(kpiRequest KPIRequest, authenticatedUser User) interface{} {
// 	myTeamLead := IsMyTeamLead()
// }
