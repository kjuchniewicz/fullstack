package seed

import (
	"log"

	"github.com/kjuchniewicz/fullstack/api/models"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Nickname: "Vuko Drakkainnen",
		Email:    "vuko.d@gmail.com",
		Password: "passwd",
	},
	{
		Nickname: "Ivi Drakkainnen",
		Email:    "ivi.d@gmail.com",
		Password: "passwd",
	},
	{
		Nickname: "Wak Drakkainnen",
		Email:    "wak.d@gmail.com",
		Password: "passwd",
	},
}

var posts = []models.Post{
	{
		Title:   "Tutuł 1",
		Content: "Siemka 1",
	},
	{
		Title:   "Tutuł 2",
		Content: "Siemka 2",
	},
	{
		Title:   "Tutuł 3",
		Content: "Siemka 3",
	},
}

func Load(db *gorm.DB) {
	var err error
	if db.Debug().Migrator().HasTable(&models.User{}) || db.Debug().Migrator().HasTable(&models.Post{}) {
		log.Println("Usuwam tabele")
		err = db.Debug().Migrator().DropTable(&models.User{}, &models.Post{})
		if err != nil {
			log.Fatalf("nie można usunąć tabeli: %v", err)
		}
	}
	// err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	// if err != nil {
	// 	log.Fatalf("nie można usunąć tabeli: %v", err)
	// }
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{})
	log.Println("Migruje tabele")
	if err != nil {
		log.Fatalf("nie można zmigrować tebel: %v", err)
	}

	// err = db.Debug().Migrator().CreateConstraint(&models.Post{},"User")
	//err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// log.Fatalf("błąd dołączania klucza obcego: %v", err)
	// }

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("nie może zasiać tabeli użytkowników: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("nie może zasiać tabeli postów: %v", err)
		}
	}
}
