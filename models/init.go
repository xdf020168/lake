package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/merico-dev/lake/config"
	"github.com/merico-dev/lake/models/domainlayer/code"
	"github.com/merico-dev/lake/models/domainlayer/crossdomain"
	"github.com/merico-dev/lake/models/domainlayer/devops"
	"github.com/merico-dev/lake/models/domainlayer/ticket"
	"github.com/merico-dev/lake/models/domainlayer/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func init() {
	connectionString := config.V.GetString("DB_URL")
	if config.V.Get("TEST") == "true" {
		connectionString = "merico:merico@tcp(localhost:3306)/lake_test"
	}
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)

	Db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("ERROR: >>> Mysql failed to connect")
		panic(err)
	}
	migrateDB()
}

func migrateDB() {
	err := Db.AutoMigrate(
		&Task{},
		&Notification{},
		&Pipeline{},
		&user.User{},
		&code.Repo{},
		&code.Commit{},
		&code.Pr{},
		&code.Note{},
		&code.RepoCommit{},
		&ticket.Board{},
		&ticket.Issue{},
		&ticket.BoardIssue{},
		&ticket.BoardSprint{},
		&ticket.Changelog{},
		&ticket.Sprint{},
		&ticket.SprintIssue{},
		&ticket.SprintIssueBurndown{},
		&devops.Job{},
		&devops.Build{},
		&ticket.Worklog{},
		&crossdomain.BoardRepo{},
		&crossdomain.IssueCommit{},
	)
	if err != nil {
		panic(err)
	}
	// TODO: create customer migration here
}
