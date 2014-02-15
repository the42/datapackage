package datapackage

import (
	"io"
	"io/ioutil"
)

type Visiter interface {
	Init() error
	Open(file string) (io.Closer, error)
	Process(io.Closer) error
	TearDown() error
}

func Visit(initdir string, recurse bool, vt Visiter) error {
	files, err := ioutil.ReadDir(initdir)
	if err != nil {
		return err
	}

	err = vt.Init()
	if err != nil {
		return err
	}

	for idx, file := range files {
		_ = idx
		if recurse && file.IsDir() {
			Visit(file.Name(), recurse, vt)
		} else {
			rc, err := vt.Open(file.Name())
			if err != nil {
				return err
			}
			if err = vt.Process(rc); err != nil {
				return err
			}
			rc.Close()
		}
	}
	err = vt.TearDown()
	return err
}
