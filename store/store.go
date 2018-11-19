package store

import (
  "log"
  "github.com/taemon1337/cpm/pkg"
)

type Store interface {
  func fetch(*pkg.Package) (string, error)
  func remove(*pkg.Package) error
  func checkout(*pkg.Package, branch string) error
}

func NewStore(cfg *StoreConfig) *Store {
  if cfg.Type == "git" {
    return NewGitStore(cfg)
  } else {
    log.Fatal("Currently the only store type supported is 'git'")
  }
}
