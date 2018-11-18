package pkg

import (
  "github.com/taemon1337/cpm/cmd"
)

func (c *ContainerPackageManager) List(args []string, opts *cmd.CommandOptions) *PackageList {
  p := &Package{
    Name: "cpm",
    Version: "1.0.0",
    Repo: "github.com/taemon1337/cpm",
    Installed: false,
  }

  packages := make([]*Package, 1)

  packages[0] = p

  return &PackageList{
    Packages: packages,
  }
}
