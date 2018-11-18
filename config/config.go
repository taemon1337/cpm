package config

type PackageManagerConfig struct {
  Version               string
  Verbose               string
}

func NewPackageManagerConfig() *PackageManagerConfig {
  return &PackageManagerConfig{
    Version:            DefaultPackageManagerConfig.Version,
    Verbose:            DefaultPackageManagerConfig.Verbose,
  }
}
