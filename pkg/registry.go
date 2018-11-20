package pkg

import (
  "github.com/taemon1337/cpm/cmd"
)

type RegistryList struct {
  Registries        []*Registry
}

type Registry struct {
  Name              string
  Repo              string
  Path              string
}

func (c *ContainerPackageManager) Registry(args []string, opts *cmd.CommandOptions) (*PackageList, error) {
  p1 := &Package{
    Name: "cpm",
    Version: "1.0.0",
    Repo: "github.com/taemon1337/cpm",
    Installed: true,
  }

  p2 := &Package{
    Name: "test",
    Version: "1.0.0",
    Repo: "github.com/taemon1337/test",
    Installed: false,
  }

  packages := make([]*Package, 2)

  packages[0] = p1
  packages[1] = p2

  return &PackageList{
    Packages: packages,
  }, nil
}
