package main

import (
	"log"
	"fmt"

	router "github.com/ShubhamSingh20/Portfolio/router"
	portfolio "github.com/ShubhamSingh20/Portfolio/portfolio"
	blog "github.com/ShubhamSingh20/Portfolio/blog"

)

//InstalledApps will load all the apps which are being used
var InstalledApps = [...]string{
	blog.AppName,
	portfolio.AppName,	
}

func main() {
	
	fmt.Println("Starting server ...")
	log.Fatal(router.HTTPServer.ListenAndServe())


}