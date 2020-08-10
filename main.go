package main

import (
	"fmt"

	helloworld "github.com/gdbu/hello-world"
	"github.com/vroomy/common"
)

var hw *helloworld.Controller

// Init will be called by vroomy on initialization
func Init(env map[string]string) (err error) {
	// Get the data directory from the enviroment
	dataDir := env["dataDir"]

	// Initialize hello world controller
	if hw, err = helloworld.New(dataDir); err != nil {
		err = fmt.Errorf("error initializing hello world controller: %v", err)
		return
	}

	return
}

// Load will be called by vroomy after plugin initialization
func Load(p common.Plugins) (err error) {
	// Referencing other plugins occurs here
	return
}

// Backend returns the underlying backend to the plugin
func Backend() interface{} {
	// Currently our plugin does not yet have a controller to reference, return nil
	return nil
}

// Close will close the plugin
func Close() (err error) {
	// Currently our plugin does not yet have a controller to close, return nil
	return nil
}
