
package config

import (
    "fmt"
	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
    Satellites   SatellitesrConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port string
}

type SatellitesrConfigurations struct {
	Satellite_1_x float64
    Satellite_1_y float64
    Satellite_2_x float64
    Satellite_2_y float64
    Satellite_3_x float64
    Satellite_3_y float64
}

func GetConfiguration(path string)(Configurations){

    viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

    return configuration

}