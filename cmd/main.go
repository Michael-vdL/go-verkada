package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Michael-vdL/verkada-slingshot/client"
	"gopkg.in/yaml.v2"
)


func main() {
	var config client.SDKConfig

	// Load config
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	err = yaml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}


  // Create client
  orgUrl := fmt.Sprintf("%s/orgs/%s", client.DefaultBaseURL, config.VerkadaCfg.OrganizationID)
  _, err = client.NewRestClient(config.VerkadaCfg.ApiKey, client.WithBaseURL(orgUrl))
  if err != nil {
    log.Fatal(err)
  }

	

}
