package config

import (
  "github.com/taemon1337/cpm/store"
)

var DefaultPackageManagerConfig = &PackageManagerConfig{
  Version:      "1.0.0",
  Verbose:      "1",
  StoreConfig:  store.DefaultStoreConfig,
}



