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
			
	// '/' for url in the end of the regex
	uuidV1Regex := "/^[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i" + "}/" 

	blogRouter.HandleFunc("/create-post/", createPostHandler)
	blogRouter.HandleFunc("/edit-post/{slug:" + uuidV1Regex, editPostHandler)
	blogRouter.HandleFunc("/delete-post/{slug:" + uuidV1Regex, deletePostHandler)
	blogRouter.HandleFunc("/preview-post/{postname:^[a-zA-Z0-9]{4,50}$}/", previewPostHandler)
}