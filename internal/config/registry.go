package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type IConfigRegistry interface {
	RegisterConfig(name string, config IConfig)
	GetConfig(name string) IConfig
	SetDefaultConfigs(v *viper.Viper)
}

type registryImpl struct {
	mu      sync.RWMutex
	configs map[string]IConfig
}

func NewConfigRegistry() IConfigRegistry {
	return &registryImpl{
		configs: make(map[string]IConfig),
	}
}

func (r *registryImpl) RegisterConfig(name string, config IConfig) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.configs[name]
	if ok {
		panic(fmt.Sprintf("a config with name %s already exist", name))
	}

	r.configs[name] = config
}

func (r *registryImpl) GetConfig(name string) IConfig {
	r.mu.RLock()
	defer r.mu.RUnlock()

	config, ok := r.configs[name]
	if !ok {
		return nil
	}

	return config
}

func (r *registryImpl) SetDefaultConfigs(v *viper.Viper) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, config := range r.configs {
		config.SetDefault(v)
	}
}
