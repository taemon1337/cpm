package store

import (
  "os"
  "path"
)

var HomeDir = os.Getenv("HOME")
var DefaultStoreConfig = &StoreConfig{
  BasePath:   path.Join(HomeDir, ".local", "git"),
}

type StoreConfig struct {
  BasePath        string
}

func NewStoreConfig(basepath string) *StoreConfig {
  return &StoreConfig{
    BasePath:     basepath,
  }
}
