package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Get wraps os.Getenv.
func Get(key string) string {
	return os.Getenv(key)
}

// Lookup wraps os.LookupEnv.
func Lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}

// GetD attempts to retrieve a string named by key. If the value is not present, def is returned instead.
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
	if v := GetString(key); v != "" {
		return v
	}
	return def
}

// GetInt retrieves an int named by key.
// int(0) is returned if the value does not exist or is not a valid int.
func GetInt(key string) int {
	return parseInt(Get(key))
}

// GetIntD attempts to retrieve an int named by key. If the value is not present, def is returned instead.
func GetIntD(key string, def int) int {
	if v := GetInt(key); v != 0 {
		return v
	}
	return def
}

// GetInt8 retrieves an int8 named by key.
// int8(0) is returned if the value does not exist or is not a valid int8.
func GetInt8(key string) int8 {
	return parseInt8(Get(key))
}

// GetInt8D attempts to retrieve an int8 named by key. If the value is not present, def is returned instead.
func GetInt8D(key string, def int8) int8 {
	if v := GetInt8(key); v != 0 {
		return v
	}
	return def
}

// GetInt16 retrieves an int16 named by key.
// int16(0) is returned if the value does not exist or is not a valid int16.
func GetInt16(key string) int16 {
	return parseInt16(Get(key))
}

// GetInt16D attempts to retrieve an int16 named by key. If the value is not present, def is returned instead.
func GetInt16D(key string, def int16) int16 {
	if v := GetInt16(key); v != 0 {
		return v
	}
	return def
}

// GetInt32 retrieves an int32 named by key.
// int32(0) is returned if the value does not exist or is not a valid int32.
func GetInt32(key string) int32 {
	return parseInt32(Get(key))
}

// GetInt32D attempts to retrieve an int32 named by key. If the value is not present, def is returned instead.
func GetInt32D(key string, def int32) int32 {
	if v := GetInt32(key); v != 0 {
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
	if v := GetInt64(key); v != 0 {
		return v
	}
	return def
}

// GetUInt retrieves an uint named by key.
// uint(0) is returned if the value does not exist or is not a valid uint.
func GetUInt(key string) uint {
	return parseUInt(Get(key))
}

// GetUIntD attempts to retrieve an uint named by key. If the value is not present, def is returned instead.
func GetUIntD(key string, def uint) uint {
	if v := GetUInt(key); v != 0 {
		return v
	}
	return def
}

// GetUInt8 retrieves an uint8 named by key.
// uint8(0) is returned if the value does not exist or is not a valid uint8.
func GetUInt8(key string) uint8 {
	return parseUInt8(Get(key))
}

// GetUInt8D attempts to retrieve an uint8 named by key. If the value is not present, def is returned instead.
func GetUInt8D(key string, def uint8) uint8 {
	if v := GetUInt8(key); v != 0 {
		return v
	}
	return def
}

// GetUInt16 retrieves an uint16 named by key.
// uint16(0) is returned if the value does not exist or is not a valid uint16.
func GetUInt16(key string) uint16 {
	return parseUInt16(Get(key))
}

// GetUInt16D attempts to retrieve an uint16 named by key. If the value is not present, def is returned instead.
func GetUInt16D(key string, def uint16) uint16 {
	if v := GetUInt16(key); v != 0 {
		return v
	}
	return def
}

// GetUInt32 retrieves an uint32 named by key.
// uint32(0) is returned if the value does not exist or is not a valid uint32.
func GetUInt32(key string) uint32 {
	return parseUInt32(Get(key))
}

// GetUInt32D attempts to retrieve an uint32 named by key. If the value is not present, def is returned instead.
func GetUInt32D(key string, def uint32) uint32 {
	if v := GetUInt32(key); v != 0 {
		return v
	}
	return def
}

// GetUInt64 retrieves an uint64 named by key.
// uint64(0) is returned if the value does not exist or is not a valid uint64.
func GetUInt64(key string) uint64 {
	return parseUInt64(Get(key))
}

// GetUInt64D attempts to retrieve an uint64 named by key. If the value is not present, def is returned instead.
func GetUInt64D(key string, def uint64) uint64 {
	if v := GetUInt64(key); v != 0 {
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
	if v := GetFloat32(key); v != 0 {
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
	if v := GetFloat64(key); v != 0 {
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
	b, ok := parseBool(GetStringD(key, fmt.Sprintf("%t", def)))
	if !ok {
		return def
	}
	return b
}

// GetDuration retrieves a time.Duration named by key.
// time.Duration(0) is returned if the value does not exist or is not parsed as a valid time.Duration.
func GetDuration(key string) time.Duration {
	d, _ := time.ParseDuration(Get(key))
	return d
}

// GetDurationD attempts to retrieve a time.Duration named by key. If the value is not present, def is returned instead.
func GetDurationD(key string, def time.Duration) time.Duration {
	d, err := time.ParseDuration(Get(key))
	if err != nil {
		return def
	}
	return d
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parseInt8(s string) int8 {
	i, _ := strconv.ParseInt(s, 10, 8)
	return int8(i)
}

func parseInt16(s string) int16 {
	i, _ := strconv.ParseInt(s, 10, 16)
	return int16(i)
}

func parseInt32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int32(i)
}

func parseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func parseUInt(s string) uint {
	return uint(parseUInt64(s))
}

func parseUInt8(s string) uint8 {
	i, _ := strconv.ParseUint(s, 10, 8)
	return uint8(i)
}

func parseUInt16(s string) uint16 {
	i, _ := strconv.ParseUint(s, 10, 16)
	return uint16(i)
}

func parseUInt32(s string) uint32 {
	i, _ := strconv.ParseUint(s, 10, 32)
	return uint32(i)
}

func parseUInt64(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
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
