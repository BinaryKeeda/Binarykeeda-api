package config
import "os"

type Config struct {
	AppPort string
	MongoURI string
	MongoDB  string
}

func Load() Config {
	return Config{
		AppPort: getEnv("APP_PORT", "8080"),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB: getEnv("MONGO_DB", "app_db"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
