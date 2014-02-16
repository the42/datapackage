package datapackage

import (
	"encoding/json"
	"time"
)

const ISO8601Format1 = "2006-01-02"
const ISO8601Format2 = "2006-01-02T15:04:05Z"

func (t *ISO8601) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	t.Raw = raw
	pt, err := time.Parse(ISO8601Format1, raw)
	if err != nil {
		if pt, err = time.Parse(ISO8601Format2, raw); err != nil {
			if pt, err = time.Parse(time.RFC3339, raw); err != nil {
				return err
			}
		}
	}
	t.Time = pt
	return nil
}

func (t *ISO8601) MarshalJSON() ([]byte, error) {
	s := t.Time.UTC().Format(ISO8601Format2)
	return json.Marshal(s)
}

func (pk *PKSpec) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] != '[' {
		var s string
		e := json.Unmarshal(data, &s)
		if e != nil {
			return e
		}
		pk.Keys = append(pk.Keys, s)
		return nil
	}
	return json.Unmarshal(data, (*[]string)(&pk.Keys))
}
