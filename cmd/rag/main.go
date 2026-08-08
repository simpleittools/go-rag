package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/simpleittools/go-rag/app"
	"github.com/simpleittools/go-rag/config"
)

func main() {
	//. - setup the app
	// - setup config
	// - setup llm client
	// - setup the read eval print loop (repl)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx, config.Load()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
