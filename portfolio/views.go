package portfolio

import (
	"net/http"
	"path"

 	view "github.com/ShubhamSingh20/Portfolio/utils"

)

var (
	templatePath = path.Join("portfolio", "templates")
)

//HomePageHandler handles the request for homepage of portfolio
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, "portfolio/index.html", nil)
}