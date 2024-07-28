package main

import (
	"os"
)

type Args struct {
	initialize bool
}

func GetArgs() Args {
	result := Args{}
	args := os.Args[1:]
	if len(args) == 0 {
		return result
	}
	switch args[0] {
	case "--initialize":
		return Args{
			initialize: true,
		}
	}
	return Args{}
}
