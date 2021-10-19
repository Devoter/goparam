package param

import (
	"flag"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"time"
)

// Bool defines a bool parameter with specified parameter name, environment variable name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag or the enviroment variable.
func Bool(envName string, paramName string, defaultValue bool, description string) *bool {
	param := defaultValue
	envVar, isSet := os.LookupEnv(envName)
	if isSet {
		val, err := strconv.ParseBool(envVar)
		exit(envName, "bool", err)
		param = val
	}

	return flag.Bool(paramName, param, description)
}

// Duration defines a time.Duration parameter with specified parameter name, environment variable name,
// default value, and usage string.
// The return value is the address of a time.Durarion variable that stores the value of the flag or the enviroment variable.
func Duration(envName string, paramName string, defaultValue time.Duration, description string) *time.Duration {
	param := defaultValue
	envVar, isSet := os.LookupEnv(envName)
	if isSet {
		val, err := strconv.ParseInt(envVar, 10, 64)
		exit(envName, "time.Duration", err)
		param = time.Duration(val)
	}

	return flag.Duration(paramName, param, description)
}

// Int defines an int parameter with specified parameter name, environment variable name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag or the enviroment variable.
func Int(envName string, paramName string, defaultValue int, description string) *int {
	param := defaultValue
	envVar, isSet := os.LookupEnv(envName)
	if isSet {
		val, err := strconv.Atoi(envVar)
		exit(envName, "int", err)
		param = val
	}

	return flag.Int(paramName, param, description)
}

// Int64 defines an int64 parameter with specified parameter name, environment variable name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag or the enviroment variable.
func Int64(envName string, paramName string, defaultValue int64, description string) *int64 {
	param := defaultValue
	envVar, isSet := os.LookupEnv(envName)
	if isSet {
		val, err := strconv.ParseInt(envVar, 10, 64)
		exit(envName, "int64", err)
		param = val
	}

	return flag.Int64(paramName, param, description)
}

// Uint defines a uint parameter with specified parameter name, environment variable name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag or the enviroment variable.
func Uint(envName string, paramName string, defaultValue uint, description string) *uint {
	param := defaultValue
	envVar, isSet := os.LookupEnv(envName)
	if isSet {
		val, err := strconv.ParseUint(envVar, 10, bits.UintSize)
		exit(envName, "uint", err)
		param = uint(val)
	}

	return flag.Uint(paramName, param, description)
}

// Uint64 defines a uint64 parameter with specified parameter name, environment variable name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag or the enviroment variable.
func Uint64(envName string, paramName string, defaultValue uint64, description string) *uint64 {
	param := defaultValue
	envVar, isSet := os.LookupEnv(envName)
	if isSet {
		val, err := strconv.ParseUint(envVar, 10, 64)
		exit(envName, "uint64", err)
		param = val
	}

	return flag.Uint64(paramName, param, description)
}

// Float64 defines a float64 parameter with specified parameter name, environment variable name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag or the enviroment variable.
func Float64(envName string, paramName string, defaultValue float64, description string) *float64 {
	param := defaultValue
	envVar, isSet := os.LookupEnv(envName)
	if isSet {
		val, err := strconv.ParseFloat(envVar, 10)
		exit(envName, "float64", err)
		param = val
	}

	return flag.Float64(paramName, param, description)
}

// String defines a string parameter with specified parameter name, environment variable name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag or the enviroment variable.
// ATTENSION: Be careful, empty string is valid.
func String(envName string, paramName string, defaultValue string, desctiption string) *string {
	param := defaultValue
	envVar, isSet := os.LookupEnv(envName)
	if isSet {
		param = envVar
	}

	return flag.String(paramName, param, desctiption)
}

// Parse parses the command-line flags from os.Args[1:].
func Parse() {
	flag.Parse()
}

// PrintDefaults prints, to standard error unless configured otherwise, a usage message showing the default settings of all defined command-line flags.
func PrintDefaults() {
	flag.PrintDefaults()
}

// exit outputs an error if `err` is not null and stops the program with status 2.
func exit(name string, typeName string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal=\"'"+name+"' environment variable shoul be '"+typeName+"'\", error=\"%s\"\n", err.Error())
		os.Exit(2)
	}
}
