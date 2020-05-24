package main

import "errors"

var (
	ErrNotJSON = errors.New("Not a JSON file.")
)
