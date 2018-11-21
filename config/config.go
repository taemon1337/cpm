package config

import (
  "os"
  "log"
  "path"
  "fmt"
  "io/ioutil"

  "gopkg.in/yaml.v2"
  "github.com/taemon1337/cpm/store"
)

type PackageManagerConfig struct {
  Version               string                    `yaml:"version"`
  Verbose               string                    `yaml:"verbose"`
  StoreConfig           *store.StoreConfig        `yaml:"store"`
}

func (c *PackageManagerConfig) Print() {
  d, err := yaml.Marshal(&c)
  if err != nil {
    log.Printf("Cannot marshal config: %s", err)
  }

  fmt.Printf(string(d))
}


func (c *PackageManagerConfig) Save(filepath string) (int, error) {
  if _, err := os.Stat(filepath); os.IsNotExist(err) {
    os.MkdirAll(path.Dir(filepath), os.ModePerm)
    _, err := os.Create(filepath)

    if err != nil {
      log.Printf("WARN: Cannot create file: %s: %s", filepath, err)
    }
  }

  d, err := yaml.Marshal(c)

  if err != nil {
    return 0, err
  }

  err = ioutil.WriteFile(filepath, d, 0644)

  if err != nil {
    return 0, err
  }

  return len(d), nil
}

func (c *PackageManagerConfig) Reset(filepath string) (int, error) {
  if _, err := os.Stat(filepath); !os.IsNotExist(err) {
    err := os.Remove(filepath)
    if err != nil {
      return 0, err
    }
  }

  cfg := NewPackageManagerConfig()
  return cfg.Save(filepath)
}

func LoadPackageManagerConfig(filepath string) (*PackageManagerConfig, error) {
  cfg := NewPackageManagerConfig()

  if _, err := os.Stat(filepath); os.IsNotExist(err) {
    log.Printf("[WARN] No package manager config found: %s, use 'cpm config init' to create it.", filepath)
    return cfg, nil
  }

  dat, err := ioutil.ReadFile(filepath)

  if err != nil {
    return nil, err
  }

  err = yaml.Unmarshal(dat, &cfg)

  if err != nil {
    return nil, err
  }

  return cfg, nil
}

func NewPackageManagerConfig() *PackageManagerConfig {
  return &PackageManagerConfig{
    Version:            DefaultPackageManagerConfig.Version,
    Verbose:            DefaultPackageManagerConfig.Verbose,
    StoreConfig:        store.DefaultStoreConfig,
  }
}
