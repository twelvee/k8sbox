// Package main - k8sbox entrypoint
package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/twelvee/k8sbox/cmd/k8sbox/internal/commands"
)

func main() {
	var (
		ctx, _ = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	)
	if err := commands.NewRootCommand().ExecuteContext(ctx); err != nil {
		log.Fatalf("Failed to execute command: %s", err.Error())
	}
}
