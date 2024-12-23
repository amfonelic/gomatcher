package env

import (
	"os"
	"strconv"
	"strings"
)

type envType interface {
	~string | ~int | bool
}

func GetEnv[T envType](key string) T {
	var val string = os.Getenv(key)
	var zeroVal T

	if val == "" {
		return zeroVal
	}

	switch any(zeroVal).(type) {
	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return zeroVal
		}
		return any(i).(T)
	case bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			return zeroVal
		}
		return any(b).(T)
	default: //string
		return any(strings.ToLower(val)).(T)
	}
}
