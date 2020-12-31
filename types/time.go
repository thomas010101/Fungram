package types

import (
	"encoding/json"
	"math"
	"strconv"
	"time"
)

const (
	layoutYear   = "2006-01-02"
	layoutSecond = "2006-01-02:15:04:05"
	emptySecond  = Second(0)
)

type Second int64

func NewSecond(s int64) Second {
	if s < 0 {
		s = time.Now().Unix()
	}
	return Second(s)
}

func (s Second) FormatAtYear() string {
	return time.Unix(s.ToInt64(), 0).Format(layoutYear)
}

func (s Second) ToDay() int {
	return int(math.Floor(float64(s) / float64(86400)))
}

func (s Second) FormatAtSecond() string {
	return time.Unix(s.ToInt64(), 0).Format(layoutSecond)
}

func (s Second) ToInt64() int64 {
	return int64(s)
}

func (s Second) Empty() bool {
	return s == emptySecond
}

func (s *Second) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] != '"' {
		var i int64
		if err := json.Unmarshal(data, &i); err != nil {
			return err
		}
		*s = Second(i)
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	*s = Second(i)
	return nil
}
