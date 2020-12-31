package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Language int

const DefaultLanguage = ZHCNLanguage

const (
	_ Language = iota
	ZHCNLanguage
	ENUSLanguage
)

var languages = []Language{
	ZHCNLanguage,
	ENUSLanguage,
}

func (l Language) String() string {
	switch l {
	case ZHCNLanguage:
		return "zh-ch"
	case ENUSLanguage:
		return "en-us"
	}
	return ""
}

func (l Language) Tag() string {
	return l.String()
}

func ToLanguage(s string) Language {
	switch s {
	case "zh":
		return ZHCNLanguage
	}
	for _, v := range languages {
		if strings.EqualFold(s, v.String()) {
			return v
		}
	}
	return DefaultLanguage
}

func (l *Language) UnmarshalJSON(b []byte) error {
	var buf string
	if err := json.Unmarshal(b, &buf); err != nil {
		return fmt.Errorf("parse language: %w", err)
	}
	*l = ToLanguage(buf)
	return nil
}
