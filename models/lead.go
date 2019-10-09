package models

import (
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
