package filecfg

import (
	"fmt"
	"io/ioutil"
	"os"
)

type (
	Configurator interface {
		// Marshal the configuration content into []bytes
		Marshal() []byte

		// Unmarshal []bytes into configuration content
		Unmarshal(b []byte) error

		// Initialize the configuration content to default setting
		Init()

		// Check the configuration content.
		// - If it found invalid value then it should fix the value and return true.
		// - If everything is ok, then it will return false.
		CheckAndFix() bool
	}
)

// Load configuration from file. Usually at program start. To use it, you
// must first implement the Configurator interface.
func Load(fileName string, cfg Configurator, isRewrite bool) {
	b, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Println("There is no existing configuration file. Default configuration will be created. Please review.")
		initAndExit(fileName, cfg)
	}

	if e := cfg.Unmarshal(b); e != nil {
		fmt.Println("Invalid configuration. New one is created. Please review.")
		initAndExit(fileName, cfg)
	}

	if cfg.CheckAndFix() {
		fmt.Println("Please review/fix the configuration to continue.")
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
	panicIf(ioutil.WriteFile(fileName, b, 0777))
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
