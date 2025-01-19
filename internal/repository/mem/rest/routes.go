package rest

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/ghostriderdev/housierBack/internal/rest/routing"
)

type IRouteRepository interface {
	AddRoute(r routing.Route) error
	GetRoute(path string) (routing.Route, bool)
	GetAllRoutes() *[]routing.Route
}

type RouteRepository struct {
	routes map[string]routing.Route
}

func New() *RouteRepository {
	return &RouteRepository{
		routes: make(map[string]routing.Route),
	}
}

func (repo *RouteRepository) AddRoute(r routing.Route) error {
	hash := generateHash(r.Path, r.Method)
	repo.routes[hash] = r
	return nil
}

func (repo *RouteRepository) GetRoute(path, method string) (routing.Route, bool) {
	hash := generateHash(path, method)
	route, exists := repo.routes[hash]
	return route, exists
}

func (repo *RouteRepository) GetAllRoutes() *[]routing.Route {
	routes := make([]routing.Route, 0, len(repo.routes))
	for _, route := range repo.routes {
		routes = append(routes, route)
	}
	return &routes
}

func generateHash(path, method string) string {
	h := sha1.New()
	pathId := fmt.Sprintf("%s_%s", method, path)
	h.Write([]byte(pathId))
	return hex.EncodeToString(h.Sum(nil))
}
