package models

import (
	"encoding/json"
)

//AddNewTeamMember adds a new member to the team
func AddNewTeamMember(member User, teamLead User) interface{} {
	var team Team
	hasTeam, team := MyTeam(teamLead)
	if hasTeam != true {
		return ErrorResponse(403, "User does not have a team")
	}

	var invitation TeamInvitation
	invitation.TeamName = team.Name
	invitation.TeamID = team.ID
	invitation.TeamLead = teamLead.FullName
	invitation.TeamLeadID = teamLead.ID
	invitation.InviteeID = member.ID
	invitation.InviteeName = member.FullName
	invitation.Status = "pending"

	Conn.Create(&invitation)

	return ValidResponse(200, teamLead, "success")
}

//MyTeam checks if i have a team and return my team if true
func MyTeam(user User) (bool, Team) {
	var myTeam Team
	if findTeam := Conn.Where("lead_id = ?", user.ID).Find(&myTeam); findTeam.Error != nil {
		return false, myTeam
	}
	return true, myTeam
}

//LeadHasTeam checks if a team Lead has a team
func LeadHasTeam(teamLead User) (bool, string) {
	var team Team
	if findTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&team); findTeam.Error != nil {
		return false, "No team created"
	}
	return true, team.Name
}

//AddNewTeam adds a new member to the team
func AddNewTeam(teamLead User, team Team) interface{} {
	//validate user role with code
	status, message := ValidateUserRoleAPI(teamLead.ID, 66)
	if status != true {
		return ErrorResponse(501, message)
	}

	code, status := VerifiTeamRoleStatus(message)
	if code != 200 {
		return ErrorResponse(501, "Error verifying API response at: "+message)
	}

	if status != true {
		return ErrorResponse(401, "User not authorized to create team.")
	}

	var checkTeam Team
	if findTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&checkTeam); findTeam.Error == nil {
		return ErrorResponse(200, "Users not allowed to own more than 1 team")
	}

	var newTeam Team
	newTeam.Lead = teamLead.FullName
	newTeam.LeadID = teamLead.ID
	newTeam.Name = team.Name
	newTeam.Department = teamLead.Department
	newTeam.DepartmentID = teamLead.DepartmentID

	Conn.Create(&newTeam)

	return ValidResponse(200, newTeam, "success")
}

//VerifiTeamRoleStatus verifies the status of the response from API
func VerifiTeamRoleStatus(body string) (uint64, bool) {
	// log.Println(body)
	verifiBody := []byte(body)
	type verificationResponse struct {
		Code uint64 `json:"code"`
		Body bool   `json:"body"`
	}
	var responseBody verificationResponse
	err := json.Unmarshal(verifiBody, &responseBody)
	if err != nil {
		return 501, false
	}
	return responseBody.Code, responseBody.Body
}

//GetTeamInfo gets a user team information
func GetTeamInfo(teamLead User) interface{} {
	var teamMembers []Members
	if findMembers := Conn.Where("team_lead_id = ?", teamLead.ID).Find(&teamMembers); findMembers.Error != nil {
		return ValidResponse(401, nil, "No Members in Team")
	}

	var allUsers []User
	Conn.Find(&allUsers)

	var allMembersArray []User

	for _, member := range teamMembers {
		for _, user := range allUsers {
			if user.ID == member.MemberID {
				allMembersArray = append(allMembersArray, user)
			}
		}
	}

	return ValidResponse(200, allMembersArray, "success")
}

//MyPendingTeamInfo retrieves all information for pending team members
func MyPendingTeamInfo(teamLead User) interface{} {
	var myTeam Team
	Conn.Where("lead_id = ?", teamLead.ID).Find(&myTeam)

	var invitations []TeamInvitation
	if findInvitation := Conn.Where("team_id = ?", myTeam.ID).Find(&invitations); findInvitation.Error != nil {
		return ValidResponse(200, nil, "success")
	}

	type invitationResponse struct {
		User             User   `json:"user"`
		InvitationStatus string `json:"status"`
	}
	var inviteStatus invitationResponse
	var inviteStatusArray []invitationResponse

	var allUsers []User
	Conn.Find(&allUsers)

	for _, user := range allUsers {
		for _, invite := range invitations {
			if invite.InviteeID == user.ID {
				inviteStatus.User = user
				inviteStatus.InvitationStatus = invite.Status

				inviteStatusArray = append(inviteStatusArray, inviteStatus)
			}
		}
	}

	return ValidResponse(200, inviteStatusArray, "success")
}

//DeletePendingTeamMember deletes pending member from the system
func DeletePendingTeamMember(uid string) interface{} {
	if deleteInvitation := Conn.Where("invitee_id = ?", uid).Delete(&TeamInvitation{}); deleteInvitation.Error != nil {
		return ErrorResponse(401, "Unable to delete Team Invitation record")
	}
	return ValidResponse(200, "Delete Successful", "success")
}

//IsMyTeamLead checks if a member is a team lead
func IsMyTeamLead(member User, teamLead User) bool {
	var members Members
	if isMyTeamLead := Conn.Where("member_id = ? AND team_lead_id = ?", member.ID, teamLead.ID).Find(&members); isMyTeamLead.Error != nil {
		return false
	}
	return true
}
