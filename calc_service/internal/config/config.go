package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port       string
	AuthSecret string
	RateLimit  float64
	RateBurst  int
	DBURL      string
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	authSecret := os.Getenv("AUTH_SECRET")
	if authSecret == "" {
		return nil, errors.New("AUTH_SECRET is required")
	}

	rateLimitStr := os.Getenv("RATE_LIMIT")
	if rateLimitStr == "" {
		rateLimitStr = "5"
	}
	rateLimit, err := strconv.ParseFloat(rateLimitStr, 64)
	if err != nil {
		return nil, errors.New("invalid RATE_LIMIT")
	}

	rateBurstStr := os.Getenv("RATE_BURST")
	if rateBurstStr == "" {
		rateBurstStr = "10"
	}
	rateBurst, err := strconv.Atoi(rateBurstStr)
	if err != nil {
		return nil, errors.New("invalid RATE_BURST")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, errors.New("DB_URL is required")
	}

	return &Config{
		Port:       port,
		AuthSecret: authSecret,
		RateLimit:  rateLimit,
		RateBurst:  rateBurst,
		DBURL:      dbURL,
	}, nil

}
