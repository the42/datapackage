package main

import (
	"fmt"
	"github.com/the42/datapackage"
)

func main() {
	p := datapackage.NewPacker(datapackage.PackRecurse | datapackage.PackCalcHash)
	err := datapackage.Visit(".", p)
	if err != nil {
		fmt.Println(err)
	}
}
