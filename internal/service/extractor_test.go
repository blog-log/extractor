package service

import (
	"context"
	"testing"
)

func Test_Extract__success_with_content(t *testing.T) {
	// setup
	repoUrl := "https://github.com/brandoncate-personal/blog-content"
	extractor := NewGitExtractor()

	// test
	repo, err := extractor.Extract(context.TODO(), repoUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	if repo.Data == nil {
		t.Errorf("expected not nil data")
	}
	if repo.Warning != nil {
		t.Errorf("expected nil warning")
	}
}

func Test_Extract__success_with_no_content(t *testing.T) {
	// setup
	repoUrl := "https://github.com/brandoncate-personal/blog-user"
	extractor := NewGitExtractor()

	// test
	repo, err := extractor.Extract(context.TODO(), repoUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	if repo.Data != nil {
		t.Errorf("expected nil data")
	}
	if repo.Warning == nil {
		t.Errorf("expected warning")
	}
}

func Test_Extract__error_not_found(t *testing.T) {
	// setup
	repoUrl := "https://github.com/brandoncate-personal/fake"
	extractor := NewGitExtractor()

	// test
	_, err := extractor.Extract(context.TODO(), repoUrl)
	if err == nil {
		t.Errorf("expected error")
	}

}
