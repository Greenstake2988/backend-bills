package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	// Cargar la configuración desde un archivo "config.json" en el directorio actual
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error leyendo archivo de configuración:", err)
		return
	}

}
