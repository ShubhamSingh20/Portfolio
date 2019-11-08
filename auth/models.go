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
	"`salt` VARCHAR(255) DEFAULT NULL" +
");"


//AuthenticateUser user validation
func AuthenticateUser(cred *Credentials) bool {
	var hashpassword, salt string

	dbInst := db.ConnectDb().Getdb()

	stmt, err := dbInst.Query(
		"SELECT `hashpasswd`, `salt` FROM `auth_superuser` WHERE `username`= ? ;",
		cred.Username,
	)

	if err != nil {
		panic(err.Error())
	}

	stmt.Scan(&hashpassword, &salt)
	
	defer dbInst.Close()

	return authenticate(cred.Password, hashpassword, salt)

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
