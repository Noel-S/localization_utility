package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"gopkg.in/yaml.v3"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	OriginFileExtension string   `yaml:"origin_file_extension"`
	OriginLocale        string   `yaml:"origin_locale"`
	OutputLocales       []string `yaml:"output_locales"`
	OutputFolder        string   `yaml:"output_folder"`
	InputFolder         string   `yaml:"input_folder"`
}

func main() {
	fmt.Println("Localization util v.0.0.1")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	data_config, err := os.ReadFile("config.yaml")
	checkErr(err)
	config := Config{}
	// parse config
	err = yaml.Unmarshal(data_config, &config)
	checkErr(err)

	// check if config is correct
	if config.OriginFileExtension == "" {
		log.Fatal("origin_file_extension is required")
	}
	if config.OriginLocale == "" {
		log.Fatal("origin_locale is required")
	}
	if len(config.OutputLocales) == 0 {
		log.Fatal("output_locales is required")
	}
	if config.OutputFolder == "" {
		log.Fatal("output_folder is required")
	}
	if config.InputFolder == "" {
		log.Fatal("input_folder is required")
	}

	fmt.Println("Origin file extension: " + config.OriginFileExtension)
	fmt.Println("Origin locale: " + config.OriginLocale)
	fmt.Println("Output locales: " + strings.Join(config.OutputLocales, ", "))
	fmt.Println("Output folder: " + config.OutputFolder)

	filepath.Walk(config.InputFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() {
			if info.Name() == config.OriginLocale+"."+config.OriginFileExtension {
				fmt.Println(path)
				data, err := os.ReadFile(path)
				checkErr(err)

				// copy original file to the new folder
				pathComponents := strings.Split(path, "/")
				destPath := config.OutputFolder + "/" + strings.Join(pathComponents[1:], "/")
				// create the directory structure in the new folder
				err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
				checkErr(err)
				// copy the file to the new folder
				err = os.WriteFile(destPath, data, 0644)
				checkErr(err)

				for _, locale := range config.OutputLocales {
					// copy the file with the locale as the name
					destPath := config.OutputFolder + "/" + strings.Join(pathComponents[1:len(pathComponents)-1], "/") + "/" + locale + "." + config.OriginFileExtension
					// create the directory structure in the new folder
					err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
					checkErr(err)
					// copy the file to the new folder
					err = os.WriteFile(destPath, data, 0644)
					checkErr(err)
				}
			}
		}
		return nil
	})
	s.Stop()
}
