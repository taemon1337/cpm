package config

import (
  "github.com/taemon1337/cpm/store"
)

type PackageManagerConfig struct {
  Version               string
  Verbose               string
  StoreConfig           *store.StoreConfig
}

func NewPackageManagerConfig() *PackageManagerConfig {
  return &PackageManagerConfig{
    Version:            DefaultPackageManagerConfig.Version,
    Verbose:            DefaultPackageManagerConfig.Verbose,
    StoreConfig:        store.DefaultStoreConfig,
  }
}
