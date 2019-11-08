package main

import (
	"os"
	"fmt"
	"log"

	auth "github.com/ShubhamSingh20/Portfolio/auth" 
	util "github.com/ShubhamSingh20/Portfolio/utils" 
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

//AppModelSchema contains the schema of all the models
var AppModelSchema = [][]string{
	auth.GetModels(),
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
		)

		util.SetupRedisCache()
		fmt.Println(
			"[+] Started redis cache " + 
			"\nPress Ctrl+C to exit",
		)
		if err := router.HTTPServer.ListenAndServe(); err != nil && err != router.HTTPServerClosed {
			log.Fatalf("Could not listen : %v\n", err)
		}

		fmt.Println("[+] Server Stopped")
		
		util.CloseRedisConnection()
		fmt.Println("[+] Closed Redis Connection")

		break
		

	case "migrate":
		dbConnection := db.ConnectDb()

		for _, schema := range AppModelSchema {
			dbConnection.Migrate(schema...)
		}

		fmt.Println(
			"[+] It seems that the datbase has been updated successfully.",
		)

		defer dbConnection.Getdb().Close()
		break


	case "createsuperuser":
		cred := &auth.Credentials{}

		fmt.Println("Enter following details:")
		fmt.Printf("|-> Enter username: ")
		fmt.Scanf("%s", &cred.Username)

		fmt.Printf("|->Enter password: ")
		fmt.Scanf("%s", &cred.Password)

		auth.CreateSuperUser(cred)
		fmt.Printf("[+] %s created successfully.\n", cred.Username)
		break
		

	case "cleardb":
		dbConnection := db.ConnectDb()
		tableNameList := dbConnection.GetTableListInDb()

		if len(tableNameList) < 1 {
			fmt.Println("[-] Database is empty")	
		} else {
			fmt.Println("[*] Following tables will be deleted ...")
			for _, tableName := range tableNameList {
				fmt.Println("-> ", tableName)
			}

			areYouSure(dbConnection.ClearDb, "Tables deleted successfully.")
		}
		
		defer dbConnection.Getdb().Close()

		break

	
	case "cleartb":
		dbConnection := db.ConnectDb()
		tableNameList := dbConnection.GetTableListInDb()

		if len(tableNameList) < 1 {
			fmt.Println("[-] Database is empty")	
		} else {
			fmt.Println("[*] All the rows of following tables will be deleted ...")
			for _, tableName := range tableNameList {
				fmt.Println("-> ", tableName)
			}

			areYouSure(dbConnection.ClearTb, "Tables emptied successfully.")
		}
		
		defer dbConnection.Getdb().Close()
		
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


func areYouSure(f func(), msg string) {
	var sure string

	fmt.Println("[?] Are you sure ... (y/n)")
	fmt.Scanf("%s", &sure)

	switch sure {
	case "Y", "y", "Yes", "yes", "YES":
		f()
		fmt.Println("[+] ", msg)
		break

	default:
		return 
	}

}
