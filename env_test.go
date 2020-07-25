package denv

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"runtime"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		value string
		fn    interface{}

		expected interface{}
	}{
		{
			value:    "BAR",
			fn:       Get,
			expected: "BAR",
		},
		{
			value:    "BAR",
			fn:       GetString,
			expected: "BAR",
		},
		{
			value:    "1",
			fn:       Get,
			expected: "1",
		},
		{
			value:    "1",
			fn:       GetString,
			expected: "1",
		},
		{
			value:    "1",
			fn:       GetInt,
			expected: 1,
		},
		{
			value:    "1",
			fn:       GetInt64,
			expected: int64(1),
		},
		{
			value:    "1",
			fn:       GetFloat32,
			expected: float32(1),
		},
		{
			value:    "1",
			fn:       GetFloat64,
			expected: float64(1),
		},
		{
			value:    "1.5",
			fn:       GetInt,
			expected: 0,
		},
		{
			value:    "1.5",
			fn:       GetInt64,
			expected: int64(0),
		},
		{
			value:    "0",
			fn:       GetBool,
			expected: false,
		},
		{
			value:    "f",
			fn:       GetBool,
			expected: false,
		},
		{
			value:    "F",
			fn:       GetBool,
			expected: false,
		},
		{
			value:    "0",
			fn:       GetBool,
			expected: false,
		},
		{
			value:    "n",
			fn:       GetBool,
			expected: false,
		},
		{
			value:    "N",
			fn:       GetBool,
			expected: false,
		},
		{
			value:    "false",
			fn:       GetBool,
			expected: false,
		},
		{
			value:    "FALSE",
			fn:       GetBool,
			expected: false,
		},
		{
			value:    "1",
			fn:       GetBool,
			expected: true,
		},
		{
			value:    "t",
			fn:       GetBool,
			expected: true,
		},
		{
			value:    "T",
			fn:       GetBool,
			expected: true,
		},
		{
			value:    "y",
			fn:       GetBool,
			expected: true,
		},
		{
			value:    "Y",
			fn:       GetBool,
			expected: true,
		},
		{
			value:    "true",
			fn:       GetBool,
			expected: true,
		},
		{
			value:    "TRUE",
			fn:       GetBool,
			expected: true,
		},
		{
			value:    "yes",
			fn:       GetBool,
			expected: true,
		},
		{
			value:    "YES",
			fn:       GetBool,
			expected: true,
		},
	}

	for i, test := range tests {
		fnName := runtime.FuncForPC(reflect.ValueOf(test.fn).Pointer()).Name()
		t.Run(fmt.Sprintf("%d %v", i, fnName), func(t *testing.T) {
			_ = os.Setenv("FOO", test.value)

			args := []reflect.Value{reflect.ValueOf("FOO")}

			result := reflect.ValueOf(test.fn).Call(args)
			require.IsType(t, test.expected, result[0].Interface())
			require.EqualValues(t, test.expected, result[0].Interface())
		})
	}
}

func TestGetD(t *testing.T) {
	tests := []struct {
		value string
		def   interface{}
		fn    interface{}

		expected interface{}
	}{
		{
			value:    "FOO",
			def:      "BAR",
			fn:       GetD,
			expected: "FOO",
		},
		{
			value:    "",
			def:      "BAR",
			fn:       GetD,
			expected: "BAR",
		},
		{
			value:    "FOO",
			def:      "BAR",
			fn:       GetStringD,
			expected: "FOO",
		},
		{
			value:    "",
			def:      "BAR",
			fn:       GetStringD,
			expected: "BAR",
		},
		{
			value:    "",
			def:      "1",
			fn:       GetD,
			expected: "1",
		},
		{
			value:    "1",
			def:      "2",
			fn:       GetStringD,
			expected: "1",
		},
		{
			value:    "",
			def:      1,
			fn:       GetIntD,
			expected: 1,
		},
		{
			value:    "2.5",
			def:      1,
			fn:       GetIntD,
			expected: 1,
		},
		{
			value:    "",
			def:      int64(1),
			fn:       GetInt64D,
			expected: int64(1),
		},
		{
			value:    "2.5",
			def:      int64(1),
			fn:       GetInt64D,
			expected: int64(1),
		},
		{
			value:    "",
			def:      float32(1.5),
			fn:       GetFloat32D,
			expected: float32(1.5),
		},
		{
			value:    "",
			def:      1.5,
			fn:       GetFloat64D,
			expected: 1.5,
		},
		{
			value:    "",
			def:      false,
			fn:       GetBoolD,
			expected: false,
		},
		{
			value:    "",
			def:      true,
			fn:       GetBoolD,
			expected: true,
		},
		{
			value:    "BAR",
			def:      true,
			fn:       GetBoolD,
			expected: true,
		},
		{
			value:    "BAR",
			def:      false,
			fn:       GetBoolD,
			expected: false,
		},
	}

	for i, test := range tests {
		fnName := runtime.FuncForPC(reflect.ValueOf(test.fn).Pointer()).Name()
		t.Run(fmt.Sprintf("%d %v", i, fnName), func(t *testing.T) {
			_ = os.Setenv("FOO", test.value)

			args := []reflect.Value{reflect.ValueOf("FOO"), reflect.ValueOf(test.def)}

			result := reflect.ValueOf(test.fn).Call(args)
			require.IsType(t, test.expected, result[0].Interface())
			require.EqualValues(t, test.expected, result[0].Interface())
		})
	}
}
