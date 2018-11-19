package pkg

import (
  "path"
  "log"
  "github.com/taemon1337/cpm/cmd"
)

func (c *ContainerPackageManager) Install(args []string, opts *cmd.CommandOptions) *Package {
  if len(args) > 0 {
    url := args[0]
    name := path.Base(url)
    log.Printf("[INSTALL] %s", url)

    repo, err := c.Store.Fetch(name, url)

    if err != nil {
      log.Fatal("[ERROR] %s", err)
    }

    log.Printf("[REPO] %s", repo)
  }

  p := &Package{
    Name: "cpm",
    Version: "1.0.0",
    Repo: "github.com/taemon1337/cpm",
    Installed: false,
  }

  return p

}
