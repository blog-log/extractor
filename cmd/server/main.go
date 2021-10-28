package main

import (
	"context"

	serverv1 "github.com/blog-log/extractor/internal/server"
)

func main() {
	ctx := context.Background()

	server := serverv1.NewExtractorServer()

	server.Run(ctx)
}
