package env

import (
	"os"
	"strconv"
	"time"
)

func GetString(key,fallback string) string{
	if value,ok := os.LookupEnv(key); ok{
		return value
	}
	return fallback
}

func GetInt(key string,fallback int) int{
	if value,ok := os.LookupEnv(key); ok{
		parsedValue,err := strconv.Atoi(value)
		if err != nil{
			return fallback
		}
		return parsedValue
	}
	return fallback
}

func GetDuration(key string,fallback time.Duration) time.Duration{
	if value,ok := os.LookupEnv(key); ok{
		parsedValue,err := time.ParseDuration(value)
		if err != nil{
			return fallback
		}
		return parsedValue
	}
	return fallback
}
