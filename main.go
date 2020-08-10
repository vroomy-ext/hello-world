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
