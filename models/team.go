package models

//GetMyTeamInformations retrieves a team information for a user
func GetMyTeamInformations(user User) interface{} {
	var team Team
	var teamMember Members
	if findTeam := Conn.Where("member_id = ?", user.ID).Find(&teamMember); findTeam.Error != nil {
		return ValidResponse(404, team, "User does not belong to any team")
	}

	if findTeamInfo := Conn.Where("id = ?", teamMember.TeamID).Find(&team); findTeamInfo.Error != nil {
		return ValidResponse(404, team, "Unable to get team Information")
	}

	return ValidResponse(200, team, "success")
}

//GetPendingTeamRequests gets a team memeber team requests
func GetPendingTeamRequests(user User) interface{} {
	var invitations []TeamInvitation
	if getInvitations := Conn.Where("invitee_id = ?", user.ID).Find(&invitations); getInvitations.Error != nil {
		return ValidResponse(404, nil, "No Pending Invitations")
	}
	return ValidResponse(200, invitations, "success")
}

//AcceptInvitation accepts a new user team request
func AcceptInvitation(user User, teamID string) interface{} {
	var teamMember Members
	if findMember := Conn.Where("member_id = ?", user.ID).Find(&teamMember); findMember.Error == nil {
		return ErrorResponse(403, "User already belongs to a team")
	}

	var teamInfo Team
	if findTeam := Conn.Where("id = ?", teamID).Find(&teamInfo); findTeam.Error != nil {
		return ErrorResponse(403, "Invalid Team ID")
	}

	var createTeam Members
	createTeam.Team = teamInfo.Name
	createTeam.TeamID = teamInfo.ID
	createTeam.TeamLead = teamInfo.Lead
	createTeam.TeamLeadID = teamInfo.LeadID
	createTeam.Member = user.FullName
	createTeam.MemberID = user.ID

	Conn.Create(&createTeam)
	Conn.Where("team_id = ? AND invitee_id = ?", teamID, user.ID).Delete(TeamInvitation{})

	return ValidResponse(200, createTeam, "Successfully joined team")
}

//GetTeamReport gets a list of all team
func GetTeamReport() []Team {
	var teams []Team
	Conn.Find(&teams)
	return teams
}

//ValidateTeamLead checks if a team lead is actually the subordinate's team lead
func ValidateTeamLead(teamLeadID uint64, subordinateID uint64) bool {
	var team Members
	if findTeam := Conn.Where("team_lead_id = ? AND member_id = ?", teamLeadID, subordinateID).Find(&team); findTeam.Error != nil {
		return false
	}
	return true
}

//BehaviourTestResults gets and saves the behaviour tests results
func BehaviourTestResults(teamLead User, tests BehaviourTest) interface{} {
	validateTeamLead := ValidateTeamLead(teamLead.ID, tests.SubordinateID)
	if validateTeamLead != true {
		return ErrorResponse(403, "Unauthorized team lead. I don't know how you got here.")
	}
	Conn.Create(&tests)
	return ValidResponse(200, tests, "success")
}
