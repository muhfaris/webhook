package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var App = struct {
	Port     int      `json:"port" mapstructure:"port"`
	Webhooks Webhooks `json:"webhooks" mapstructure:"webhooks"`
}{}

type WebhookConfig struct {
	ID               string            `json:"id" mapstructure:"id"`
	Token            string            `json:"token" mapstructure:"token"`
	Workdir          string            `json:"workdir" mapstructure:"workdir"`
	ExecuteCommand   string            `json:"execute_command" mapstructure:"execute_command"`
	CommandArguments []CommandArgument `json:"command_arguments" mapstructure:"command_arguments"`
}

type CommandArgument struct {
	Source string `json:"source" mapstructure:"source"`
	Name   string `json:"name" mapstructure:"name"`
}

type Webhooks []WebhookConfig

func (w Webhooks) ByID(webhookID string) (WebhookConfig, error) {
	for _, webhook := range w {
		if webhook.ID == webhookID {
			return webhook, nil
		}
	}
	return WebhookConfig{}, fmt.Errorf("webhook not found")
}

func ReadConfig(path string) error {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.AddConfigPath(path)
		viper.SetConfigName("config")
		viper.SetConfigType("json")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error read config: %v", err)
	}

	if err := viper.Unmarshal(&App); err != nil {
		return fmt.Errorf("error parse the config: %v", err)
	}

	if App.Port == 0 {
		App.Port = 8081
	}

	return nil
}
