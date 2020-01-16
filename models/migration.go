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
	// MigrateTeam()
	MigrateMembers()
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

//MigrateMembers migrates
func MigrateMembers() {
	var allMemberArray []Members
	stmt, err := OldDB.Query(`SELECT member, lead from members`)
	if err != nil {
		panic(err.Error())
	}
	var thisMember Members
	for stmt.Next() {
		stmt.Scan(&thisMember.MemberID, &thisMember.TeamLeadID)
		allMemberArray = append(allMemberArray, thisMember)
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

	var allTeam []Team
	Conn.Find(&allTeam)

	var newUseArray []User
	Conn.Find(&newUseArray)

	var newUserArray []User
	newUserArray = SortUser(allUsersArray, newUseArray)

	var teamMembers Members
	for _, eachMember := range allMemberArray {
		var member User
		member = GetUserFromArray(newUserArray, eachMember.ID)
		teamMembers.Member = member.FullName
		teamMembers.MemberID = member.ID
		var lead User
		lead = GetUserFromArray(newUserArray, eachMember.ID)
		teamMembers.TeamLeadID = lead.ID
		teamMembers.TeamLead = lead.FullName
		var team Team
		team = GetTeamFromArray(allTeam, lead.ID)
		teamMembers.Team = team.Name
		teamMembers.TeamID = team.ID

		Conn.Create(&teamMembers)
	}
}

//SortUser sorts
func SortUser(oldDB []User, newDB []User) []User {
	var sortedUsers []User
	for _, firstUsers := range oldDB {
		for _, secondUsers := range newDB {
			if firstUsers.Email == secondUsers.Email {
				sortedUsers = append(sortedUsers, secondUsers)
			}
		}
	}
	return sortedUsers
}

//GetUserFromArray retrieves user data
func GetUserFromArray(userArray []User, userID uint64) User {
	for _, user := range userArray {
		if user.ID == userID {
			return user
		}
	}
	return User{}
}

//GetTeamFromArray gets
func GetTeamFromArray(teamArray []Team, leadID uint64) Team {
	for _, team := range teamArray {
		if team.LeadID == leadID {
			return team
		}
	}
	return Team{}
}
