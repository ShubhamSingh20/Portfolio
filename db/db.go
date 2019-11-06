package db


import (
	"os"
	"log"
	"time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // import for using mysql-db support
)

var dbDriver = "mysql"
var dbUser =  os.Getenv("PORTFOLIO_DB_USER")
var dbPassWord =  os.Getenv("PORTFOLIO_DB_PASSWORD")
var dbName = "PortfolioDb"


//EnvDb for managing queries regarding statements
type EnvDb struct {
	db *sql.DB
}

//Getdb returns the original SQL connection object
func (envDb *EnvDb) Getdb() *sql.DB {
	return envDb.db
}


//GetdbName get the name of the db
func (envDb *EnvDb) GetdbName() string {
	return dbName
}

// ConnectDb establishes a connection between database create a
// new new connection for each time an db operation is to be done
func ConnectDb() *EnvDb {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassWord+"@/"+dbName)
	
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	db.SetConnMaxLifetime(time.Hour/2)

	err = db.Ping() ; if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	
	envDb := &EnvDb{db:db}

	return  envDb
}