package configer

import "fmt"

type ServeConfig struct {
	Host string
	Port string
}

func LoadServeConfig() ServeConfig {
	serveConfig := ServeConfig{}
	configLoader("serve.json", &serveConfig)

	return serveConfig
}

func (config *ServeConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}
