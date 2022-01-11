package client

import (
	"fmt"
	"os"

	"github.com/dghubble/sling"
	"gopkg.in/yaml.v2"
)

const (
	MajorVersion    = 0
	MinorVersion    = 1
	PatchVersion    = 0
	UserAgentPrefix = "verkada-api-go-client"
)

type NotFound string

func (nf NotFound) Error() string {
	return string(nf)
}

var Version = fmt.Sprintf("%d.%d.%d", MajorVersion, MinorVersion, PatchVersion)

type SDKConfig struct {
	VerkadaCfg VerkadaConfig `yaml:"verkada"`
	ElasticCfg ElasticConfig `yaml:"elastic"`
}

type VerkadaConfig struct {
	OrganizationID string `yaml:"organization_id"`
	ApiKey         string `yaml:"api_key"`
}

type ElasticConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type APIClient struct {
	baseURL string
	token   string
}

// Move to configurable flags or env vars
const (
	DefaultBaseURL = "https://api.verkada.com"
	OrganizationID = "Enter a valid organization ID (config.yaml)"
	ApiKey         = "Please Enter a valid API Key (config.yaml)"
)


type Client interface {
	AuditLogs() AuditLogClient
}

type OptFunc func(c *APIClient) error


func WithBaseURL(baseURL string) OptFunc {
	return func(c *APIClient) error {
		c.baseURL = baseURL
		return nil
	}
}

func NewRestClient(token string, opts ...OptFunc) (*APIClient, error) {
	c := &APIClient{
		baseURL: DefaultBaseURL,
		token:   token,
	}

	for _, f := range opts {
		if err := f(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func ConfigurationFromFile() (SDKConfig, error) {
	var config SDKConfig
	configFile, err := os.Open("config.yaml")
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	err = yaml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func NewRestClientFromConfigFile() (*APIClient, error) {
	config, err := ConfigurationFromFile()
	if err != nil {
		return nil, err
	}
	
	orgUrl := fmt.Sprintf("%s/orgs/%s/", DefaultBaseURL, config.VerkadaCfg.OrganizationID)
	c := &APIClient{
		baseURL: orgUrl,
		token:   config.VerkadaCfg.ApiKey,
	}
	
	return c, nil
}

func (c *APIClient) client() *sling.Sling {
	return sling.New().Base(c.baseURL).
		Set("User-Agent", fmt.Sprintf("%s (%s)", UserAgentPrefix, Version)).
		Set("x-api-key", c.token)
}

func (c *APIClient) AuditLog() AuditLogClient {
  return &RESTAuditLogClient{client: c}
}