package router

import (
	"net/http"
	"time"
	mux "github.com/gorilla/mux"
)

//Router global variable
var Router *mux.Router

//HTTPServer for listining request
var HTTPServer *http.Server

func init() {
	Router = mux.NewRouter()

	Router.NotFoundHandler = http.HandlerFunc(Custom404ErrorView)
	Router.MethodNotAllowedHandler = http.HandlerFunc(Custom403ErrorView)
	
	Router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	HTTPServer = &http.Server{
		Handler:      Router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}