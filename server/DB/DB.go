package DB

import (
	"log"
	"os"

	"github.com/maulerrr/book-addict-server/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DATABASE_DSN")

	log.Println(dsn)

	DB, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(
			"Failed to connect to the database! \n",
			err,
		)
	}

	log.Println("Connected to database!")
	log.Println("Running migrations")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.BookTab{},
	)

	if err != nil {
		log.Fatal("Failed to connect to migrate! \n", err)
	}

	log.Println("Migrations done!")
}
