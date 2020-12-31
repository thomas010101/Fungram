package types

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Int64 int64

func NewInt64(i int64) Int64 {
	return Int64(i)
}

func (c *Int64) Decode(b []byte) error {
	i, _ := binary.Varint(b)
	*c = Int64(i)
	return nil
}

func (c Int64) Encode() ([]byte, error) {
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(b, c.Int64())
	return b[:n], nil
}

func (c Int64) Int64() int64 {
	return int64(c)
}

func (c Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(c), 10))
}

func (c *Int64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	switch data[0] {
	case '"':
		if len(data) < 2 {
			return nil
		}
		data = data[1 : len(data)-1]
	}
	u, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*c = NewInt64(u)
	return nil
}

type Uint64 uint64

func NewUInt64(i uint64) Uint64 {
	return Uint64(i)
}

func (c Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(c), 10))
}

func (c Uint64) String() string {
	return strconv.FormatUint(c.Uint64(), 10)
}

func (c Uint64) Empty() bool {
	return c == 0
}

func (c *Uint64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	switch data[0] {
	case '"':
		if len(data) < 2 {
			return nil
		}
		data = data[1 : len(data)-1]
	}
	u, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}
	*c = NewUInt64(u)

	return nil
}

func (c Uint64) Uint64() uint64 {
	return uint64(c)
}

type UpperString string

func NewUpperString(s string) UpperString {
	return UpperString(strings.ToUpper(s))
}

func (s UpperString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *UpperString) UnmarshalJSON(b []byte) error {
	if b[0] != '"' || len(b) < 2 {
		return fmt.Errorf("type wrong(%s)", string(b))
	}
	*s = NewUpperString(string(b[1 : len(b)-1]))
	return nil
}

func (s *UpperString) DecodeStr(st string) error {
	*s = NewUpperString(st)
	return nil
}

func (s UpperString) String() string {
	return string(s)
}
