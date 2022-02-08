package settings

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Configuration *viper.Viper

func init() {
	environment := os.Getenv("ENVIROMENT")
	if environment == "" {
		environment = "development"
		var builder strings.Builder
		builder.WriteString("\033[33m")
		builder.WriteString("\nenvironment variable is not set, setting environment to: ")
		builder.WriteString(environment)
		builder.WriteString("\033[0m")
	}

	Configuration = viper.New()
	Configuration.SetConfigType("yaml")
	Configuration.SetConfigName(environment)
	Configuration.AddConfigPath("../settings")
	Configuration.AddConfigPath("settings/")
	err := Configuration.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
