package middleware

import (
	"net/http"
	"fmt"
)


//IsUserAuthenticated through a custom 404 page
var IsUserAuthenticated = func(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check for request ...
		handlerFunc(w, r)
		// check for response
		fmt.Println()
	}
}