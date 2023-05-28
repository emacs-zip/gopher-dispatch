package config

import "os"

type Config struct {
    DBUsername string
    DBPassword string
    DBName string
    DBHost string
    JWTSecret string
}

// Gets the variable from env, returns "fallback" if env var empty
func GetEnv(key string, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }

    return fallback
}

func Get() Config {
    config := Config{
        DBUsername: "postgres",
        DBPassword: "root@123",
        DBName: "gopher-dispatch",
        DBHost: "localhost",
        JWTSecret: GetEnv("JWT_SECRET", "SuperSecretKey"),
    }

    return config
}
