package env

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

var env sync.Map

func init() {
	for _, e := range os.Environ() {
		splits := strings.Split(e, "=")
		env.Store(splits[0], os.Getenv(splits[0]))
	}
	s_logger.Debug(os.Environ())
	s_logger.Debug("init done")
	// env.Store("GOBIN", GetGOBIN())   // GOBIN is the path to the go binary
	// env.Store("GOPATH", GetGOPATH()) // GOPATH is the path to the go source code
}

// func getGOENV(string, error) {
// 	if file := os.Getenv("GOENV"); file != "" {
// 		if file == "off" {
// 			return "", fmt.Errorf("GOENV=off")
// 		}
// 		return file, nil
// 	}
// 	dir, err := os.UserConfigDir()
// 	if err != nil {
// 		return "", err
// 	}
// 	if dir == "" {
// 		return "", fmt.Errorf("missing user-config dir")
// 	}
// 	file, err := filepath.Join(dir, "go", "env")
// 	if err != nil {
// 		return "", err
// 	}
// 	if file == "" {
// 		return "", fmt.Errorf("missing runtime env file")
// 	}

// 	var runtimeEnv string
// 	data, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		return "", err
// 	}
// 	envStrings := strings.Split(string(data), "\n")
// 	for _, envItem := range envStrings {
// 		envItem = strings.TrimSuffix(envItem, "\r")
// 		envKeyValue := strings.Split(envItem, "=")
// 		if len(envKeyValue) == 2 && strings.TrimSpace(envKeyValue[0]) == key {
// 			runtimeEnv = strings.TrimSpace(envKeyValue[1])
// 		}
// 	}
// 	return runtimeEnv, nil
// }

// func GetGOBIN() string {
// 	// The one set by user explicitly by `export GOBIN=/path` or `env GOBIN=/path command`
// 	gobin := strings.TrimSpace(Get("GOBIN", ""))
// 	if gobin == "" {
// 		var err error
// 		// The one set by user by running `go env -w GOBIN=/path`
// 		gobin, err = GetRuntimeEnv("GOBIN")
// 		if err != nil {
// 			// The default one that Golang uses
// 			return filepath.Join(build.Default.GOPATH, "bin")
// 		}
// 		if gobin == "" {
// 			return filepath.Join(build.Default.GOPATH, "bin")
// 		}
// 		return gobin
// 	}
// 	return gobin
// }

// // GetGOPATH returns GOPATH environment variable as a string.
// // It will NOT be an empty string.
// func GetGOPATH() string {
// 	// The one set by user explicitly by `export GOPATH=/path` or `env GOPATH=/path command`
// 	gopath := strings.TrimSpace(Get("GOPATH", ""))
// 	if gopath == "" {
// 		var err error
// 		// The one set by user by running `go env -w GOPATH=/path`
// 		gopath, err = GetRuntimeEnv("GOPATH")
// 		if err != nil {
// 			// The default one that Golang uses
// 			return build.Default.GOPATH
// 		}
// 		if gopath == "" {
// 			return build.Default.GOPATH
// 		}
// 		return gopath
// 	}
// 	return gopath
// }

func Get[T int | string](key string, defaultValue T) T {
	if val, ok := env.Load(key); ok {
		switch any(defaultValue).(type) {
		case int:
			if i, err := strconv.Atoi(val.(string)); err == nil {
				return T(i)
			}

		}
		return val.(T)
	}
	return defaultValue
}

func IsPrd() bool {
	return Get(runEnv, prd) == prd
}

func IsDev() bool {
	return Get(runEnv, prd) == dev
}

func IsStg() bool {
	return Get(runEnv, prd) == stg
}

func CurrEnv() string {
	return Get(runEnv, prd)
}

func Port() int {
	return Get("port", 8080)
}
