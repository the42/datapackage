package datapackage

import (
	"path/filepath"
)

type Visiter interface {
	Init() (filepath.WalkFunc, error)
	TearDown() error
}

func Visit(initdir string, vt Visiter) error {
	wp, err := vt.Init()
	if err != nil {
		return err
	}

	if err = filepath.Walk(initdir, wp); err != nil {
		return err
	}
	return vt.TearDown()
}
