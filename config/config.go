package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Read(env string, config interface{}) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if env != "" {
		f, err := os.Open("config." + env + ".yml")
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		defer f.Close()
		viper.MergeConfig(f)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
