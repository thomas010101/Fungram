package types

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	defaultAuthTimeout = 480 * 60 * 60
)

type Authorization struct {
	AccountId uint64 `json:"-"`
	Auth      string `json:"auth"`
	TailTime  int64  `json:"tail_time"`
}

func (u *Authorization) Encode() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Authorization) Decode(b []byte) error {
	if err := json.Unmarshal(b, u); err != nil {
		return err
	}
	return nil
}

func (u *Authorization) ToJWT() map[string]interface{} {
	return map[string]interface{}{
		"id":   u.AccountId,
		"auth": u.Auth,
	}
}

func (u *Authorization) Valid(now Authorization) error {
	if u.Auth != now.Auth {
		return fmt.Errorf("auth (%s != %s)", u.Auth, now.Auth)
	}
	if !(u.TailTime > 0 && u.TailTime <= now.TailTime && u.TailTime+defaultAuthTimeout > now.TailTime) {
		fmt.Println(u.TailTime > 0)
		fmt.Println(u.TailTime <= now.TailTime)
		fmt.Println(u.TailTime < now.TailTime+defaultAuthTimeout)
		return fmt.Errorf("auth current(%d) req(%d) default(%d)",
			u.TailTime, now.TailTime, defaultAuthTimeout)
	}
	return nil
}

func ParseJwtToUserAuthorization(m map[string]interface{}) (Authorization, error) {
	tmpId, ok := m["id"]
	if !ok {
		return Authorization{}, errors.New("id not found")
	}
	id, ok := tmpId.(float64)
	if !ok {
		return Authorization{}, errors.New("id type parse failed")
	}
	tmpAuth, ok := m["auth"]
	if !ok {
		return Authorization{}, errors.New("auth not found")
	}
	auth, ok := tmpAuth.(string)
	if !ok {
		return Authorization{}, errors.New("auth type parse failed")
	}
	return Authorization{
		AccountId: uint64(id),
		Auth:      auth,
		TailTime:  NewSecond(-1).ToInt64(),
	}, nil
}
