package utils

import (
  "errors"
  "os"

  "github.com/jgrancell/kubenv/settings"
)

func ValidateEnv(target string, conf settings.Configuration) (bool, error) {
  if _, err := os.Stat(Truename(target, conf)); err != nil {
    if os.IsNotExist(err) {
      return false, errors.New(target + " does not exist in Kubenv.")
    } else {
      return false, err
    }
  }
  return true, nil
}
