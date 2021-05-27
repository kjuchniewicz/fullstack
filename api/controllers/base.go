package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjuchniewicz/fullstack/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "sqlite" {
		server.DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			fmt.Printf("Nie można się połączyć z baza danych %s", Dbdriver)
			log.Fatal("To jest błąd:", err)
		} else {
			fmt.Printf("Jesteśmy podłączeni do bazy danych %s\n", Dbdriver)
		}
	}
	if Dbdriver == "mssql" {
		// sqlserver://Kartal:1903@127.0.0.14:1433?database=Besiktas
		DBURL := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(sqlserver.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Nie można się połączyć z baza danych %s", Dbdriver)
			log.Fatal("To jest błąd:", err)
		} else {
			fmt.Printf("Jesteśmy podłączeni do bazy danych %s", Dbdriver)
		}
	}
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Nasłuchuje na porcie 8088")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
