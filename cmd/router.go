package main

import (
	"net/http"
)

const RouteNotFound = "Route not found"

func router(w http.ResponseWriter, r *http.Request, app *App) {
	result, isPresent := app.RoutingTable.GetRoute(r.URL.Path, r.Method)

	if !isPresent {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(RouteNotFound))
		return
	}

	result.Handler(w, r)
}
