package pkg

import (
  "fmt"
  "github.com/taemon1337/cpm/cmd"
)

type HelpPrintable struct {}

func (h *HelpPrintable) Print() {
  fmt.Println("Help Menu: ")
}

func (c *ContainerPackageManager) Help(args []string, opts *cmd.CommandOptions) *HelpPrintable {
  return &HelpPrintable{}
}
