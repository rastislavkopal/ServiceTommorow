package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a SqLite3 database and save the reference to `Database` struct.
// func Init() *gorm.DB {
// 	db, err := gorm.Open("sqlite3", "./../gorm.db")
// 	if err != nil {
// 		fmt.Println("db err: (Init) ", err)
// 	}
// 	db.DB().SetMaxIdleConns(10)
// 	//db.LogMode(true)
// 	DB = db
// 	return DB
// }

// Opening a MariaDB database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open("mysql", "test_user:pwd123@(mariadb:3306)/appka?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	fmt.Println("MariaDB connected oh yeaah")
	//db.LogMode(true)
	DB = db
	return DB
}

// This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("sqlite3", "./../gorm_test.db")
	if err != nil {
		fmt.Println("db err: (TestDBInit) ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

// Delete the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("./../gorm_test.db")
	return err
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
