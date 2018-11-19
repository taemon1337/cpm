package store

import (
  "path"

  "github.com/taemon1337/cpm/pkg"
  "github.com/taemon1337/cpm/config"
  "gopkg.in/src-d/go-git.v4"
)

type GitStore struct {
  BasePath              string
}

func (g *GitStore) fetch(p *pkg.Package) (string, error) {
  repo, err := git.PlainClone(path.Join(g.BasePath, p.Name), false, &git.CloneOptions{
    URL: p.Repo,
    RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
  })

  if err != nil {
    return nil, err
  }

  return repo
}

func NewGitStore(cfg *StoreConfig) *GitStore {
  return &GitStore{
    BasePath:   cfg.BasePath,
  }
}
