package adapter

import (
	"testing"

	"github.com/blog-log/extractor/internal/model"
)

func Test_GitFileAdapter__success_one(t *testing.T) {
	// setup
	files := []*model.GitFile{{
		Name: "fake-name",
		Content: `
---
title: fake-title
extra: fake-extra
---

# Content Title
bla bla bla
		`,
	}}

	// test
	_, err := GitFileAdapter(files)
	if err != nil {
		t.Errorf(err.Error())
	}

}
