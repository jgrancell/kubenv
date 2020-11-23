package cmd

import (
  "testing"
)

func TestLoadConfiguration(t *testing.T) {
  result, err := Help("1.0.0")

  if err != nil {
    t.Errorf(err.Error())
  }
  if ! result {
    t.Errorf("Help function should never return false.")
  }
}
