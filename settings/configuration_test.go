package settings

import (
  "testing"
)

func TestLoadConfiguration(t *testing.T) {
  defaultConf, err := LoadConfiguration()
  if err != nil {
    t.Errorf(err.Error())
  }

  if defaultConf.KubeDir != defaultConf.EnvDir {
      t.Errorf("Expected KubeDir and EnvDir to be equal, got %s %s", defaultConf.KubeDir, defaultConf.EnvDir)
  }

  // Testing with custom parameters
  // os.Setenv("KUBENV_CONFIG", "testsuite/.kubenv.json")
  // conf, err := LoadConfiguration()
  // if err != nil {
  //   t.Errorf(err.Error())
  // }

  if defaultConf.KubeDir == "" {
    t.Errorf("KubeDir should not be nil. Got %s", defaultConf.KubeDir)
  }

  if defaultConf.EnvDir == "" {
    t.Errorf("EnvDir should not be nil. Got %s", defaultConf.EnvDir)
  }

  if defaultConf.KubeConf == "" {
    t.Errorf("KubeConf should not be nil. Got %s", defaultConf.KubeConf)
  }
}
