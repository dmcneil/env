package env

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPrefix_Get(t *testing.T) {
	prefix := Prefix("FOO")

	tests := []struct {
		value string
		fn    interface{}

		expected interface{}
	}{
		{
			value:    "BAR",
			fn:       prefix.Get,
			expected: "BAR",
		},
		{
			value:    "BAR",
			fn:       prefix.GetString,
			expected: "BAR",
		},
		{
			value:    "1",
			fn:       prefix.Get,
			expected: "1",
		},
		{
			value:    "1",
			fn:       prefix.GetString,
			expected: "1",
		},
		{
			value:    "1",
			fn:       prefix.GetInt,
			expected: 1,
		},
		{
			value:    "1",
			fn:       prefix.GetInt8,
			expected: int8(1),
		},
		{
			value:    "1",
			fn:       prefix.GetInt16,
			expected: int16(1),
		},
		{
			value:    "1",
			fn:       prefix.GetInt32,
			expected: int32(1),
		},
		{
			value:    "1",
			fn:       prefix.GetInt64,
			expected: int64(1),
		},
		{
			value:    "-1",
			fn:       prefix.GetUInt,
			expected: uint(0),
		},
		{
			value:    "1",
			fn:       prefix.GetUInt,
			expected: uint(1),
		},
		{
			value:    "1",
			fn:       prefix.GetUInt8,
			expected: uint8(1),
		},
		{
			value:    "1",
			fn:       prefix.GetUInt16,
			expected: uint16(1),
		},
		{
			value:    "1",
			fn:       prefix.GetUInt32,
			expected: uint32(1),
		},
		{
			value:    "1",
			fn:       prefix.GetUInt64,
			expected: uint64(1),
		},
		{
			value:    "1",
			fn:       prefix.GetFloat32,
			expected: float32(1),
		},
		{
			value:    "1",
			fn:       prefix.GetFloat64,
			expected: float64(1),
		},
		{
			value:    "1.5",
			fn:       prefix.GetInt,
			expected: 0,
		},
		{
			value:    "1.5",
			fn:       prefix.GetInt64,
			expected: int64(0),
		},
		{
			value:    "0",
			fn:       prefix.GetBool,
			expected: false,
		},
		{
			value:    "f",
			fn:       prefix.GetBool,
			expected: false,
		},
		{
			value:    "F",
			fn:       prefix.GetBool,
			expected: false,
		},
		{
			value:    "0",
			fn:       prefix.GetBool,
			expected: false,
		},
		{
			value:    "n",
			fn:       prefix.GetBool,
			expected: false,
		},
		{
			value:    "N",
			fn:       prefix.GetBool,
			expected: false,
		},
		{
			value:    "false",
			fn:       prefix.GetBool,
			expected: false,
		},
		{
			value:    "FALSE",
			fn:       prefix.GetBool,
			expected: false,
		},
		{
			value:    "1",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "t",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "T",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "y",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "Y",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "true",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "TRUE",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "yes",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "YES",
			fn:       prefix.GetBool,
			expected: true,
		},
		{
			value:    "1s",
			fn:       prefix.GetDuration,
			expected: 1 * time.Second,
		},
		{
			value:    "100",
			fn:       prefix.GetDuration,
			expected: time.Duration(0),
		},
	}

	for i, test := range tests {
		fnName := runtime.FuncForPC(reflect.ValueOf(test.fn).Pointer()).Name()
		t.Run(fmt.Sprintf("%d %v", i, fnName), func(t *testing.T) {
			_ = os.Setenv("FOO_FOO", test.value)

			args := []reflect.Value{reflect.ValueOf("FOO")}

			result := reflect.ValueOf(test.fn).Call(args)
			require.IsType(t, test.expected, result[0].Interface())
			require.EqualValues(t, test.expected, result[0].Interface())
		})
	}
}

func TestPrefix_GetD(t *testing.T) {
	prefix := Prefix("FOO")

	tests := []struct {
		value string
		def   interface{}
		fn    interface{}

		expected interface{}
	}{
		{
			value:    "FOO",
			def:      "BAR",
			fn:       prefix.GetStringD,
			expected: "FOO",
		},
		{
			value:    "",
			def:      "BAR",
			fn:       prefix.GetStringD,
			expected: "BAR",
		},
		{
			value:    "FOO",
			def:      "BAR",
			fn:       prefix.GetStringD,
			expected: "FOO",
		},
		{
			value:    "",
			def:      "BAR",
			fn:       prefix.GetStringD,
			expected: "BAR",
		},
		{
			value:    "",
			def:      "1",
			fn:       prefix.GetStringD,
			expected: "1",
		},
		{
			value:    "1",
			def:      "2",
			fn:       prefix.GetStringD,
			expected: "1",
		},
		{
			value:    "",
			def:      1,
			fn:       prefix.GetIntD,
			expected: 1,
		},
		{
			value:    "2.5",
			def:      1,
			fn:       prefix.GetIntD,
			expected: 1,
		},
		{
			value:    "",
			def:      int8(1),
			fn:       prefix.GetInt8D,
			expected: int8(1),
		},
		{
			value:    "2.5",
			def:      int8(1),
			fn:       prefix.GetInt8D,
			expected: int8(1),
		},
		{
			value:    "",
			def:      int16(1),
			fn:       prefix.GetInt16D,
			expected: int16(1),
		},
		{
			value:    "2.5",
			def:      int16(1),
			fn:       prefix.GetInt16D,
			expected: int16(1),
		},
		{
			value:    "",
			def:      int32(1),
			fn:       prefix.GetInt32D,
			expected: int32(1),
		},
		{
			value:    "2.5",
			def:      int32(1),
			fn:       prefix.GetInt32D,
			expected: int32(1),
		},
		{
			value:    "",
			def:      int64(1),
			fn:       prefix.GetInt64D,
			expected: int64(1),
		},
		{
			value:    "2.5",
			def:      int64(1),
			fn:       prefix.GetInt64D,
			expected: int64(1),
		},
		{
			value:    "",
			def:      uint(1),
			fn:       prefix.GetUIntD,
			expected: uint(1),
		},
		{
			value:    "2.5",
			def:      uint(1),
			fn:       prefix.GetUIntD,
			expected: uint(1),
		},
		{
			value:    "",
			def:      uint8(1),
			fn:       prefix.GetUInt8D,
			expected: uint8(1),
		},
		{
			value:    "2.5",
			def:      uint8(1),
			fn:       prefix.GetUInt8D,
			expected: uint8(1),
		},
		{
			value:    "",
			def:      uint16(1),
			fn:       prefix.GetUInt16D,
			expected: uint16(1),
		},
		{
			value:    "2.5",
			def:      uint16(1),
			fn:       prefix.GetUInt16D,
			expected: uint16(1),
		},
		{
			value:    "",
			def:      uint32(1),
			fn:       prefix.GetUInt32D,
			expected: uint32(1),
		},
		{
			value:    "2.5",
			def:      uint32(1),
			fn:       prefix.GetUInt32D,
			expected: uint32(1),
		},
		{
			value:    "",
			def:      uint64(1),
			fn:       prefix.GetUInt64D,
			expected: uint64(1),
		},
		{
			value:    "2.5",
			def:      uint64(1),
			fn:       prefix.GetUInt64D,
			expected: uint64(1),
		},
		{
			value:    "",
			def:      float32(1.5),
			fn:       prefix.GetFloat32D,
			expected: float32(1.5),
		},
		{
			value:    "",
			def:      1.5,
			fn:       prefix.GetFloat64D,
			expected: 1.5,
		},
		{
			value:    "",
			def:      false,
			fn:       prefix.GetBoolD,
			expected: false,
		},
		{
			value:    "",
			def:      true,
			fn:       prefix.GetBoolD,
			expected: true,
		},
		{
			value:    "BAR",
			def:      true,
			fn:       prefix.GetBoolD,
			expected: true,
		},
		{
			value:    "BAR",
			def:      false,
			fn:       prefix.GetBoolD,
			expected: false,
		},
		{
			value:    "5s",
			def:      3 * time.Second,
			fn:       prefix.GetDurationD,
			expected: 5 * time.Second,
		},
		{
			value:    "FOO",
			def:      5 * time.Second,
			fn:       prefix.GetDurationD,
			expected: 5 * time.Second,
		},
	}

	for i, test := range tests {
		fnName := runtime.FuncForPC(reflect.ValueOf(test.fn).Pointer()).Name()
		t.Run(fmt.Sprintf("%d %v", i, fnName), func(t *testing.T) {
			_ = os.Setenv("FOO_FOO", test.value)

			args := []reflect.Value{reflect.ValueOf("FOO"), reflect.ValueOf(test.def)}

			result := reflect.ValueOf(test.fn).Call(args)
			require.IsType(t, test.expected, result[0].Interface())
			require.EqualValues(t, test.expected, result[0].Interface())
		})
	}
}
