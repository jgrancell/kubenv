package cmd

import (
  "fmt"
)

func Help(version string) (bool, error) {
  fmt.Println(`Kubenv is a Kubernetes configuration environment manager.
Usage:
    kubenv [command] [target]

Commands:
    kubenv help

    kubenv use    <target>        Updates the environment to use your chosen target.
    kubenv list                   Lists all K8S environments registered with Kubenv.

    kubenv import <target(s)>     Imports one or more K8S environment files into Kubenv.
    kubenv delete <target(s)>     Deletes one or more K8S environment files from Kubenv.


Kubenv v`+version+` | https://github.com/jgrancell/kubenv`)
  return true, nil
}
