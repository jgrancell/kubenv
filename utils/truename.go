package utils

import "github.com/jgrancell/kubenv/settings"

func Truename(name string, conf settings.Configuration) string {
  truename := conf.EnvDir + "/" + name + "-kubenv"
  return truename
}
