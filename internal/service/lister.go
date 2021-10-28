package service

import (
	"github.com/blog-log/extractor/internal/model"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Lister func(repo string) ([]*model.GitFile, error)

func GitList(repoUrl string) ([]*model.GitFile, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repoUrl,
	})
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	var files []*model.GitFile
	if err = tree.Files().ForEach(func(f *object.File) error {
		content, err := f.Contents()
		if err != nil {
			return err
		}

		files = append(files, &model.GitFile{
			Name:    f.Name,
			Content: content,
		})

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
