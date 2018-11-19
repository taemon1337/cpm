package store

import (
  "path"

  "gopkg.in/src-d/go-git.v4"
)

type GitStore struct {
  BasePath              string
}

func (g *GitStore) Fetch(name, url string) (*git.Repository, error) {
  repo, err := git.PlainClone(path.Join(g.BasePath, name), false, &git.CloneOptions{
    URL: url,
    RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
  })

  if err != nil {
    return nil, err
  }

  return repo, nil
}

func (g *GitStore) Remove(name string) error {
  return nil
}

func (g *GitStore) Checkout(name, branch string) error {
  return nil
}

func NewGitStore(cfg *StoreConfig) *GitStore {
  return &GitStore{
    BasePath:   cfg.BasePath,
  }
}
