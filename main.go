package main

import (
	"fmt"

	"github.com/Hatch1fy/httpserve"
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
	// Return hello world controller
	return hw
}

// New is the handler for creating a new Entry
func New(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		e   helloworld.Entry
		err error
	)

	// Parse request body as JSON
	if err = ctx.BindJSON(&e); err != nil {
		// Error parsing request body, return error
		err = fmt.Errorf("error parsing request body: %v", err)
		return httpserve.NewJSONResponse(400, err)
	}

	var createdID string
	// Attempt to insert parsed Entry into the helloworld.Controller
	if createdID, err = hw.New(e); err != nil {
		// Error inserting new Entry, return error
		err = fmt.Errorf("error creating new entry: %v", err)
		return httpserve.NewJSONResponse(400, err)
	}

	// Entry was successfully inserted, return newly created ID
	return httpserve.NewJSONResponse(200, createdID)
}

// Close will close the plugin
func Close() (err error) {
	// Currently our plugin does not yet have a controller to close, return nil
	return hw.Close()
}
