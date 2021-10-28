package adapter

import (
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/blog-log/extractor/internal/model"
)

type FileAdapter func(files []*model.GitFile) ([]*model.File, error)

func GitFileAdapter(gitFiles []*model.GitFile) ([]*model.File, error) {
	var files []*model.File
	for _, gitFile := range gitFiles {

		var matter model.Matter
		_, err := frontmatter.Parse(strings.NewReader(gitFile.Content), &matter)
		if err != nil {
			return nil, err
		}

		if matter.Title != "" {
			files = append(files, &model.File{
				Name:  gitFile.Name,
				Title: matter.Title,
			})
		}
	}

	return files, nil
}
