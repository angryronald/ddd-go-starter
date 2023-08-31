package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/vault/api"
)

var secret *api.Secret

func init() {
	vaultClient, err := api.NewClient(&api.Config{
		Address: getEnv(VAULT_ADDRESS),
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		log.Fatalf("Cannot connect to vault: %v", err)
	}

	// only needed if the env for vault token have different name (not VAULT_TOKEN)
	// vaultClient.SetToken(getEnv(VAULT_TOKEN))
	secret, err = vaultClient.Logical().Read(fmt.Sprintf("medicplus/%s/%s", getEnv(SERVICE_NAME), getEnv(ENV)))
	if err != nil || secret == nil {
		log.Fatalf("Cannot collect secret: %v", err)
	}
}

var defaultConfig = map[string]string{
	// Common
	ENV: ENV_DEVELOPMENT,

	// Database
	DATABASE_CONNECTION_STRING: "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	DATABASE_DRIVER:            "postgres",
	DATABASE_HOST:              "localhost",
	DATABASE_PORT:              "5432",
	DATABASE_NAME:              "postgres",
	DATABASE_USER:              "postgres",
	DATABASE_PASS:              "password",
	DATABASE_SSL:               "disable",
	REDIS_DB:                   "0",
	REDIS_ADDRESS:              "127.0.0.1:6379",
	REDIS_PASSWORD:             "",

	// Migration
	MIGRATION_PATH: "database/migrations",

	// Transport
	HTTP_PORT:          ":8001",
	HTTP_PROFILER_PORT: ":9100",

	// Vault
	VAULT_ADDRESS: "http://127.0.0.1:8200/v1/kv/",
	VAULT_TOKEN:   "",
}

func getEnv(KEY string) string {
	r := os.Getenv(KEY)

	if r == "" {
		if val, ok := defaultConfig[KEY]; ok {
			return val
		}

		return ""
	}

	return r
}

func GetValue(key string) string {
	value, ok := secret.Data[key]
	if !ok {
		log.Fatalf("Secret for %s is not found, looking for mock data", key)
	}
	return value.(string)
}

func SetValue(key string, value string) {
	secret.Data[key] = value
}

func GetFlavor() string {
	return getEnv(ENV)
}

func GetWorkingDirectory() (string, error) {
	return os.Getwd()
}
