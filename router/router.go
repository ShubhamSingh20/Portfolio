package router

import (
	"net/http"
	"time"
	mux "github.com/gorilla/mux"
	errView "github.com/ShubhamSingh20/Portfolio/utils/view"
)

//Router global variable
var Router *mux.Router

//HTTPServer for listining request
var HTTPServer *http.Server

//HTTPServerClosed for listining request
var HTTPServerClosed = http.ErrServerClosed

func init() {
	Router = mux.NewRouter()

	Router.NotFoundHandler = http.HandlerFunc(errView.Custom404ErrorView)
	Router.MethodNotAllowedHandler = http.HandlerFunc(errView.Custom403ErrorView)
	
	Router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	HTTPServer = &http.Server{
		Handler:      Router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}