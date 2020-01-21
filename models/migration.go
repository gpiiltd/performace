package models

import (
	"database/sql"

	//Mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/astaxie/beego"
)

//OldDB for old database using manual queries
var OldDB *sql.DB

//NewDB with gorm
var NewDB *gorm.DB

//SetupOldDatabase sets up the data needed to mine the old DB
func SetupOldDatabase() {
	var DB = new(DBConfig)
	DB.Host = beego.AppConfig.String("databaseHost")
	DB.User = beego.AppConfig.String("databaseUsername")
	DB.Password = beego.AppConfig.String("databasePassword")
	DB.Database = beego.AppConfig.String("olddatabaseName")
	conn, err := sql.Open("mysql", DB.User+":"+DB.Password+"@/"+DB.Database+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	OldDB = conn

}

//SetupNewDatabase sets us new database
func SetupNewDatabase() {
	var DB = new(DBConfig)
	DB.Host = beego.AppConfig.String("databaseHost")
	DB.User = beego.AppConfig.String("databaseUsername")
	DB.Password = beego.AppConfig.String("databasePassword")
	DB.Database = beego.AppConfig.String("newdatabaseName")
	conn, err := gorm.Open("mysql", DB.User+":"+DB.Password+"@/"+DB.Database+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	NewDB = conn
}

//UpdateUserInfo updates
func UpdateUserInfo() {
	var allUsersArray []User
	NewDB.Find(&allUsersArray)
	for _, user := range allUsersArray {
		Conn.Create(&user)
	}
	return
}

//StartMining starts the process of extracting the data from old DB
func StartMining() {
	MigrateTeam()
	// MigrateMembers()
	return
}

//MigrateTeam blah
func MigrateTeam() {
	var allTeamArray []Team
	stmt, err := OldDB.Query(`SELECT name, lead from teams`)
	if err != nil {
		panic(err.Error())
	}
	var thisTeam Team
	for stmt.Next() {
		stmt.Scan(&thisTeam.Name, &thisTeam.LeadID)
		allTeamArray = append(allTeamArray, thisTeam)
	}

	var allUsersArray []User
	stmts, err := OldDB.Query(`SELECT ID, name from companyEmployees`)
	if err != nil {
		panic(err.Error())
	}
	var user User
	for stmts.Next() {
		stmts.Scan(&user.ID, &user.FullName)
		allUsersArray = append(allUsersArray, user)
	}

	for _, tempUser := range allUsersArray {
		for _, tempTeam := range allTeamArray {
			if tempUser.ID == tempTeam.LeadID {
				CreateTeams(tempTeam, tempUser)
			}
		}
	}

	return
}

//CreateTeams creates  user team
func CreateTeams(team Team, lead User) {
	var teamLead User
	if findLead := Conn.Where("full_name = ?", lead.FullName).Find(&teamLead); findLead.Error != nil {
		panic(findLead.Error.Error())
	}
	var newTeam Team
	newTeam.Name = team.Name
	newTeam.LeadID = teamLead.ID
	newTeam.Lead = teamLead.FullName
	newTeam.Department = teamLead.Department
	newTeam.DepartmentID = teamLead.DepartmentID

	Conn.Create(&newTeam)
	return
}

//MigrateMembers gets a list of all old users
func MigrateMembers() {
	var allOldTeamMemberArray []Members
	allOldTeamMemberArray = GetAllOldMembers()

	var teamLead User
	var teamMember User
	var teamData Team
	var newMember Members
	var allNewMembers []Members
	var counter uint64
	counter = 0
	for _, member := range allOldTeamMemberArray {
		counter = counter + 1
		newMember.ID = counter
		teamLead = GetNewUser(member.TeamLeadID)
		newMember.TeamLead = teamLead.FullName
		newMember.TeamLeadID = teamLead.ID
		teamMember = GetNewUser(member.MemberID)
		newMember.Member = teamMember.FullName
		newMember.MemberID = teamMember.ID
		teamData = GetNewTeam(member.TeamID)
		newMember.Team = teamData.Name
		newMember.TeamID = teamData.ID

		allNewMembers = append(allNewMembers, member)
		Conn.Create(&newMember)
	}

	return
}

//GetNewTeam get a new team data
func GetNewTeam(teamID uint64) Team {
	var oldTeam Team
	stmt, _ := OldDB.Query(`SELECT name, lead, leadName from teams where id = ?`, teamID)
	for stmt.Next() {
		stmt.Scan(&oldTeam.Name, &oldTeam.LeadID, &oldTeam.Lead)
	}

	var newTeam Team
	Conn.Where("name = ?", oldTeam.Name).Find(&newTeam)
	return newTeam
}

//GetAllOldMembers gets a list of all old members
func GetAllOldMembers() []Members {
	var allOldMembers []Members
	stmt, _ := OldDB.Query(`SELECT member, lead, team from members`)
	var user Members
	for stmt.Next() {
		stmt.Scan(&user.MemberID, &user.TeamLeadID, &user.TeamID)
		allOldMembers = append(allOldMembers, user)
	}

	return allOldMembers
}

//GetNewUser tansforms the old user information to the new information
func GetNewUser(uid uint64) User {
	stmt, _ := OldDB.Query(`SELECT ID, name, email from companyEmployees where ID = ?`, uid)
	var user User
	for stmt.Next() {
		stmt.Scan(&user.ID, &user.FullName, &user.Email)
	}

	var newUser User
	Conn.Where("email = ?", user.Email).Find(&newUser)
	return newUser
}
