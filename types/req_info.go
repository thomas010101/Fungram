package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Origin int64

const (
	Default Origin = iota
	Android
	Iphone
)

var (
	OriginLs = []Origin{
		Android,
		Iphone,
		Default,
	}
)

func ParseOrigin(useragent string) Origin {
	//获取请求来源（安卓或苹果或者其他）
	if strings.Contains(useragent, "iPhone") || strings.Contains(useragent, "iPad") {
		return Iphone
	} else if strings.Contains(useragent, "Android") {
		return Android
	} else {
		return Default
	}
}

func (o *Origin) String() string {
	switch *o {
	case Android:
		return "android"
	case Iphone:
		return "iphone"
	default:
		return "default"
	}
}

func (o *Origin) Uint8() uint8 {
	return uint8(*o)
}

func (o *Origin) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

func (o *Origin) DecodeStr(s string) error {
	for _, v := range OriginLs {
		if strings.EqualFold(v.String(), s) {
			*o = v
		}
	}
	return nil
}

func (o *Origin) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	for _, v := range OriginLs {
		if strings.EqualFold(v.String(), tmp) {
			*o = v
		}
	}
	return nil
}
