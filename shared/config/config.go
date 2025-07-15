package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// explain why we have the "shared" folder, why we have a config here and another config in seperate projects in the lecture?
type DbConfig struct {
	MongoDuration  time.Duration
	MongoClientURI string
}

var cfgs = map[string]DbConfig{
	"prod": {
		MongoDuration:  time.Second * 10,
		MongoClientURI: "mongodb://root:root1234@mongodb_docker:27017",
	},
	"qa": {
		MongoDuration:  time.Second * 10,
		MongoClientURI: "mongodb://root:root1234@mongodb_docker:27017",
	},
	"dev": {
		MongoDuration:  time.Second * 10,
		MongoClientURI: EnvLoad(),
	},
}

func GetDBConfig(env string) *DbConfig {
	config, isExist := cfgs[env]
	if !isExist {
		panic("config does not exist")
	}
	return &config
}

func EnvLoad() string {
	if err := godotenv.Load("./media/.env"); err != nil {
		panic("environment variable did not load")
	}

	return os.Getenv("MONGO_URI")
}
