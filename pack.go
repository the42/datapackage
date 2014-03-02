package datapackage

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	PackRecurse    = 1 << iota
	PackCalcHash   = 1 << iota
	PackInlineData = 1 << iota
)

type Packer struct {
	Visiter
	w        *bufio.Writer
	f        *os.File
	data     Datapackage
	pattern  []string
	calchash bool
	recurse  bool
	inline   bool
}

func newString(s string) *string {
	st := s
	return &st
}

func processFile(path string, info os.FileInfo, calc_hash, inline bool) (*Resource, error) {
	res := &Resource{}
	var rawdata *[]byte

	res.Path = newString(path)
	ext := filepath.Ext(path)
	res.Format = newString(ext)

	res.Name = newString(path) // TODO: realy use the filepath as the Name?

	size := info.Size()
	res.Bytes = &size

	// specialized handling of data types
	switch ext {
	case "csv":
		res.Mediatype = newString("text/csv")
	}

	// generic handling of data types. requires checking if struct members
	// have already been set in the specialized handling above

	if res.Mediatype == nil {
		// TODO: try to determine the media type by eg. unix `file`?
	}

	if calc_hash {
		rawdata, _ = lazyGetData(path, rawdata)
		res.Hash = newString(MD5SumforData(*rawdata))
	}

	if inline && res.Data == nil {
		rawdata, _ = lazyGetData(path, rawdata)

		if ext == "json" {
			// if the data is already json, include it verbatim;
			// no sanity check is happening here; if the input is corrupt,
			// the resulting datapackage file will be corrupt too

			// An intermediate variable is necessary as json.Mashal requires a pointer receiver
			// see http://code.google.com/p/go/issues/detail?id=6528
			imed := json.RawMessage(*rawdata)
			res.Data = &imed
		} else {
			// otherwise base64 encode it. The JSON-serializer will base64 encode our []byte stream
			// so no explicit encoding necessary here
			res.Data = rawdata
		}
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
		res, err := processFile(path, info, p.calchash, p.inline)
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

// NewPacker initialises a new datapackage Packer. The operation mode has to be
// provided using the Pack... constants
func NewPacker(om int) *Packer {
	return &Packer{
		calchash: om&PackCalcHash != 0,
		recurse:  om&PackRecurse != 0,
		inline:   om&PackInlineData != 0,
	}
}
