package pkg

import (
  "path"
  "log"
  "errors"

  "github.com/taemon1337/cpm/cmd"
)

func (c *ContainerPackageManager) Install(args []string, opts *cmd.CommandOptions) (*Package, error) {
  if len(args) != 1 {
    log.Printf("Invalid Parameters: %s", args)
    return nil, errors.New("Usage: cpm install <git-repo-url>")
  }

  var pull = true
  url := args[0]
  branch := opts.Branch
  name := path.Base(url)
  log.Printf("[INSTALL] %s", url)

  repo, err := c.Store.Load(name, url, branch, pull)

  if err != nil {
    log.Printf("Could not clone git repo: %s: %s", url, err)
    return nil, err
  }

  log.Printf("REPO: ", repo)

  p := &Package{
    Name: "cpm",
    Version: "1.0.0",
    Repo: "github.com/taemon1337/cpm",
    Installed: false,
  }

  return p, nil
}
