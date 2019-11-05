package blog

import router "github.com/ShubhamSingh20/Portfolio/router"



//AppName denotes the name of the entire app
const (
	AppName = "blog"
)

func init() {
	mux := router.Router

	mux.HandleFunc("/admin/", adminLoginHandler).Methods("GET")
	mux.HandleFunc("/admin/", authenticateUser).Methods("POST")
	mux.HandleFunc("/admin-dashboard/", adminDashboardHandler)

	blogRouter := mux.PathPrefix("/blog").Subrouter()		

	blogRouter.HandleFunc("/create-post/", createPostHandler)
	blogRouter.HandleFunc("/preview-post/", previewPostHandler)
	blogRouter.HandleFunc("/edit-post/", editPostHandler)
	blogRouter.HandleFunc("/delete-post/", deletePostHandler)
}