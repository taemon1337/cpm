package pkg

import (
  "fmt"
  "errors"

  "github.com/taemon1337/cpm/cmd"
  "github.com/taemon1337/cpm/config"
)

func (c *ContainerPackageManager) Config(args []string, opts *cmd.CommandOptions) (cfg *config.PackageManagerConfig, err error) {
  if len(args) < 1 {
    return nil, errors.New(fmt.Sprintf("Invalid Arguments: %s", args))
  }

  cfg, err = config.LoadPackageManagerConfig(opts.ConfigFile)
  if err != nil {
    return cfg, err
  }

  switch args[0] {
    case "init":
      _, err := cfg.Save(opts.ConfigFile)
      if err != nil {
        return cfg, err
      }
    case "reset":
      _, err := cfg.Reset(opts.ConfigFile)
      if err != nil {
        return cfg, err
      }
      fmt.Printf("[RESET] Successfully deleted and re-initialized %s\n", opts.ConfigFile)
  }
  return cfg, nil
}
