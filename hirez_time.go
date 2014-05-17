package smiteapi

import (
	"encoding/json"
	"time"
)

const Format = "1/2/2006 3:4:5 PM"

type SmiteTime time.Time

func (st *SmiteTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(Format, s)
	if err != nil {
		return err
	}
	*st = SmiteTime(t)
	return nil
}

func (st SmiteTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(st).Format(Format))
}
