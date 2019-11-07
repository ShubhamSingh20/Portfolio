package auth

import (
	db "github.com/ShubhamSingh20/Portfolio/db"
)


var userModel = "" +
	"CREATE TABLE IF NOT EXISTS `auth_superuser` (" +
	"`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY," + 
	"`is_logged_in` BOOLEAN DEFAULT FALSE," +
	"`username` VARCHAR(100) NOT NULL UNIQUE," +
	"`hashpasswd` VARCHAR(255) DEFAULT NULL," +
	"`salt` VARCHAR(255) DEFAULT NULL" +
");"


//AuthenticateUser user validation
func AuthenticateUser(username, inputPassword string) bool {
	var hashpassword, salt string

	dbInst := db.ConnectDb().Getdb()

	stmt, err := dbInst.Query(
		"SELECT `hashpasswd`, `salt` FROM `auth_superuser` WHERE `username`= ? ;",
		username,
	)

	if err != nil {
		panic(err.Error())
	}

	for stmt.Next() {
		stmt.Scan(&hashpassword, &salt)
	}

	defer dbInst.Close()

	return authenticate(inputPassword, hashpassword, salt)

}

//CreateSuperUser takes in username & password and create superuser.
func CreateSuperUser(username, password string) {

	password, salt := getHashedPasswordAndSalt(password)

	dbInst := db.ConnectDb().Getdb()
	stmt, err := dbInst.Prepare(
		"INSERT INTO `auth_superuser`(username, hashpasswd, salt) " +
		"VALUES (?,?,?);",
	)

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(username, password, salt)

	defer dbInst.Close()

}

//GetModels get SQL code for all the models
func GetModels() []string {
	return []string{userModel}
}
