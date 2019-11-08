package server

import (
	"fmt"
	"log"
	"os"
	router "github.com/ShubhamSingh20/Portfolio/router"
)


//StartServer start the server
func StartServer() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	fmt.Println(
		"[+] Starting development server at ", 
		router.HTTPServer.Addr,
	)

	SetupRedisCache()
	fmt.Println(
		"[+] Started redis cache " + 
		"\nPress Ctrl+C to exit",
	)
	if err := router.HTTPServer.ListenAndServe(); err != nil && err != router.HTTPServerClosed {
		logger.Fatalf("Could not listen : %v\n", err)
	}

	fmt.Println("[+] Server Stopped")
	
	CloseRedisConnection()
	fmt.Println("[+] Closed Redis Connection")

}