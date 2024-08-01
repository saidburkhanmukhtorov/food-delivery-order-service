package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config struct holds the configuration settings.
type Config struct {
	HTTPPort string

	// PostgreSQL Configuration (Development)
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	// PostgreSQL Configuration (Testing)
	PostgresHostTest     string
	PostgresPortTest     int
	PostgresUserTest     string
	PostgresPasswordTest string
	PostgresDBTest       string

	// MongoDB Configuration
	MongoHost          string
	MongoPort          int
	MongoUser          string
	MongoPassword      string
	MongoDB            string
	MongoContainerName string // Docker Compose container name

	// Kafka Configuration
	KafkaBrokers     []string
	KafkaBrokersTest []string

	LOG_PATH string
}

// Load loads the configuration from environment variables.
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPPort = cast.ToString(coalesce("HTTP_PORT", ":9090"))

	// PostgreSQL Configuration (Development)
	config.PostgresHost = cast.ToString(coalesce("POSTGRES_HOST", "postgres_dock"))
	config.PostgresPort = cast.ToInt(coalesce("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(coalesce("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(coalesce("POSTGRES_PASSWORD", "root"))
	config.PostgresDB = cast.ToString(coalesce("POSTGRES_DB", "memory"))

	// PostgreSQL Configuration (Testing)
	config.PostgresHostTest = cast.ToString(coalesce("POSTGRES_HOST_TEST", "localhost"))
	config.PostgresPortTest = cast.ToInt(coalesce("POSTGRES_PORT_TEST", 5432))
	config.PostgresUserTest = cast.ToString(coalesce("POSTGRES_USER_TEST", "sayyidmuhammad"))
	config.PostgresPasswordTest = cast.ToString(coalesce("POSTGRES_PASSWORD_TEST", "root"))
	config.PostgresDBTest = cast.ToString(coalesce("POSTGRES_DB_TEST", "postgres"))

	// MongoDB Configuration
	config.MongoHost = cast.ToString(coalesce("MONGO_HOST", "mongo"))
	config.MongoPort = cast.ToInt(coalesce("MONGO_PORT", 27017))
	config.MongoUser = cast.ToString(coalesce("MONGO_USER", "root"))
	config.MongoPassword = cast.ToString(coalesce("MONGO_PASSWORD", "example"))
	config.MongoDB = cast.ToString(coalesce("MONGO_DB", "your_mongodb_name"))

	// Kafka
	config.KafkaBrokers = cast.ToStringSlice(coalesce("KAFKA_BROKERS", []string{"kafka:9092"}))
	config.KafkaBrokersTest = cast.ToStringSlice(coalesce("KAFKA_BROKERS_Test", []string{"localhost:9092"}))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
