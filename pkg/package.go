package pkg

import (
  "fmt"
)

type PackageList struct {
  Packages      []*Package
}

func (l *PackageList) Print() {
  fmt.Println("Name\t\tVersion\t\tInstalled\t\tRepository")
  fmt.Println("----\t\t-------\t\t---------\t\t----------")
  for i := range l.Packages {
    l.Packages[i].Print()
  }
}


type Package struct {
  Name          string
  Version       string
  Repo          string
  Installed     bool
}

func (p *Package) Print() {
  i := "false"
  if p.Installed { i = "true" }
  fmt.Printf("%s\t\t%s\t\t%s\t\t%s\n", p.Name, p.Version, i, p.Repo)
}


