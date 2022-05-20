package config

import "database/sql"

// var db *gorm.DB

// func Connect() *gorm.DB {
// 	var err error
// 	dbUser := "root"
// 	dbPassword := "17081998"
// 	dbHost := "localhost"
// 	dbName := "quanlydatvexetructuyen"
// 	dsn := dbUser + ":" + dbPassword + "@tcp" + "(" + dbHost + ")/" + dbName
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("Error connecting to database : error=%v", err)
// 		return nil
// 	}
// 	db.AutoMigrate(&models.Xe{})
// 	return db
// }

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "17081998"
	dbHost := "localhost:3306"
	dbName := "quanlydatvexetructuyen"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
