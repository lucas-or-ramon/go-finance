package config

import (
	"fmt"
	"os"
	"strconv"
)

func getStringOrDefault(variable, fallback string) string {
	value, ok := os.LookupEnv(variable)
	if ok {
		return value
	}
	return fallback
}

func getIntOrDefault(variable string, fallback int) int {
	value, ok := os.LookupEnv(variable)
	if ok {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return intValue
	}
	return fallback
}

func getDatabaseConfiguration() Database {
	url := getStringOrDefault("postgres_url", "")
	databaseName := getStringOrDefault("database", "")

	return Database{
		URL:          url,
		DatabaseName: databaseName,
	}
}

func getWebConfiguration() Web {
	port := getIntOrDefault("port", 8080)
	address := fmt.Sprintf("0:0:0:0:%d", port)
	return Web{
		Address: address,
	}
}

func getConfiguration() Configuration {
	return Configuration{
		Web:      getWebConfiguration(),
		Database: getDatabaseConfiguration(),
	}
}
