package main

import "os"

func commandExit(config *Config, arg string) error {
	os.Exit(0)
	return nil
}
