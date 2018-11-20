package store

import (
  "os"
  "log"
  "fmt"
  "path"
  "errors"

  "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/plumbing"
)

type GitStore struct {
  BasePath              string
}

func (g *GitStore) Load(name, url, branch string, pull bool) (repo *git.Repository, err error) {
  var cloned = false
  gitpath := path.Join(g.BasePath, name)

  // if repo does not exist, then clone it
  if !g.Exists(name) {
    repo, err = g.Clone(name, url)

    if err != nil {
      log.Printf("Could not clone git repo: %s: %s", gitpath, err)
      return nil, err
    }

    cloned = true
  }

  // if repo did exist, open it
  if repo == nil {
    repo, err = git.PlainOpen(gitpath)

    if err != nil {
      log.Printf("Could not open git repo: %s: %s", gitpath, err)
      return nil, err
    }
  }

  // open the worktree
  w, err := repo.Worktree()

  if err != nil {
    log.Printf("Could not open git tree: %s: %s", gitpath, err)
    return nil, err
  }

  // checkout the branch
  err = w.Checkout(&git.CheckoutOptions{
    Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
  })

  if err != nil {
    log.Printf("Cannot checkout branch %s of %s: %s", branch, gitpath, err)
    return repo, err
  }

  // if we didn't just clone this and pull is set, then pull latest of this branch
  if !cloned && pull {
    err = w.Pull(&git.PullOptions{RemoteName: "origin"})

    if err != nil {
      log.Printf("Could not pull git repo: %s: %s", gitpath, err)
      return repo, err
    }
  }

  return repo, nil
}

func (g *GitStore) Exists(name string) bool {
  if _, err := os.Stat(path.Join(g.BasePath, name)); os.IsNotExist(err) {
    return false
  } else {
    return true
  }
}

func (g *GitStore) Clone(name, url string) (*git.Repository, error) {
  gitpath := path.Join(g.BasePath, name)

  repo, err := git.PlainClone(gitpath, false, &git.CloneOptions{
    URL: url,
    RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
  })

  if err != nil {
    return nil, err
  }

  return repo, nil
}

func (g *GitStore) Remove(name string) error {
  if name == "" {
    return errors.New("Invalid name of repo to remove")
  }

  dir := path.Join(g.BasePath, name)
  d, err := os.Open(dir)
  if err != nil {
      return err
  }
  defer d.Close()
  names, err := d.Readdirnames(-1)
  if err != nil {
      return err
  }
  for _, name := range names {
    err = os.RemoveAll(path.Join(dir, name))
    if err != nil {
      return err
    }
  }
  return nil
}

func NewGitStore(cfg *StoreConfig) *GitStore {
  return &GitStore{
    BasePath:   cfg.BasePath,
  }
}
