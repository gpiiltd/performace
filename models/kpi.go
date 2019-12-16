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
	var members Members
	if findMember := Conn.Where("team_lead_id = ? AND member_id = ?", teamLead.ID, member.ID).Find(&members); findMember.Error != nil {
		return ErrorResponse(403, "Not authorized to view user KPI")
	}
	var kpiInfo []KPI
	if findKPI := Conn.Where("employee_id = ? AND start_date = ?", member.ID, month).Find(&kpiInfo); findKPI.Error != nil {
		return ErrorResponse(403, "No KPI information for that user.")
	}
	return ValidResponse(200, kpiInfo, "Success")
}
