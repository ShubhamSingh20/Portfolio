package main

import (
	"os"
	"fmt"
	"log"

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