package cmd

import (
  "fmt"
  "io/ioutil"
  "regexp"
  "strings"

  "github.com/jgrancell/kubenv/settings"
)

func List(conf settings.Configuration) (bool, error) {

  kubedir := conf.EnvDir

  files, err := ioutil.ReadDir(kubedir)
  if err != nil {
    return false, err
  }

  var matches []string
  count := 0

  for _, file := range files {
    if matched, _ := regexp.Match(`.*-kubenv`, []byte(file.Name())); matched {
      count++
      matches = append(matches, file.Name())
    }
  }

  if count > 0 {
    fmt.Println("Managed environments:")
    for _, match := range matches {
      fmt.Println("  -", strings.TrimSuffix(match, "-kubenv"))
    }
  } else {
    fmt.Println("No environments currently managed by Kubenv.")
  }
  return true, nil
}
