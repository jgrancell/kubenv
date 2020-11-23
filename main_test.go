package main

import (
  "os"
  "os/exec"
  "testing"
)

func TestNoArguments(t *testing.T) {
  if os.Getenv("NO_ARGUMENTS") == "1" {
    // Running with an empty slice as arguments
    Run(make([]string, 0))
    return
  }
  cmd := exec.Command(os.Args[0], "-test.run=TestNoArguments")
  cmd.Env = append(os.Environ(), "NO_ARGUMENTS=1")
  err := cmd.Run()
  if e, ok := err.(*exec.ExitError); ok && !e.Success() {
    return
  }
  t.Fatalf("process ran with err %v, want exit status 1", err)
}
