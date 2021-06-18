package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Application returned error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	httpClient := NewHttpClient()

	jmsepathClient, err := NewJmsepathClient("result")
	if err != nil {
		return err
	}

	ctx := context.Background()
	opaClient, err := NewOpaClient(ctx)
	if err != nil {
		return err
	}

	handlerClient := NewHandlerClient(httpClient, jmsepathClient, opaClient, "http://opa-test:8181/v1/data/nginx/authz")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	errGroup, ctx := errgroup.WithContext(ctx)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(stopCh)

	address := ":8082"
	srv := &http.Server{
		Addr: address,
	}

	http.HandleFunc("/proxy", handlerClient.OpaProxyHandler)
	http.HandleFunc("/rego", handlerClient.OpaRegoHandler)

	errGroup.Go(func() error {
		fmt.Printf("Starting server: %s\n", address)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	select {
	case <-stopCh:
		break
	case <-ctx.Done():
		break
	}

	cancel()

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeoutCancel()

	errGroup.Go(func() error {
		if err := srv.Shutdown(timeoutCtx); err != nil {
			return err
		}

		return nil
	})

	fmt.Println("Shutting down server")

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("Error group returned error: %w", err)
	}

	fmt.Println("Server gracefully shutdown")

	return nil
}
