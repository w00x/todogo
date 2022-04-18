package shared

import (
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

func GetEnv(key string) string {
	env := os.Getenv("GIN_MODE")
	if env == "" {
		envPath := dotEnvPath()
		godotenv.Load(envPath)
	}
	return os.Getenv(key)
}

func dotEnvPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	envPath := ""

	for envPath == "" {
		files, err := ioutil.ReadDir(currentPath)
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			if !f.IsDir() {
				if f.Name() == ".env" {
					envPath = currentPath + "/" + f.Name()
				}
			}
		}
		currentPath += "/.."
	}

	return envPath
}
