package main

import "github.com/vroomy/common"

// Init will be called by vroomy on initialization
func Init(env map[string]string) (err error) {
	// Plugin initialization occurs here
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
