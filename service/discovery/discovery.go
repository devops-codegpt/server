package discovery

import (
	"fmt"
	"github.com/devops-codegpt/server/container"
	"github.com/hashicorp/consul/api"
)

type Discover interface {
	GetServiceByName(name string) (ls *LLMService, err error)
}

type discover struct {
	container container.Container
}

type LLMService struct {
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
}

func (d discover) GetServiceByName(name string) (*LLMService, error) {
	conf := d.container.GetConfig()
	// Create Consul client
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", conf.Consul.Address, conf.Consul.Port)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	// Use the consul client to query service instances
	instances, _, err := client.Health().Service(name, "", true, nil)
	if err != nil {
		return nil, err
	}
	if len(instances) == 0 {
		return nil, fmt.Errorf("no instances available for service: %s", name)
	}
	service := &LLMService{
		Address: instances[0].Service.Address,
		Port:    instances[0].Service.Port,
	}
	return service, nil
}

func New(c container.Container) Discover {
	return &discover{container: c}
}
