package env

import (
	"time"
)

// Prefix retrieves the values of environment variables with a common prefix.
// An underscore will be appended to the Prefix when retreiving values.
// For example, Prefix("FOO").GetString("BAR") would return the value of FOO_BAR.
type Prefix string

// Get wraps os.Getenv.
func (p Prefix) Get(key string) string {
	return Get(p.format(key))
}

// GetD attempts to retrieve a string named by key. If the value is not present, def is returned instead.
func (p Prefix) GetD(key string, def string) string {
	return GetD(p.format(key), def)
}

// Lookup wraps os.LookupEnv.
func (p Prefix) Lookup(key string) (string, bool) {
	return Lookup(p.format(key))
}

// GetString retrieves a string named by key.
// It is functionally the same as Get.
func (p Prefix) GetString(key string) string {
	return GetString(p.format(key))
}

// GetStringD attempts to retrieve a string named by key. If the value is not present, def is returned instead.
// It is functionally the same as GetD.
func (p Prefix) GetStringD(key string, def string) string {
	return GetStringD(p.format(key), def)
}

// GetInt retrieves an int named by key.
// int(0) is returned if the value does not exist or is not a valid int.
func (p Prefix) GetInt(key string) int {
	return GetInt(p.format(key))
}

// GetIntD attempts to retrieve an int named by key. If the value is not present, def is returned instead.
func (p Prefix) GetIntD(key string, def int) int {
	return GetIntD(p.format(key), def)
}

// GetInt8 retrieves an int8 named by key.
// int8(0) is returned if the value does not exist or is not a valid int8.
func (p Prefix) GetInt8(key string) int8 {
	return GetInt8(p.format(key))
}

// GetInt8D attempts to retrieve an int8 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetInt8D(key string, def int8) int8 {
	return GetInt8D(p.format(key), def)
}

// GetInt16 retrieves an int16 named by key.
// int16(0) is returned if the value does not exist or is not a valid int16.
func (p Prefix) GetInt16(key string) int16 {
	return GetInt16(p.format(key))
}

// GetInt16D attempts to retrieve an int16 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetInt16D(key string, def int16) int16 {
	return GetInt16D(p.format(key), def)
}

// GetInt32 retrieves an int32 named by key.
// int32(0) is returned if the value does not exist or is not a valid int32.
func (p Prefix) GetInt32(key string) int32 {
	return GetInt32(p.format(key))
}

// GetInt32D attempts to retrieve an int32 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetInt32D(key string, def int32) int32 {
	return GetInt32D(p.format(key), def)
}

// GetInt64 retrieves an int64 named by key.
// int64(0) is returned if the value does not exist or is not a valid int64.
func (p Prefix) GetInt64(key string) int64 {
	return GetInt64(p.format(key))
}

// GetInt64D attempts to retrieve an int64 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetInt64D(key string, def int64) int64 {
	return GetInt64D(p.format(key), def)
}

// GetUInt retrieves an uint named by key.
// uint(0) is returned if the value does not exist or is not a valid uint.
func (p Prefix) GetUInt(key string) uint {
	return GetUInt(p.format(key))
}

// GetUIntD attempts to retrieve an uint named by key. If the value is not present, def is returned instead.
func (p Prefix) GetUIntD(key string, def uint) uint {
	return GetUIntD(p.format(key), def)
}

// GetUInt8 retrieves an uint8 named by key.
// uint8(0) is returned if the value does not exist or is not a valid uint8.
func (p Prefix) GetUInt8(key string) uint8 {
	return GetUInt8(p.format(key))
}

// GetUInt8D attempts to retrieve an uint8 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetUInt8D(key string, def uint8) uint8 {
	return GetUInt8D(p.format(key), def)
}

// GetUInt16 retrieves an uint16 named by key.
// uint16(0) is returned if the value does not exist or is not a valid uint16.
func (p Prefix) GetUInt16(key string) uint16 {
	return GetUInt16(p.format(key))
}

// GetUInt16D attempts to retrieve an uint16 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetUInt16D(key string, def uint16) uint16 {
	return GetUInt16D(p.format(key), def)
}

// GetUInt32 retrieves an uint32 named by key.
// uint32(0) is returned if the value does not exist or is not a valid uint32.
func (p Prefix) GetUInt32(key string) uint32 {
	return GetUInt32(p.format(key))
}

// GetUInt32D attempts to retrieve an uint32 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetUInt32D(key string, def uint32) uint32 {
	return GetUInt32D(p.format(key), def)
}

// GetUInt64 retrieves an uint64 named by key.
// uint64(0) is returned if the value does not exist or is not a valid uint64.
func (p Prefix) GetUInt64(key string) uint64 {
	return GetUInt64(p.format(key))
}

// GetUInt64D attempts to retrieve an uint64 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetUInt64D(key string, def uint64) uint64 {
	return GetUInt64D(p.format(key), def)
}

// GetFloat32 retrieves a float32 named by key.
// float32(0) is returned if the value does not exist or is not a valid float32.
func (p Prefix) GetFloat32(key string) float32 {
	return GetFloat32(p.format(key))
}

// GetFloat32D attempts to retrieve a float32 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetFloat32D(key string, def float32) float32 {
	return GetFloat32D(p.format(key), def)
}

// GetFloat64 retrieves a float64 named by key.
// float64(0) is returned if the value does not exist or is not a valid float64.
func (p Prefix) GetFloat64(key string) float64 {
	return GetFloat64(p.format(key))
}

// GetFloat64D attempts to retrieve a float64 named by key. If the value is not present, def is returned instead.
func (p Prefix) GetFloat64D(key string, def float64) float64 {
	return GetFloat64D(p.format(key), def)
}

// GetBool retrieves a bool named by key.
// bool(false) is returned if the value does not exist or is not a valid bool.
func (p Prefix) GetBool(key string) bool {
	return GetBool(p.format(key))
}

// GetBoolD attempts to retrieve a bool named by key. If the value is not present, def is returned instead.
func (p Prefix) GetBoolD(key string, def bool) bool {
	return GetBoolD(p.format(key), def)
}

// GetDuration retrieves a time.Duration named by key.
// time.Duration(0) is returned if the value does not exist or is not parsed as a valid time.Duration.
func (p Prefix) GetDuration(key string) time.Duration {
	return GetDuration(p.format(key))
}

// GetDurationD attempts to retrieve a time.Duration named by key. If the value is not present, def is returned instead.
func (p Prefix) GetDurationD(key string, def time.Duration) time.Duration {
	return GetDurationD(p.format(key), def)
}

func (p Prefix) format(key string) string {
	return string(p) + "_" + key
}
