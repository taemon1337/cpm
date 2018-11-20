package pkg

import (
  "github.com/taemon1337/cpm/cmd"
)

func (c *ContainerPackageManager) Restart(args []string, opts *cmd.CommandOptions) (*Package, error) {
  _, err := c.Stop(args, opts)

  if err != nil {
    return nil, err
  }

  return c.Start(args, opts)
}
