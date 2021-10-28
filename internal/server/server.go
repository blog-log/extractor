package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/blog-log/extractor/internal/handler"
	"github.com/gorilla/mux"
)

type Server interface {
	Run(ctx context.Context) error
}

type ExtractorServer struct{}

func NewExtractorServer() *ExtractorServer {
	return &ExtractorServer{}
}

func (s *ExtractorServer) Run(ctx context.Context) error {

	rootRouter := mux.NewRouter().StrictSlash(true)

	// setup ExtractHandler
	handler := handler.NewExtractHandler()

	// define routes
	rootRouter.HandleFunc("/extract", handler.Extract).Methods("POST")

	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", rootRouter); err != nil {
		return err
	}

	return nil
}
