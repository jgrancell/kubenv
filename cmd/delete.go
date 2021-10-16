package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/jgrancell/kubenv/settings"
	"github.com/jgrancell/kubenv/utils"
)

func Delete(targets []string, conf settings.Configuration) (bool, error) {
	// Making sure we have at least one target
	if len(targets) >= 1 {
		deleteErrors := false
		// Looping through all targets
		for _, target := range targets {
			// Validating the existance of the target Environment
			if _, err := utils.ValidateEnv(target, conf); err == nil {
				r := os.Remove(utils.Truename(target, conf))
				if r != nil {
					deleteErrors = true
					fmt.Println(r)
				}
			} else {
				deleteErrors = true
				fmt.Println(err)
			}
		}
		if deleteErrors {
			return false, errors.New("one or more target environment could not be deleted")
		}
		return true, nil
	} else {
		// Missing a target. Erroring.
		return false, errors.New("you must provide at least one target file to import")
	}
}
