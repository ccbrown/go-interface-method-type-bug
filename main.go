package main

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/loader"
)

func main() {
	conf := &loader.Config{
		ImportPkgs: map[string]bool{
			"./thepackage": true,
		},
	}

	prog, err := conf.Load()
	if err != nil {
		panic(err)
	}

	for _, pkg := range prog.Imported {
		for _, obj := range pkg.Defs {
			if fn, ok := obj.(*types.Func); ok {
				fmt.Println(fn)
			}
		}
	}
}
