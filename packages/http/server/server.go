package server

import (
	"context"
	"errors"
	"example/service/packages/http/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pavlico/cacheRedis"
)

type ServerStarter interface {
	Serve(server *http.Server)
}

func createChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return stopCh, func() {
		close(stopCh)
	}
}

func start(server *http.Server) {
	log.Println("service started")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	} else {
		log.Println("service stopped gracefully")
	}
}

func shutdown(ctx context.Context, server *http.Server) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		log.Println("service shutdowned")
	}
}

func Serve() {
	routes := http.NewServeMux()
	routes.HandleFunc("/user/", cacheRedis.HandlerCacheMiddleware(handler.GetUserById()))
	s := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	go start(s)
	stopCh, closeCh := createChannel()
	defer shutdown(context.Background(), s)

	defer closeCh()

	log.Println("notified:", <-stopCh)
}

func CheckConnections() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
}
