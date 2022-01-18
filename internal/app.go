package internal

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/di/config"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	container := di.NewContainer(cfg)

	server := createServer(cfg, container)

	log.Println("starting http server at", cfg.Port)
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
}

func createServer(cfg *config.Config, container *di.Container) *http.Server {
	var handler http.Handler
	handler = createRouter(container)
	handler = handlers.CombinedLoggingHandler(os.Stdout, handler)
	handler = cors.AllowAll().Handler(handler)

	return &http.Server{Addr: ":" + strconv.Itoa(int(cfg.Port)), Handler: handler}
}
