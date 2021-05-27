package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kjuchniewicz/fullstack/api/controllers"
	"github.com/kjuchniewicz/fullstack/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Błąd podczas pobierania .env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("TestDbDriver")
	TestDbUser := os.Getenv("TestDbDriver")
	TestDbPassword := os.Getenv("TestDbDriver")
	TestDbHost := os.Getenv("TestDbDriver")
	TestDbPort := os.Getenv("TestDbDriver")
	TestDbName := os.Getenv("TestDbDriver")

	if TestDbDriver == "sqlite" {
		server.DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			fmt.Printf("Nie można się połączyć z baza danych %s", TestDbDriver)
			log.Fatal("To jest błąd:", err)
		} else {
			fmt.Printf("Jesteśmy podłączeni do bazy danych %s\n", TestDbDriver)
		}
	}
	if TestDbDriver == "mssql" {
		// sqlserver://Kartal:1903@127.0.0.14:1433?database=Besiktas
		DBURL := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", TestDbUser, TestDbPassword, TestDbHost, TestDbPort, TestDbName)
		server.DB, err = gorm.Open(sqlserver.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Nie można się połączyć z baza danych %s", TestDbDriver)
			log.Fatal("To jest błąd:", err)
		} else {
			fmt.Printf("Jesteśmy podłączeni do bazy danych %s", TestDbDriver)
		}
	}
	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", TestDbUser, TestDbPassword, TestDbHost, TestDbPort, TestDbName)
		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", TestDbDriver)
		}
	}
	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", TestDbHost, TestDbPort, TestDbUser, TestDbName, TestDbPassword)
		server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	var err error
	if server.DB.Migrator().HasTable(&models.User{}) {
		err = server.DB.Migrator().DropTable(&models.User{})
		if err != nil {
			return err
		}
	}
	err = server.DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	log.Printf("Pomyślnie odświeżono tabelę")
	return nil
}

func seedOneUser() (models.User, error) {
	refreshUserTable()

	user := models.User{
		Nickname: "TestPet",
		Email:    "test.pet@gmail.com",
		Password: "passwd",
	}
	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("Nie można zasiać tabeli użytkowników: %v", err)
	}
	return user, nil
}

func seedUsers() error {

	users := []models.User{
		models.User{
			Nickname: "Steven victor",
			Email:    "steven@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Kenny Morris",
			Email:    "kenny@gmail.com",
			Password: "password",
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndPostTable() error {
	var err error
	if server.DB.Migrator().HasTable(&models.User{}) || server.DB.Migrator().HasTable(&models.Post{}) {
		err = server.DB.Migrator().DropTable(&models.User{}, &models.Post{})
		if err != nil {
			return err
		}
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		return err
	}
	log.Printf("Pomyślnie odświeżono tabelę")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {

	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}
	user := models.User{
		Nickname: "Sam Phil",
		Email:    "sam@gmail.com",
		Password: "password",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}
	post := models.Post{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func seedUsersAndPosts() ([]models.User, []models.Post, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Post{}, err
	}
	var users = []models.User{
		models.User{
			Nickname: "Steven victor",
			Email:    "steven@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Magu Frank",
			Email:    "magu@gmail.com",
			Password: "password",
		},
	}
	var posts = []models.Post{
		models.Post{
			Title:   "Title 1",
			Content: "Hello world 1",
		},
		models.Post{
			Title:   "Title 2",
			Content: "Hello world 2",
		},
	}

	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return users, posts, nil
}
