package pkg

import (
  "github.com/taemon1337/cpm/config"
)

type ContainerPackageManager struct {
  Version                 string
  Registries              *RegistryList
}

func NewContainerPackageManager(cfg *config.PackageManagerConfig) *ContainerPackageManager {
  return &ContainerPackageManager{
    Version:              cfg.Version,
    Registries:           &RegistryList{
      Registries:           make([]*Registry, 0),
    },
  }
}
