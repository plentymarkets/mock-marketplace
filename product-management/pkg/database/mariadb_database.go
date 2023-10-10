package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type MariaDBDatabase struct {
	db *gorm.DB
}

func NewMariaDBDatabase() MariaDBDatabase {
	return MariaDBDatabase{
		db: createConnection(),
	}
}

func (db *MariaDBDatabase) GetConnection() *gorm.DB {
	return db.db
}

func createConnection() *gorm.DB {

	//dsn := fmt.Sprintf(
	//	"%s:%s@tcp(database:3306)/%s",
	//	os.Getenv("MARIADB_USER"),
	//	os.Getenv("MARIADB_PASSWORD"),
	//	os.Getenv("MARIADB_USER"))

	dsn := "root:myrootpassword@tcp(host.docker.internal:3306)/mydb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
