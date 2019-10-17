package models

import (
	"time"

	"github.com/astaxie/beego"
	//Mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

//DBConfig holds database connection object
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

//DB hold new instance of databse objects
var DB = new(DBConfig)

//Conn db
var Conn *gorm.DB

func init() {
	DB.Host = beego.AppConfig.String("databaseHost")
	DB.User = beego.AppConfig.String("databaseUsername")
	DB.Password = beego.AppConfig.String("databasePassword")
	DB.Database = beego.AppConfig.String("databaseName")

	conn, err := gorm.Open("mysql", DB.User+":"+DB.Password+"@/"+DB.Database+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	Conn = conn

	CreateTables()
}

//CreateTables creates all database tables
func CreateTables() {
	Conn.AutoMigrate(&User{})
	Conn.AutoMigrate(&Team{})
	Conn.AutoMigrate(&Members{})
	return
}

//ErrorResponseData sends response data if situation is false.
type ErrorResponseData struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

//Model holds the default gorm models
type Model struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

//User struct shows models for users
type User struct {
	Model
	FullName     string `gorm:"type:varchar(100)" json:"full_name"`
	JobTitle     string `gorm:"type:varchar(100)" json:"job_title"`
	Gender       string `gorm:"type:varchar(100)" json:"gender"`
	Location     string `gorm:"type:varchar(100)" json:"location"`
	Number       string `gorm:"type:varchar(100)"  json:"number"`
	Email        string `gorm:"type:varchar(100); unique_index" json:"email"`
	Role         uint64 `gorm:"type:int(10)" json:"role"`
	Department   string `gorm:"department" json:"department"`
	DepartmentID uint64 `gorm:"type:int(10)" json:"department_id"`
	Image        string `gorm:"type:varchar(100)" json:"image"`
}

//Team holds the object for creating a new team on the system
type Team struct {
	Model
	Name         string `gorm:"type:varchar(100)" json:"name"`
	LeadID       uint64 `gorm:"type:int(10)" json:"lead_id"`
	Lead         string `gorm:"type:varchar(100)" json:"lead"`
	Department   string `gorm:"type:varchar(100)" json:"department"`
	DepartmentID uint64 `gorm:"type:int(10)" json:"department_id"`
}

//Members holds struct of team members
type Members struct {
	Model
	Team       string `gorm:"type:varchar(100)" json:"team"`
	TeamID     uint64 `gorm:"type:int(10)" json:"team_id"`
	TeamLead   string `gorm:"type:varchar(100)" json:"team_lead"`
	TeamLeadID uint64 `gorm:"type:int(10)" json:"team_lead_id"`
	Member     string `gorm:"type:varchar(100)" json:"member"`
	MemberID   uint64 `gorm:"type:int(10)" json:"member_id"`
}
