package portfolio

import router "github.com/ShubhamSingh20/Portfolio/router"


//AppName denotes the name of the entire app
const (
	AppName = "portfolio"
)

func init() {
	mux := router.Router
	mux.HandleFunc("/", HomePageHandler).Methods("GET")
}