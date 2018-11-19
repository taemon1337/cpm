package main

import (
  "os"
  "flag"

  "github.com/taemon1337/cpm/cmd"
  "github.com/taemon1337/cpm/pkg"
  "github.com/taemon1337/cpm/config"
  "github.com/taemon1337/cpm/store"
)

func command(cpm *pkg.ContainerPackageManager, c string, args []string, options *cmd.CommandOptions) pkg.Printable {
  switch c {
    case "list":
      return cpm.List(args, options)
    case "search":
      return cpm.Search(args, options)
    case "build":
      return cpm.Build(args, options)
    case "install":
      return cpm.Install(args, options)
    case "update":
      return cpm.Update(args, options)
    case "remove":
      return cpm.Remove(args, options)
    case "start":
      return cpm.Start(args, options)
    case "stop":
      return cpm.Stop(args, options)
    case "restart":
      return cpm.Restart(args, options)
    case "pause":
      return cpm.Pause(args, options)
    case "registry":
      return cpm.Registry(args, options)
    default:
      return cpm.Help(args, options)
  }
}

func main() {
  cfg := config.NewPackageManagerConfig()
  defaults := cmd.DefaultCommandOptions
  options := cmd.NewCommandOptions()

  verbose := flag.String("verbose", defaults.Verbose, "Set the verbosity level (1-5)")
  flag.Parse()

  cfg.Verbose = *verbose

  cpm := pkg.NewContainerPackageManager(cfg)
  args := flag.Args()

  if len(args) > 0 {
    res := command(cpm, args[0], args[1:], options)
    res.Print()

    os.Exit(0)
  } else {
    os.Exit(1)
  }
  return
}
