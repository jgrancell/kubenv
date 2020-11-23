package cmd

import (
  "errors"
  "fmt"
  "path/filepath"
  "os"

  "github.com/jgrancell/kubenv/settings"
  "github.com/jgrancell/kubenv/utils"
)

func Import(targets []string, conf settings.Configuration) (bool, error) {
  // Making sure we have at least one target
  if len(targets) >= 1 {

    for _, target := range targets {
      fmt.Println("Importing K8S environment", target)
      basename := filepath.Base(target)
      truename := utils.Truename(basename, conf)
      err := os.Rename(target, truename)

      if err != nil {
        return false, errors.New("Unable to import " + target + " to environment directory " + conf.EnvDir)
      }
    }
    return true, nil
  } else {
    // Missing a target. Erroring.
    return false, errors.New("You must provide at least one target file to import.")
  }
}
