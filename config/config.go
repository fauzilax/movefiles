package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env") // specify that it's an .env file
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v", err)
	}

	// Read config values
	port := viper.GetString("PORTSERVER")
	user := viper.GetString("USERNAMESERVER")
	password := viper.GetString("PASSWORDSERVER")
	srcDir := viper.GetString("SOURCEDIRECTORY")

	fmt.Println("PORTSERVER:", port)
	fmt.Println("USERNAMESERVER:", user)
	fmt.Println("PASSWORDSERVER:", password)
	fmt.Println("SOURCEDIRECTORY:", srcDir)
}
