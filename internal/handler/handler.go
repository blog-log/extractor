package handler

import (
	"github.com/blog-log/extractor/internal/service"
)

type ExtractHandler struct {
	extractor service.Extractor
}

func NewExtractHandler() *ExtractHandler {
	gitExtractor := service.NewGitExtractor()

	return &ExtractHandler{
		extractor: gitExtractor.Extract,
	}
}
