package filecfg

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gopyai/go-err"
)

type (
	Configurator interface {
		// Marshal the configuration content into []bytes
		Marshal() []byte

		// Unmarshal []bytes into configuration content
		Unmarshal(b []byte) error

		// Initialize the configuration content to default setting
		Init()

		// Check the configuration content. If it found invalid value then
		// it should fix the value and return error. If there is no invalid
		// value, then it will return nil.
		CheckAndFix() error
	}
)

// Load configuration from file. Usually at program start. To use it, you
// must first implement the Configurator interface.
func Load(fileName string, cfg Configurator, isRewrite bool) {
	b, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Println("There is no existing configuration file. Default configuration will be created. Please check and update if necessary.")
		initAndExit(fileName, cfg)
	}

	if e := cfg.Unmarshal(b); e != nil {
		fmt.Println("Invalid configuration. New one is created. Please check and update if necessary.")
		initAndExit(fileName, cfg)
	}

	if e := cfg.CheckAndFix(); e != nil {
		fmt.Println(e.Error())
		fmt.Println("Configuration has been fixed. Please check and update if necessary.")
		saveToFile(fileName, cfg.Marshal())
		os.Exit(1)
	}

	if isRewrite {
		saveToFile(fileName, cfg.Marshal())
	}
}

func initAndExit(fileName string, cfg Configurator) {
	cfg.Init()
	saveToFile(fileName, cfg.Marshal())
	os.Exit(1)
}

func saveToFile(fileName string, b []byte) {
	err.Panic(ioutil.WriteFile(fileName, b, 0600))
}
