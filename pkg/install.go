package pkg

import (
  "path"
  "log"
  "github.com/taemon1337/cpm/cmd"
)

func (c *ContainerPackageManager) Install(args []string, opts *cmd.CommandOptions) (*Package, error) {
  if len(args) > 0 {
    var pull = true
    url := args[0]
    name := path.Base(url)
    log.Printf("[INSTALL] %s", url)

    repo, err := c.Store.Load(name, url, opts.Branch, pull)

    if err != nil {
      log.Printf("Could not clone git repo: %s: %s", url, err)
      return nil, err
    }

    log.Printf("REPO: ", repo)
  }

  p := &Package{
    Name: "cpm",
    Version: "1.0.0",
    Repo: "github.com/taemon1337/cpm",
    Installed: false,
  }

  return p, nil
}
