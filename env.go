package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Get wraps os.Getenv.
func Get(key string) string {
	return os.Getenv(key)
}

// Lookup wraps os.LookupEnv.
func Lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}

// GetD attempts to retrieve a string named by key. If the value is not present, def
// is returned instead.
func GetD(key string, def string) string {
	v, ok := Lookup(key)
	if !ok || v == "" {
		return def
	}
	return v
}

// GetString retrieves a string named by key.
// It is functionally the same as Get.
func GetString(key string) string {
	return Get(key)
}

// GetStringD attempts to retrieve a string named by key. If the value is not present, def is returned instead.
// It is functionally the same as GetD.
func GetStringD(key string, def string) string {
	return GetD(key, def)
}

// GetInt retrieves an int named by key.
// int(0) is returned if the value does not exist or is not a valid int.
func GetInt(key string) int {
	return parseInt(Get(key))
}

// GetIntD attempts to retrieve an int named by key. If the value is not present, def is returned instead.
func GetIntD(key string, def int) int {
	if v := parseInt(GetD(key, fmt.Sprintf("%d", def))); v != 0 {
		return v
	}
	return def
}

// GetInt64 retrieves an int64 named by key.
// int64(0) is returned if the value does not exist or is not a valid int64.
func GetInt64(key string) int64 {
	return parseInt64(Get(key))
}

// GetInt64D attempts to retrieve an int64 named by key. If the value is not present, def is returned instead.
func GetInt64D(key string, def int64) int64 {
	if v := parseInt64(GetD(key, fmt.Sprintf("%d", def))); v != 0 {
		return v
	}
	return def
}

// GetFloat32 retrieves a float32 named by key.
// float32(0) is returned if the value does not exist or is not a valid float32.
func GetFloat32(key string) float32 {
	return parseFloat32(Get(key))
}

// GetFloat32D attempts to retrieve a float32 named by key. If the value is not present, def is returned instead.
func GetFloat32D(key string, def float32) float32 {
	if v := parseFloat32(GetD(key, fmt.Sprintf("%f", def))); v != 0 {
		return v
	}
	return def
}

// GetFloat64 retrieves a float64 named by key.
// float64(0) is returned if the value does not exist or is not a valid float64.
func GetFloat64(key string) float64 {
	return parseFloat64(Get(key))
}

// GetFloat64D attempts to retrieve a float64 named by key. If the value is not present, def is returned instead.
func GetFloat64D(key string, def float64) float64 {
	if v := parseFloat64(GetD(key, fmt.Sprintf("%f", def))); v != 0 {
		return v
	}
	return def
}

// GetBool retrieves a bool named by key.
// bool(false) is returned if the value does not exist or is not a valid bool.
func GetBool(key string) bool {
	b, _ := parseBool(Get(key))
	return b
}

// GetBoolD attempts to retrieve a bool named by key. If the value is not present, def is returned instead.
func GetBoolD(key string, def bool) bool {
	b, ok := parseBool(GetD(key, fmt.Sprintf("%t", def)))
	if !ok {
		return def
	}
	return b
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func parseFloat32(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

func parseFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func parseBool(s string) (bool, bool) {
	switch strings.ToUpper(s) {
	case "1", "T", "Y", "TRUE", "YES":
		return true, true
	case "0", "F", "N", "FALSE", "NO":
		return false, true
	}

	return false, false
}
