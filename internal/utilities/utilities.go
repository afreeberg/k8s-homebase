package utilities

import (
	"github.com/afreeberg/k8s-homebase/internal/models"
	"github.com/spf13/viper"
)

func GetConfiguration() (config models.Configuration, err error) {

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}
