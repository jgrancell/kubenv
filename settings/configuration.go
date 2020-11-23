package settings

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "os"
)

type Configuration struct {
  KubeDir  string `json:"kube_dir"`
  KubeConf string `json:"kube_conf"`
  EnvDir   string `json:"env_dir"`
}

func LoadConfiguration() (Configuration, error) {
  conf := Configuration{}

  var homedir string
  // Checking for $HOME environment variable
  if homedir = os.Getenv("HOME"); homedir != "" {
    // Ensuring the $HOME directory exists
    if _, err := os.Stat(homedir); err == nil {
      // Attempting to read our configuration file
      if _, err := os.Stat(homedir + "/.kubenv.json"); err == nil {
        file, err := ioutil.ReadFile(homedir + "/.kubenv.json")
        if err == nil {
          _ = json.Unmarshal([]byte(file), &conf)
        } else {
          return Configuration{}, errors.New("Kubenv configuration file at " + homedir + "/.kubenv.json is not valid or is unreadable.")
        }
      }
    } else {
      // We cannot read the homedir. Erroring.
      return Configuration{}, errors.New("$HOME env variable points to " + homedir + " which does not exist or is unreadable.")
    }
  } else {
    // $HOME appears to be unset. Erroring.
    return Configuration{}, errors.New("Unable to determine user's home directory from HOME env variable.")
  }

  if conf.KubeDir == "" {
    conf.KubeDir = homedir + "/.kube"
  }
  kErr := setupDir(conf.KubeDir)
  if kErr != nil {
    return Configuration{}, kErr
  }

  if conf.KubeConf == "" {
    conf.KubeConf = conf.KubeDir + "/config"
  }

  if conf.EnvDir == "" {
    conf.EnvDir = homedir + "/.kube"
  }
  eErr := setupDir(conf.KubeDir)
  if eErr != nil {
    return Configuration{}, eErr
  }

  return conf, nil
}

func setupDir(dir string) error {
  // Checking to make sure your EnvDir exists
  if _, err := os.Stat(dir); err != nil {
    // It doesn't, so checking to determine if we can create it.
    if os.IsNotExist(err) {
      mkErr := os.Mkdir(dir, 0700)
      if mkErr != nil {
        return errors.New("The directory " + dir + " does not exist and cannot be created by Kubenv.")
      }
    } else {
      return errors.New("The directory " + dir + " cannot be properly accessed or created for use.")
    }
  }
  return nil
}
