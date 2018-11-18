package pkg

import (
  "github.com/taemon1337/cpm/cmd"
)

func (c *ContainerPackageManager) Restart(args []string, opts *cmd.CommandOptions) *Package {
  c.Stop(args, opts)
  return c.Start(args, opts)
}
