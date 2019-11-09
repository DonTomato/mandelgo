package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

// MConfig is a configuration structure
type MConfig struct {
	dataPath  string
	logEnable bool
}

var instance *MConfig
var initialized uint32
var mu sync.Mutex

// Get returns singleton of config of an application
func Get() *MConfig {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		*instance = load()
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

func getConfigPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func load() MConfig {
	dir := getConfigPath()

	data, le := ioutil.ReadFile(filepath.Join(dir, "config.json"))
	if le != nil {
		log.Fatal(le)
		panic(le)
	}

	result := MConfig{}
	json.Unmarshal(data, &result)

	if result.logEnable {
		log.Printf("Config is loaded successfully %v\n", result)
		fmt.Printf("Config directory: %v\n", dir)
	}

	return result
}
