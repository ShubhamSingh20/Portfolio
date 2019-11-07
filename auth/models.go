package auth


var userModel = "" +
	"CREATE TABLE IF NOT EXISTS `auth_superuser` (" +
	"`auth_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY," + 
	"`auth_email` VARCHAR(100) NOT NULL," +
	"`auth_hashpasswd` VARCHAR(255) DEFAULT NULL," +
	"`auth_salt` VARCHAR(255) DEFAULT NULL" +
");"


//GetModels get SQL code for all the models
func GetModels() []string {
	return []string{userModel}
}