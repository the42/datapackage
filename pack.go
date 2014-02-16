package datapackage

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Packer struct {
	Visiter
	w        *bufio.Writer
	f        *os.File
	recurse  bool
	data     Datapackage
	pattern  []string
	calchash bool
}

func newString(s string) *string {
	st := s
	return &st
}

func processFile(path string, info os.FileInfo, calc_hash bool) (*Resource, error) {
	res := &Resource{}
	var rawdata *[]byte

	res.Path = newString(path)
	ext := filepath.Ext(path)
	res.Format = newString(ext)

	res.Name = newString(path) // TODO: realy use the filepath as the Name?

	size := info.Size()
	res.Bytes = &size

	switch ext {
	case "csv":
		res.Mediatype = newString("text/csv")
	}

	if res.Mediatype == nil {
		// TODO: try to determine the media type by eg. unix `file`?
	}

	if calc_hash {
		if rawdata == nil {
			rdata, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, err
			}
			rawdata = &rdata
		}
		res.Hash = newString(MD5SumforData(*rawdata))
	}

	res.Modified = &ISO8601{Time: info.ModTime()}

	return res, nil
}

func (p *Packer) walkerFunc(path string, info os.FileInfo, err error) error {
	if path == "." || path == Filename {
		return nil
	}
	// when it's a directory, check if directories should be processed recursively
	if info.IsDir() {
		if !p.recurse {
			return filepath.SkipDir
		} else {
			// directory hit. Signal to recurse
			return nil
		}
	}
	// os it's a regular file, process it
	match := true
	// If a pattern should be matched, check if the received file matches the pattern
	if len(p.pattern) > 0 {
		var err error
		for _, pattern := range p.pattern {
			match, err = filepath.Match(pattern, path)
			if err != nil {
				return err
			}
			if match {
				break
			}
		}
	}

	if match {
		res, err := processFile(path, info, p.calchash)
		if err != nil {
			return err
		}
		p.data.Resources = append(p.data.Resources, *res)
	}
	return nil
}

func (p *Packer) Init() (filepath.WalkFunc, error) {
	// TODO: check if realy sensible
	// if a datapackage.json - file already exists, fail
	f, err := os.OpenFile(Filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return nil, err
	}
	p.w = bufio.NewWriter(f)
	p.f = f
	return p.walkerFunc, nil
}

func (p *Packer) TearDown() error {
	enc := json.NewEncoder(p.w)
	if err := enc.Encode(p.data); err != nil {
		return err
	}
	p.w.Flush()
	return p.f.Close()
}

func NewPacker(recurse, calchash bool) *Packer {
	return &Packer{calchash: calchash, recurse: recurse}
}
