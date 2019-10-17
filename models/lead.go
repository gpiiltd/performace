package models

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
)

//AddNewTeamMember adds a new member to the team
func AddNewTeamMember(member User) interface{} {
	isTeamLead, err := http.Get(beego.AppConfig.String("coreapi"))
	if err != nil {
		return ErrorResponse(200, "User is not a team lead: %s"+err.Error())
	}

	var teamMember Members
	teamMember.Member = member.FullName
	teamMember.MemberID = member.ID

	Conn.Find(&teamMember)

	return ValidResponse(200, isTeamLead, "success")
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
		return ErrorResponse(501, "Error verifying API response")
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
	var team Team
	if findTeam := Conn.Where("lead_id = ?", teamLead.ID).Find(&team); findTeam.Error != nil {
		return ErrorResponse(404, "Team not Found")
	}
	return ValidResponse(200, team, "success")
}
