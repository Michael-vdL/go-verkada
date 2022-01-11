package client

import (
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestConfigReader(t *testing.T) {

  // Load config
  var config SDKConfig
  configFile, err := os.Open("config.yaml")
  if err != nil {
    t.Fatal(err)
  }
  defer configFile.Close()

  err = yaml.NewDecoder(configFile).Decode(&config)
  if err != nil {
    t.Fatal(err)
  }

  t.Logf("%+v", config.VerkadaCfg)
}

func TestClientFromtConfiguration(t *testing.T) {
  // Load config
  var config SDKConfig
  configFile, err := os.Open("config.yaml")
  if err != nil {
    t.Fatal(err)
  }
  defer configFile.Close()

  err = yaml.NewDecoder(configFile).Decode(&config)
  if err != nil {
    t.Fatal(err)
  }

	// Prepare
  ac, err := NewRestClientFromConfigFile()
  if err != nil {
    t.Fatal(err)
  }
  // Test client

  // Expect last URL path to match organization ID
  clientUrl := strings.Split(ac.baseURL, "/")
  t.Logf("Expected %s, got %s", config.VerkadaCfg.OrganizationID, clientUrl[len(clientUrl)-1])

}