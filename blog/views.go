package blog

import (
	"net/http"
	"path"
	"html/template"
	"fmt"
)


var (
	templatePath = path.Join("blog", "templates")
)

func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	
	adminLoginTemplate := template.Must(template.ParseFiles(
		path.Join(templatePath, "admin-login.html"),
	))
	err := adminLoginTemplate.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	
}

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username, password)

}


func adminDashboardHandler(w http.ResponseWriter, r *http.Request) {
}



func createPostHandler(w http.ResponseWriter, r *http.Request) {
}



func previewPostHandler(w http.ResponseWriter, r *http.Request) {
}



func editPostHandler(w http.ResponseWriter, r *http.Request) {
}



func deletePostHandler(w http.ResponseWriter, r *http.Request) {
}



/*
CREATE USER 'shubham'@'localhost' IDENTIFIED BY '#include';
GRANT ALL PRIVILEGES ON * . * TO 'shubham'@'localhost';
*/