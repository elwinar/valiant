package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
)

// Load a json file into the destination struct
func Load(filename string, destination interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, destination)
	if err != nil {
		return err
	}

	return nil
}

// Bootstrap a cli.App by using the *Context to generate loggers and load configuration
func Bootstrap(ctx *cli.Context, cfg interface{}) (debug, info, error *log.Logger) {
	var flags int
	if ctx.Bool("verbose") {
		flags = log.Ltime | log.Lshortfile
	} else {
		flags = log.Ltime
	}
	
	error = log.New(os.Stdout, "error ", flags)
	
	if ctx.Bool("debug") {
		debug = log.New(os.Stdout, "debug ", flags)
	} else {
		debug = log.New(ioutil.Discard, "", flags)
	}
	
	if ctx.Bool("quiet") {
		info = log.New(ioutil.Discard, "", flags)
	} else {
		info = log.New(os.Stdout, "info ", flags)
	}
	
	err := Load(ctx.String("configuration"), cfg)
	if err != nil {
		error.Fatalln(err)
	}
	
	return
}
