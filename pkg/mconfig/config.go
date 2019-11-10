package mconfig

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
	DataPath  string
	LogEnable bool
}

var instance *MConfig
var initialized uint32
var mu sync.Mutex

// Get returns singleton of config of an application
func Get() (*MConfig, error) {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance, nil
	}

	mu.Lock()
	defer mu.Unlock()

	var err error

	if initialized == 0 {
		instance, err = load()
		if err != nil {
			atomic.StoreUint32(&initialized, 1)
		}
	}

	return instance, err
}

func getConfigPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir, err
}

func load() (*MConfig, error) {
	dir, err := getConfigPath()

	data, err := ioutil.ReadFile(filepath.Join(dir, "config.json"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result := new(MConfig)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	if result.LogEnable {
		log.Printf("Config is loaded successfully %v\n", result)
		fmt.Printf("Config directory: %v\n", dir)
	}

	return result, nil
}
