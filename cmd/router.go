package main

import (
	"net/http"

	"github.com/ghostriderdev/housierBack/pkg/utils"
	"github.com/ghostriderdev/housierBack/routes"
	"github.com/ghostriderdev/housierBack/templates"
)

func router(w http.ResponseWriter, r *http.Request) {
	route := utils.ExtractPath(r.URL, 0)

	switch route {
	case "auth":
		routes.AuthRouter(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(templates.PageNotFound))
	}
}
