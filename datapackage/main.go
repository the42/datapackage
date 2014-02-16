package main

import (
	"fmt"
	"github.com/the42/datapackage"
)

func main() {
	p := datapackage.NewPacker(true, true)
	err := datapackage.Visit(".", p)
	if err != nil {
		fmt.Println(err)
	}
}
