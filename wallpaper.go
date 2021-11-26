/*
Author: Alistair Rameau
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/reujab/wallpaper"
	"github.com/spf13/viper"
)

type Configuration struct {
	Wallpaper string
}

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	var config Configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return
	}

	usbDir, _ := os.Getwd()                             //directory of the ubs key where the app and image are
	filepath := filepath.Join(usbDir, config.Wallpaper) //complete filepath of the filename

	if err := wallpaper.SetFromFile(filepath); err != nil {
		fmt.Printf("Unable to set iamge as wallpaper, %v", err)
	}
	if err := wallpaper.SetMode(wallpaper.Stretch); err != nil {
		fmt.Printf("Unable to set wallpaper mode, %v", err)
	}
	return
}
