package pkg

import (
  "github.com/taemon1337/cpm/cmd"
)

func (c *ContainerPackageManager) Install(args []string, opts *cmd.CommandOptions) *Package {
  p := &Package{
    Name: "cpm",
    Version: "1.0.0",
    Repo: "github.com/taemon1337/cpm",
    Installed: false,
  }

  return p
}
