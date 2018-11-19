package store

import (

)

type StoreConfig {
  Type            string
}

func NewStoreConfig(storetype string) *StoreConfig {
  return &StoreConfig{
    Type:     storetype,
  }
}
