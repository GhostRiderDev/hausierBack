package routes

import (
	"log"
	"net/http"

	"github.com/ghostriderdev/housierBack/pkg/utils"
	"github.com/ghostriderdev/housierBack/templates"
)

func AuthRouter(w http.ResponseWriter, r *http.Request) {
	route := utils.ExtractPath(r.URL, 1)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	switch route {
	case "signup":
		log.Println("Signup****")
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(templates.PageNotFound))
	}
}
