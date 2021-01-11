package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/enkhalifapro/users-serverless/internal/users/api"

	"go.uber.org/zap"
)

func main() {
	// Setup logger.
	logger, _ := zap.NewProduction()

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()
	u, err := url.Parse("localhost:9090")
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", "localhost:9090", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	h := api.BuildHTTPHandler(logger)

	wg.Add(1)
	go func() {
		defer wg.Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Info(fmt.Sprintf("HTTP server listening on host %s", u.Host))
			errc <- http.ListenAndServe("localhost:9090", h)
		}()
		<-ctx.Done()
		logger.Info("shutting down HTTP server at",
			zap.String("host", u.Host))

		// Shutdown gracefully with a 30s timeout.
		var _, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
	}()

	// Wait for signal.
	logger.Info("exiting", zap.Error(<-errc))

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Info("exited")
}
