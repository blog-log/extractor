package service

import "github.com/blog-log/extractor/internal/model"

type Adapter func(repo string, files []*model.File) *model.Repo

func FileAdapter(repo string, files []*model.File) *model.Repo {
	var docs []*model.Document
	for _, file := range files {
		docs = append(docs, &model.Document{
			Path:  file.Name,
			Title: file.Title,
		})
	}

	data := &model.Repo{
		Repo:   repo,
		Branch: "main",
		Data:   docs,
	}

	return data
}
