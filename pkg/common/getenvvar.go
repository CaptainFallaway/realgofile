package common

import (
	"os"
	"strconv"
)

func GetEnvVar[T any](key string, fallback T) T {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	switch any(fallback).(type) {
	case string:
		return any(value).(T)
	case int:
		if v, err := strconv.Atoi(value); err == nil {
			return any(v).(T)
		}
	case bool:
		if v, err := strconv.ParseBool(value); err == nil {
			return any(v).(T)
		}
	case float64:
		if v, err := strconv.ParseFloat(value, 64); err == nil {
			return any(v).(T)
		}
	}

	return fallback
}
