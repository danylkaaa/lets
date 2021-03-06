package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/lets-cli/lets/cmd"
	"github.com/lets-cli/lets/env"
	"github.com/lets-cli/lets/logging"
)

func main() {
	ctx := getContext()

	logging.InitLogging(env.IsDebug())

	rootCmd := cmd.CreateRootCommand(ctx, os.Stdout, GetVersion())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func getContext() context.Context {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := <-ch
		logging.Log.Printf("lets: signal received: %s", sig)
		cancel()
	}()

	return ctx
}
