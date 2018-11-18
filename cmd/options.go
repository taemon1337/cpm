package cmd

import (

)

var DefaultCommandOptions = &CommandOptions{
  Verbose:        "1",
}

type CommandOptions struct {
  Verbose         string
}

func NewCommandOptions() *CommandOptions {
  return &CommandOptions{}
}
