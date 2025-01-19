package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/ghostriderdev/housierBack/cfg"
	"github.com/ghostriderdev/housierBack/internal/controller"
	"github.com/ghostriderdev/housierBack/internal/mapper"
	repository "github.com/ghostriderdev/housierBack/internal/repository/mem/data"
	"github.com/ghostriderdev/housierBack/internal/repository/mem/rest"
	"github.com/ghostriderdev/housierBack/internal/rest/routing"
	"github.com/ghostriderdev/housierBack/internal/service"
)

type App struct {
	Config       *cfg.ApiConfig
	RoutingTable *rest.RouteRepository
}

func main() {
	config, err := cfg.LoadConfig()
	if err != nil {
		panic(err)
	}

	routingTable := rest.New()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		loadRoutes(routingTable)
	}()

	wg.Wait()

	app := App{
		Config:       config,
		RoutingTable: routingTable,
	}

	startServer(&app)
}

func loadRoutes(routingTable *rest.RouteRepository) {
	userMemRepo := repository.NewUserRepository()
	userMapper := mapper.NewUserMapper()
	authService := service.NewUserService(userMemRepo, userMapper)
	authController := &controller.AuthController{
		Service: authService,
	}

	routes := []routing.Route{
		{
			Path:    "/auth/signup",
			Method:  http.MethodPost,
			Handler: authController.Signup,
		},
		{
			Path:    "/",
			Method:  http.MethodGet,
			Handler: authController.Signup,
		},
	}

	for _, route := range routes {
		routingTable.AddRoute(route)
	}
}

func startServer(app *App) {
	serverPort := app.Config.ServerCfg.Port

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", serverPort),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			router(w, r, app)
		}),
	}

	log.Printf("Running server on http://127.0.0.1%s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
