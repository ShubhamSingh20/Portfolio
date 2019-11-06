package main

import (
	"os"
	"fmt"
	"log"

	db "github.com/ShubhamSingh20/Portfolio/db"
	router "github.com/ShubhamSingh20/Portfolio/router"
	portfolio "github.com/ShubhamSingh20/Portfolio/portfolio"
	blog "github.com/ShubhamSingh20/Portfolio/blog"

)

//InstalledApps will load all the apps which are being used
var InstalledApps = [...]string{
	blog.AppName,
	portfolio.AppName,	
}

func executeFromCommandLine() {

	// all the arguments of single command
	if len(os.Args) < 2 {
		fmt.Println("[-] No arguments provided")
		os.Exit(1)
	}

	programArgument := os.Args[1]
	
	switch programArgument {
		
	case "runserver":
		
		fmt.Println(
			"[+] Starting development server at ", 
			router.HTTPServer.Addr,
			"\nPress Ctrl+C to exit",
		)

		log.Fatal(router.HTTPServer.ListenAndServe())
		break

	case "createsuperuser":
		break
	
	case "cleardb":
		cleardb()
		break
	
	case "cleartb":
		break
	
	case "help", "h":
		fmt.Println(
			"[+] HELP MANUAL \n" +
			"runserver -- to run development server\n" +
			"createsuperuser -- to create superuser in db\n" +
			"cleardb -- clear the database of the program (will delete all the tables) \n" +
			"cleartb -- clear all the tables in database",
		)
		break
	
	default: 
		fmt.Println(programArgument, " no such argument avialable")
	}
}

func main() {
	executeFromCommandLine()
}

func cleardb() {
	dbConnection := db.ConnectDb()
	dbInst := dbConnection.Getdb()

	dbName := dbConnection.GetdbName()

	tableList, err := dbInst.Query(
		"select table_name FROM information_schema.tables where table_schema=?",
		dbName,
	)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	var tableNameList []string

	fmt.Println("[*] Following tables will be deleted ...")

	for tableList.Next() {
		var tableName string

		err = tableList.Scan(&tableName)

        if err != nil {
            panic(err.Error())
		}
		
		tableNameList = append(tableNameList, tableName)
		fmt.Println(tableName)
	}

	var sure string

	fmt.Println("[?] Are you sure ... (y/n)")
	fmt.Scanf("%s", &sure)

	switch sure {
	case "Y", "y", "Yes", "yes", "YES":
		break

	default:
		return 
	}

	for _, tableName := range tableNameList {
		dropStatement, err := dbInst.Prepare("DROP TABLE IF EXISTS " + tableName)

		if err != nil {
            panic(err.Error())
		}

		dropStatement.Exec()
		fmt.Println("[+] ",tableName," deleted.")
	}

	defer dbInst.Close()
}