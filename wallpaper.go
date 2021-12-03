/*
Author: Alistair Rameau
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/reujab/wallpaper"
	"github.com/spf13/viper"
)

type Configuration struct {
	Wallpaper string
	ImagesDir string
	FillMode  string
	Random    bool
}

var IMAGE_EXTENSIONS []string = []string{"JPG", "BMP", "JPEG", "PNG"}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	rand.Seed(time.Now().UnixNano())

	var config Configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return
	}
	currentDir, err := os.Getwd() //directory of the ubs key where the app and image are
	if err != nil {
		fmt.Printf("Unable to get current directory, %v", err)
		return
	}

	if config.Random {
		imagesPath := filepath.Join(currentDir, config.ImagesDir) //complete filepath of the filename
		imageList, err := GetImageListFromDirectory(imagesPath)
		if err != nil {
			fmt.Printf("Unable to retrieve images directory's files, %v", err)
			return
		}
		randomImage := GetRandomEntryFromList(imageList)
		err = SetImageAsWallpaper(randomImage, ModeStringToConst(config.FillMode))
		if err != nil {
			fmt.Printf("Unable to set image as wallpaper, %v", err)
			return
		}
	} else {
		imagePath := filepath.Join(currentDir, config.Wallpaper) //complete filepath of the filename
		err := SetImageAsWallpaper(imagePath, ModeStringToConst(config.FillMode))
		if err != nil {
			fmt.Printf("Unable to set image as wallpaper, %v", err)
			return
		}
	}
	return
}

func GetImageListFromDirectory(directoryPath string) ([]string, error) {
	directoryPath = strings.TrimSpace(directoryPath)

	//open the directory and read all its files
	directory, err := os.Open(directoryPath)
	if err != nil {
		return nil, err
	}
	dirFiles, err := directory.Readdir(0) //read all files in the directory, 0 = no limit
	if err != nil {
		return nil, err
	}
	var fileList []string = make([]string, 0)
	//looping over the directory's files
	for index := range dirFiles {
		file := dirFiles[index]
		if isAnImage(file.Name()) {
			fileList = append(fileList, filepath.Join(directoryPath, file.Name()))
		}
	}
	return fileList, nil
}

func GetRandomEntryFromList(list []string) string {
	return list[rand.Intn(len(list))]
}

func SetImageAsWallpaper(imagePath string, mode wallpaper.Mode) error {
	if err := wallpaper.SetFromFile(imagePath); err != nil {
		return err
	}
	if err := wallpaper.SetMode(mode); err != nil {
		return err
	}
	return nil
}

func ModeStringToConst(mode string) wallpaper.Mode {
	switch strings.ToUpper(mode) {
	case "STRETCH":
		return wallpaper.Stretch
	case "FIT":
		return wallpaper.Fit
	case "CENTER":
		return wallpaper.Center
	case "CROP":
		return wallpaper.Crop
	case "SPAN":
		return wallpaper.Span
	case "TILE":
		return wallpaper.Tile
	default:
		return wallpaper.Stretch
	}
}

func isAnImage(name string) bool {
	splits := strings.Split(name, ".")
	if len(splits) > 0 {
		return ListContains(IMAGE_EXTENSIONS, splits[len(splits)-1])
	} else {
		return false
	}
}

func ListContains(list []string, value string) bool {
	if len(list) < 1 {
		return false
	}
	for i := 0; i < len(list); i++ {
		if strings.EqualFold(list[i], value) {
			return true
		}
	}
	return false
}
