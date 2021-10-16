package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

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

			if err := os.Rename(target, truename); err != nil {
				return false, errors.New("unable to import " + target + " to environment directory " + conf.EnvDir)
			}

			if err := os.Chmod(truename, 0600); err != nil {
				return false, errors.New("unable to set kubenv file for " + target + " to 0600")
			}
		}
		return true, nil
	} else {
		// Missing a target. Erroring.
		return false, errors.New("no target file specified for import")
	}
}
