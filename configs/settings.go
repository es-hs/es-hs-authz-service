package config

import (
	"os"
)

const ()

var (
	Mode     string
	GrpcPort string
	RedisURL string
)

func init() {
	Mode = os.Getenv("MODE")
	GrpcPort = os.Getenv("GRPC_PORT")
	RedisURL = os.Getenv("REDIS_URL")
}

func IsProduction() bool {
	return os.Getenv("MODE") == "production"
}

func IsStaging() bool {
	return os.Getenv("MODE") == "staging"
}

func IsDev() bool {
	return !IsProduction() && !IsStaging()
}
