package models

import "errors"

//GetReport retrieves the user monthly report
func GetReport(tokenString string, month string) interface{} {
	status, user := GetUserFromTokenString(tokenString)
	if status != 200 {
		LogError(errors.New("Invalid Token String"))
		return ValidResponse(403, "Invalid Token String", "error")
	}

	// var teamMembers Members
	var allTeamMembers []Members

	if findAllMembers := Conn.Where("team_lead_id = ?", user.ID).Find(&allTeamMembers); findAllMembers.Error != nil {
		LogError(findAllMembers.Error)
		return ValidResponse(403, findAllMembers.Error.Error(), "error")
	}

	var thisMonthKPIs []KPI
	if findThisMonthKPIs := Conn.Where("start_date = ?", month).Find(&thisMonthKPIs); findThisMonthKPIs.Error != nil {
		LogError(findThisMonthKPIs.Error)
		return ValidResponse(403, findThisMonthKPIs.Error.Error(), "error")
	}

	var kpiReport []KPIReport
	var thisReport KPIReport

	for _, teamMembers := range allTeamMembers {
		for _, kpi := range thisMonthKPIs {
			if teamMembers.MemberID == kpi.EmployeeID {
				thisReport.MemberKPI = kpi
				thisReport.TeamMember = teamMembers

				kpiReport = append(kpiReport, thisReport)
			}
		}
	}

	response := ValidResponse(200, kpiReport, "fine")
	return response
}
