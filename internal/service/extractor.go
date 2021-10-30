package service

import (
	"context"

	"github.com/blog-log/extractor/internal/adapter"
	"github.com/blog-log/extractor/internal/model"
)

type Extractor func(ctx context.Context, repoUrl string) (*model.Repo, error)

type GitExtractor struct {
	lister      Lister
	fileAdapter adapter.FileAdapter
	repoAdapter adapter.RepoAdapter
}

func NewGitExtractor() *GitExtractor {
	return &GitExtractor{
		lister:      GitList,
		fileAdapter: adapter.GitFileAdapter,
		repoAdapter: adapter.GitRepoAdapter,
	}
}

func (e *GitExtractor) Extract(ctx context.Context, repoUrl string) (*model.Repo, error) {
	// list files in repo
	gitFiles, err := e.lister(repoUrl)
	if err != nil {
		return nil, err
	}

	// filter and enrich gitFile data with frontmatter (if exists)
	files, err := e.fileAdapter(gitFiles)
	if err != nil {
		return nil, err
	}

	// adapt data to repo object for consumer return
	repo := e.repoAdapter(repoUrl, files)
	return repo, nil
}
