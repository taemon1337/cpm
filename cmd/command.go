package cmd

import (
)

type CommandInterface interface {
  execute([]string, *CommandOptions) int
}
