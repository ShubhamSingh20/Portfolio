package auth

import (
	db "github.com/ShubhamSingh20/Portfolio/db"
)


var userModel = "" +
	"CREATE TABLE IF NOT EXISTS `auth_superuser` (" +
	"`auth_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY," + 
	"`auth_username` VARCHAR(100) NOT NULL," +
	"`auth_hashpasswd` VARCHAR(255) DEFAULT NULL," +
	"`auth_salt` VARCHAR(255) DEFAULT NULL" +
");"


//CreateSuperUser takes in username & password and create superuser.
func CreateSuperUser(username, password string) {

	password, salt := getHashedPasswordAndSalt(password)

	dbInst := db.ConnectDb().Getdb()
	stmt, err := dbInst.Prepare(
		"INSERT INTO `auth_superuser`(auth_username, auth_hashpasswd, auth_salt) " +
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
