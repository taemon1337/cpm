package pkg

import (
  "github.com/taemon1337/cpm/config"
  "github.com/taemon1337/cpm/store"
)

type ContainerPackageManager struct {
  Version                 string
  Registries              *RegistryList
  Store                   *store.Store
}

func NewContainerPackageManager(cfg *config.PackageManagerConfig) *ContainerPackageManager {
  return &ContainerPackageManager{
    Version:              cfg.Version,
    Registries:           &RegistryList{
      Registries:           make([]*Registry, 0),
    },
    Store:                store.NewStoreConfig(cfg.StoreType),
  }
}
