package pkg

import (
  "github.com/taemon1337/cpm/config"
  "github.com/taemon1337/cpm/store"
)

type ContainerPackageManager struct {
  Version                 string
  Registries              []*Registry
  Store                   *store.GitStore
}

func NewContainerPackageManager(cfg *config.PackageManagerConfig) *ContainerPackageManager {
  return &ContainerPackageManager{
    Version:              cfg.Version,
    Registries:           make([]*Registry, 0),
    Store:                store.NewGitStore(cfg.StoreConfig),
  }
}
