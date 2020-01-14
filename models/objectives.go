package models

//CreateStrategicObjective handles creating a strategic objective
func CreateStrategicObjective(obj StrategicObjective, teamLead User) interface{} {
	var team Team
	if verifyTeamLead := Conn.Where("lead_id = ?", teamLead.ID).Find(&team); verifyTeamLead.Error != nil {
		return ErrorResponse(403, "User is not a Team Lead")
	}

	var strategicObjective StrategicObjective
	strategicObjective.Objective = obj.Objective
	strategicObjective.Team = team.Name
	strategicObjective.TeamID = team.ID
	strategicObjective.Status = "in progress"

	if createObj := Conn.Create(&strategicObjective); createObj.Error != nil {
		return ErrorResponse(403, createObj.Error.Error())
	}
	return ValidResponse(200, strategicObjective, "success")
}

//GetTeamLeadStrategicObj gets the team lead's strategic objectives
func GetTeamLeadStrategicObj(teamLead User) ([]StrategicObjective, error) {
	var myTeam Team
	if getMyTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&myTeam); getMyTeam.Error != nil {
		return nil, getMyTeam.Error
	}
	var allObj []StrategicObjective
	if findObjectives := Conn.Where("team_id = ?", myTeam.ID).Find(&allObj); findObjectives.Error != nil {
		return nil, findObjectives.Error
	}

	return allObj, nil
}

//GetTeamMemberStrategicObj gets a teams strategic objective
func GetTeamMemberStrategicObj(teamMember User) ([]StrategicObjective, error) {
	var member Members
	if getMyTeamLead := Conn.Where("member_id = ?", teamMember.ID).Find(&member); getMyTeamLead.Error != nil {
		return nil, getMyTeamLead.Error
	}
	var teamLead User
	teamLead.ID = member.TeamLeadID
	var allStrategivObjective []StrategicObjective
	allStrategivObjective, err := GetTeamLeadStrategicObj(teamLead)
	if err != nil {
		return nil, err
	}

	return allStrategivObjective, nil
}
