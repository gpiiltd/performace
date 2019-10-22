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
