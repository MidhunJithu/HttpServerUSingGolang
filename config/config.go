package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configurations struct {
	Server    ServerConfigurations
	AppName   string `yaml:"AppName"`
	WebhokUrl string `yaml:"WebhokUrl"`
}

type ServerConfigurations struct {
	Port int
	Host string
}

func Initconfig() *Configurations {
	config := Configurations{}
	if fileName, err := checkAndSetEnvVariable("env", "local"); err != nil {
		log.Default().Printf("[Error] [InitConfig] : %s", err)
		panic("Some issue in setting up the env variable")
	} else {
		file, err := os.Open(fmt.Sprintf("config/%s.yaml", fileName))
		defer file.Close()
		if err != nil {
			log.Default().Printf("[Error] [ReadFile] : %s", err)
			panic("Some issue in reading the config file")
		}
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			log.Default().Printf("[Error] [ParseConfig] : %s", err)
			panic("Some issue in parsing the config file")
		}
		log.Default().Printf("This is environment - %s", fileName)
		log.Default().Printf("This is config - %v", &config)
		log.Default().Printf("Starting the application - %s", config.AppName)
	}

	return &config
}

func checkAndSetEnvVariable(key string, defaultvalue interface{}) (string, error) {
	if val, ok := os.LookupEnv(key); ok && len(val) > 0 {
		return val, nil
	}
	err := os.Setenv(key, fmt.Sprintf("%s", defaultvalue))
	return fmt.Sprintf("%s", defaultvalue), err
}
