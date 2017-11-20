package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	// the server port. Defaults to 8080
	ServerPort         int    `mapstructure:"server_port"`
	// the data source name for connecting to the database. required.
	DataSource                string `mapstructure:"data_source"`

	// image storage path
	ImageStorePath          string `mapstructure:"image_store_path"`
	// dribble access keys
	ClientID		   string `mapstructure:"client_id"`
	ClientSecret	   string `mapstructure:"client_secret"`
	ClientAccessToken  string `mapstructure:"client_access_token"`
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetDefault("server_port", 8080)
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return nil
}



