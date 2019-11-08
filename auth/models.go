package auth

import (
	db "github.com/ShubhamSingh20/Portfolio/db"
)


//Credentials for passing around user cred
type Credentials  struct {
	Username string
	Password string
}


var userModel = "" +
	"CREATE TABLE IF NOT EXISTS `auth_superuser` (" +
	"`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY," + 
	"`is_logged_in` BOOLEAN DEFAULT FALSE," +
	"`username` VARCHAR(100) NOT NULL UNIQUE," +
	"`hashpasswd` VARCHAR(255) DEFAULT NULL," +
	"`last_logged_in` datetime DEFAULT NULL," +
	"`salt` VARCHAR(255) DEFAULT NULL" +
");"


//AuthenticateUser user validation
func AuthenticateUser(cred *Credentials) bool {
	var id int
	var hashpassword, salt string

	dbInst := db.ConnectDb().Getdb()
	defer dbInst.Close()

	stmt, err := dbInst.Query(
		"SELECT `id`, `hashpasswd`, `salt` FROM `auth_superuser` WHERE `username`= ? ;",
		cred.Username,
	)

	if err != nil {
		panic(err.Error())
	}

	stmt.Scan(&id, &hashpassword, &salt)

	userValid := authenticate(cred.Password, hashpassword, salt)

	setUserToLoggedIn := func(id int){
		upStmt, err := dbInst.Prepare(
			"UPDATE `auth_superuser` SET `is_logged_in`=true, `last_logged_in`=now() " +
			"WHERE `id`=? ;",
		)

		if err != nil {
			panic(err.Error())
		}

		upStmt.Exec(id)
	}

	if userValid {
		setUserToLoggedIn(id)
	}

	return userValid
}

//CreateSuperUser takes in username & password and create superuser.
func CreateSuperUser(cred *Credentials) {

	passwordHashed, salt := getHashedPasswordAndSalt(cred.Password)

	dbInst := db.ConnectDb().Getdb()
	stmt, err := dbInst.Prepare(
		"INSERT INTO `auth_superuser`(username, hashpasswd, salt) " +
		"VALUES (?,?,?);",
	)

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(cred.Username, passwordHashed, salt)

	defer dbInst.Close()

}

//GetModels get SQL code for all the models
func GetModels() []string {
	return []string{userModel}
}
