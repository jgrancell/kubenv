package cmd

import (
	"errors"
	"os"

	"github.com/jgrancell/kubenv/settings"
	"github.com/jgrancell/kubenv/utils"
)

func Activate(targets []string, conf settings.Configuration) (bool, error) {
	// Making sure we have a target
	if len(targets) == 1 {
		target := targets[0]
		r := os.Remove(conf.KubeConf)
		if r != nil && !os.IsNotExist(r) {
			return false, errors.New("the current KUBECONFIG file at " + conf.KubeConf + " cannot be moved")
		}
		linkerr := os.Link(utils.Truename(target, conf), conf.KubeConf)
		if linkerr != nil {
			return false, linkerr
		}
		return true, nil
	}
	return false, errors.New("you must specify exactly one environment to activate")
}
