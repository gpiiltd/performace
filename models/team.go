package models

import (
	"log"
	"sync"
)

//GetMyTeamInformations retrieves a team information for a user
func GetMyTeamInformations(user User) interface{} {
	// var team Team
	var teamMember []Members
	if findTeam := Conn.Where("member_id = ?", user.ID).Find(&teamMember); findTeam.Error != nil {
		return ValidResponse(404, teamMember, "User does not belong to any team")
	}

	// if findTeamInfo := Conn.Where("id = ?", teamMember.TeamID).Find(&team); findTeamInfo.Error != nil {
	// 	return ValidResponse(404, team, "Unable to get team Information")
	// }

	return ValidResponse(200, teamMember, "My Team Information")
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
func AcceptInvitation(user User, invitationID string) interface{} {
	log.Println(invitationID)
	var teamMembers []Members
	Conn.Where("member_id = ?", user.ID).Find(&teamMembers)
	if len(teamMembers) > 2 {
		return ErrorResponse(403, "User already belongs to more than 2 teams.")
	}

	var invitationInfo TeamInvitation
	if findInvitation := Conn.Where("id = ?", invitationID).Find(&invitationInfo); findInvitation.Error != nil {
		return ErrorResponse(403, "Invalid Invitation ID")
	}

	var teamInfo Team
	if findTeam := Conn.Where("id = ?", invitationInfo.TeamID).Find(&teamInfo); findTeam.Error != nil {
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
	Conn.Where("team_id = ? AND invitee_id = ?", teamInfo.ID, user.ID).Delete(TeamInvitation{})

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

//TeamLeadHasTeam checks if a user has a team
func TeamLeadHasTeam(teamLead User) ValidationResponseData {
	var team Team
	if findMyTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&team); findMyTeam.Error != nil {
		return ValidationResponse(200, false)
	}
	return ValidationResponse(200, true)
}

//DeleteTeamFunc deletes a particular team from the system using the team id
func DeleteTeamFunc(teamLead User) ValidResponseData {
	var team Team
	if findMyTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&team); findMyTeam.Error != nil {
		return ValidResponse(403, findMyTeam.Error.Error(), "error")
	}
	var teamMembers []Members
	Conn.Where("team_id = ?", team.ID).Find(&teamMembers)
	if len(teamMembers) > 0 {
		return ValidResponse(403, "Team still has members in it. Delete Members", "error")
	}
	Conn.Where("lead_id = ?", teamLead.ID).Delete(&Team{})
	return ValidResponse(200, "Successfully deleted team.", "success")
}

//DeletePendingTeamInvitation deletes a pending team invitation from the system using the team id
func DeletePendingTeamInvitation(teamLead User, invitationID int) ValidResponseData {
	var team Team
	if findMyTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&team); findMyTeam.Error != nil {
		return ValidResponse(403, "User not a team lead", findMyTeam.Error.Error())
	}
	var pendingInvitation TeamInvitation
	if findPendingInvitation := Conn.Where("team_lead_id = ? AND invitee_id = ?", teamLead.ID, invitationID).Find(&pendingInvitation); findPendingInvitation.Error != nil {
		return ValidResponse(403, "User not authorized to delete invitation. ", findPendingInvitation.Error.Error())
	}
	Conn.Where("invitee_id = ?", invitationID).Delete(&TeamInvitation{})
	return ValidResponse(200, "Successfully deleted pending invitation.", "success")
}

//DeleteTeamMemberFunc deletes a pending team invitation from the system using the team id
func DeleteTeamMemberFunc(teamLead User, memberID int) ValidResponseData {
	var team Team
	if findMyTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&team); findMyTeam.Error != nil {
		return ValidResponse(403, "User not a team lead", findMyTeam.Error.Error())
	}
	var teamMember Members
	if findMember := Conn.Where("team_lead_id = ? AND member_id = ?", teamLead.ID, memberID).Find(&teamMember); findMember.Error != nil {
		return ValidResponse(403, "User not authorized to delete member. ", findMember.Error.Error())
	}
	Conn.Where("member_id = ? AND team_lead_id = ?", memberID, teamLead.ID).Delete(&Members{})
	return ValidResponse(200, "Successfully deleted member.", "success")
}

//GetNonMembers retrieves an array of other users that are not in User's team
func GetNonMembers(teamLead User) []User {
	var teamMembers []Members
	if findMembers := Conn.Where("team_lead_id = ?", teamLead.ID).Find(&teamMembers); findMembers.Error != nil {
		LogError(findMembers.Error)
		return nil
	}

	var allUsers []User
	Conn.Find(&allUsers)
	userArrayLength := len(allUsers)

	var sortedArray []User

	var wg sync.WaitGroup
	wg.Add(userArrayLength)

	for i := 0; i < userArrayLength; i++ {
		go func(i int) {
			defer wg.Done()
			var thisUser User
			thisUser = allUsers[i]
			for _, member := range teamMembers {
				if thisUser.ID != member.ID {
					sortedArray = append(sortedArray, thisUser)
				}
			}
		}(i)
	}

	wg.Wait()

	return sortedArray

}

//AppendUserdataFromArrayID extends an operation for func GetNonMembers
func AppendUserdataFromArrayID(allUsers []User, userID uint64) User {
	for _, user := range allUsers {
		if userID != user.ID {
			return user
		}
	}

	return User{}
}
