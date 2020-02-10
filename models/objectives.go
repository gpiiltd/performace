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
	strategicObjective.Year = obj.Year

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

//DeleteStrategicObjective deletes a team's strategic objective using it's ID
func DeleteStrategicObjective(teamLead User, objectiveID int) ValidResponseData {
	var team Team
	if findMyTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&team); findMyTeam.Error != nil {
		return ValidResponse(403, "User not a team lead", findMyTeam.Error.Error())
	}

	var strategicObjective StrategicObjective
	if findObjective := Conn.Where("team_id = ? AND id = ?", team.ID, objectiveID).Find(&strategicObjective); findObjective.Error != nil {
		return ValidResponse(403, "User not authorized o delete strategic objective", findObjective.Error.Error())
	}

	Conn.Where("id = ?", objectiveID).Delete(&StrategicObjective{})
	return ValidResponse(200, "Successfully deleted strategic Objective.", "success")
}

//MarkObjComplete marks a strategic objective as complete.
func MarkObjComplete(user User, objectiveID int) ValidResponseData {
	objectiveStatus := "completed"

	var myTeam Team
	if findMyTeam := Conn.Where("lead_id = ?", user.ID).Find(&myTeam); findMyTeam.Error != nil {
		return ValidResponse(403, "User does not have a team", findMyTeam.Error.Error())
	}

	var strategicObj StrategicObjective
	if findObjective := Conn.Where("team_id = ? AND id = ?", myTeam.ID, objectiveID); findObjective.Error != nil {
		return ValidResponse(403, "Invalid Strategic Objective ID", findObjective.Error.Error())
	}

	Conn.Model(&strategicObj).Where("id = ? AND team_id = ?", objectiveID, myTeam.ID).Updates(StrategicObjective{Status: objectiveStatus})

	return ValidResponse(200, strategicObj, "success")
}
