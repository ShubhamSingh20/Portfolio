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


//GetTableListInDb get list of all the tables in the database
func (envDb *EnvDb) GetTableListInDb() []string {
	
	tableList, err := envDb.db.Query(
		"SELECT table_name FROM information_schema.tables WHERE table_schema=?",
		dbName,
	)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	var tableNameList []string

	for tableList.Next() {
		var tableName string

		err = tableList.Scan(&tableName)

        if err != nil {
            panic(err.Error())
		}
		
		tableNameList = append(tableNameList, tableName)
	}

	return tableNameList
}


//ClearDb delete all the tables in the database
func (envDb *EnvDb) ClearDb() {

	for _, tableName := range envDb.GetTableListInDb() {
		dropStatement, err := envDb.db.Prepare("DROP TABLE IF EXISTS " + tableName)

		if err != nil {
            panic(err.Error())
		}

		dropStatement.Exec()
	}
}

//ClearTb delete all rows in the tables in the database
func (envDb *EnvDb) ClearTb() {

	for _, tableName := range envDb.GetTableListInDb() {
		deleteStatement, err := envDb.db.Prepare("DELETE FROM " + tableName)

		if err != nil {
            panic(err.Error())
		}

		deleteStatement.Exec()
	}
}



//Migrate provide models and it will migrate them
func(envDb *EnvDb) Migrate(tableSchema ...string) {
	for _, schema := range tableSchema {
		stmt, err := envDb.db.Prepare(schema)
		
		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec()

		if err != nil {
			panic(err.Error())
		}
	}
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