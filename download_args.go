package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type downloadArgs struct {
	repoURL    string
	outDirPath string
}

func newDownloadArgs(rawRepoURL, outDirPath string) (*downloadArgs, error) {
	repoURL, err := parseRepoURL(rawRepoURL)
	if err != nil {
		return nil, errors.WithMessage(err, "url error")
	}

	if outDirPath == "" {
		goSrcDir := filepath.Join(os.Getenv("GOPATH"), "src")
		repoDir := filepath.Join(repoURL.Hostname(), strings.TrimSuffix(repoURL.RequestURI(), ".git"))

		outDirPath = filepath.Join(goSrcDir, repoDir)
	}

	res := &downloadArgs{
		repoURL:    repoURL.String(),
		outDirPath: outDirPath,
	}

	return res, nil
}

func parseRepoURL(rawURL string) (*url.URL, error) {
	repo, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.WithMessagef(err, "bad url")
	}

	repo.RawQuery = ""
	repo.RawFragment = ""

	if repo.Host == "" && repo.Path == "" {
		return nil, fmt.Errorf("full url is required")
	}

	if repo.Scheme == "" {
		repo.Scheme = "http"
		repo, _ = url.Parse(repo.String())
	}

	if !strings.HasSuffix(repo.Path, ".git") {
		repo.Path += ".git"
	}

	return repo, nil
}
