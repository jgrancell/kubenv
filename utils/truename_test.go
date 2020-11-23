package utils

import (
  "testing"

  "github.com/jgrancell/kubenv/settings"
)

func TestTruename(t *testing.T) {
  conf := settings.Configuration {
    EnvDir: "testdir/envs",
    KubeDir: "testdir/kube",
    KubeConf: "testdir/kube/conf",
  }

  got := Truename("example", conf)
  if got != "testdir/envs/example-kubenv" {
    t.Errorf("Got: %s; Wanted: testdir/envs/example-kubenv", got)
  }
}
