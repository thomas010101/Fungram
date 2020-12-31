package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	floatRegexp        = regexp.MustCompile("^[0-9]*\\.?[0-9]*$")
	integerTooLargeErr = fmt.Errorf("integer too large")
)

// 绝对值
func Abs(i int64) int64 {
	y := i >> 63
	c := (i ^ y) - y
	return c
}

// 取相反值
func Turn(a int64) int64 {
	a = ^a + 1
	return a
}

func Int64ToString(u int64, precision, showPrecision int) string {
	c := Abs(u)
	s := Uint64ToString(uint64(c), precision, showPrecision)
	if u < 0 {
		return "-" + s
	}
	return s
}

func StringToInt64(s string, precision int) (int64, error) {
	if len(s) == 0 {
		return 0, nil
	}
	var isTurn bool
	if s[0] == '-' {
		isTurn = true
		s = s[1:]
	}
	u, err := StringToUint64(s, precision)
	if err != nil {
		return 0, err
	}
	if u > math.MaxInt64 {
		return 0, fmt.Errorf("greater than maxInt64")
	}
	i := int64(u)
	if isTurn {
		i = Turn(i)
	}
	return i, nil
}

func Uint64ToString(u uint64, precision, showPrecision int) string {
	if u == 0 {
		return "0"
	}
	s := strconv.FormatUint(u, 10)
	if precision == 0 {
		return s
	}
	sl := len(s)
	var (
		integer  string
		tDecimal string
	)
	//var isZero bool
	if sl > precision {
		integer = s[:sl-precision]
		tDecimal = s[sl-precision:]
		//return fmt.Sprintf("%s.%s", s[:sl-precision], s[sl-precision:])
	} else if sl == precision {
		integer = "0"
		tDecimal = s
		//isZero = true
		//return fmt.Sprintf("0.%s", s)
	} else {
		integer = "0"
		tDecimal = fmt.Sprintf("%s%s", strings.Repeat("0", precision-sl), s)
		//isZero = true
		//return fmt.Sprintf("0.%s%s", strings.Repeat("0", precision-sl), s)
	}
	if precision <= showPrecision {
		return fmt.Sprintf("%s.%s", integer, tDecimal)
	}
	if showPrecision >= 0 {
		tDecimal = tDecimal[:showPrecision]
	} else {
		for len(tDecimal) > 0 {
			if tDecimal[len(tDecimal)-1] == '0' {
				tDecimal = tDecimal[:len(tDecimal)-1]
			} else {
				break
			}
		}
	}
	if len(tDecimal) == 0 {
		return integer
	}
	return fmt.Sprintf("%s.%s", integer, tDecimal)
}

func FloatToUint64(f float64, precision int) (uint64, error) {
	s := big.NewFloat(f).Text('f', precision+1)
	return StringToUint64(s, precision)
}

func StringToUint64(s string, precision int) (uint64, error) {
	if len(s) == 0 || s == "0" {
		return 0, nil
	}
	if precision > 19 {
		return 0, integerTooLargeErr
	}
	if !floatRegexp.MatchString(s) {
		return 0, fmt.Errorf("parse wrong: %s", s)
	}
	splitStr := strings.Split(s, ".")
	var integer uint64
	if len(splitStr[0]) == 0 || splitStr[0] == "0" {
		integer = 0
	} else {
		if len(splitStr[0])+precision > 19 {
			return 0, integerTooLargeErr
		}
		var err error
		integer, err = strconv.ParseUint(splitStr[0], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("parse interger: %s", s)
		}
		rate := uint64(math.Pow10(precision))
		integer = integer * rate
	}
	if len(splitStr) == 1 || len(splitStr[1]) == 0 || splitStr[1] == "0" {
		return integer, nil
	}
	decimalLen := len(splitStr[1])
	if decimalLen > precision {
		decimalLen = precision
		splitStr[1] = splitStr[1][:precision]
	}
	de, err := strconv.ParseUint(splitStr[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse decimal: %s", s)
	}
	if decimalLen == precision {
		return integer + de, nil
	}
	rate := uint64(math.Pow10(precision - decimalLen))
	return integer + de*rate, nil
}

func Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func Decode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}

func DecimalFixed(d decimal.Decimal, places int32) decimal.Decimal {
	ua := decimal.New(1, places)
	return d.Mul(ua).Floor().Div(ua)
}

func StrToDecimalFixed(s string, places int32) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Decimal{}, fmt.Errorf("string to decimal: %w", err)
	}
	return DecimalFixed(d, places), nil
}

func ToBytes(val ...interface{}) []byte {
	data := make([]byte, 0)
	for _, v := range val {
		switch v := v.(type) {
		case string:
			data = append(data, []byte(v)...)
		case int32:
			tmp := make([]byte, binary.MaxVarintLen32)
			n := binary.PutVarint(tmp, int64(v))
			data = append(data, tmp[:n]...)
		case int64:
			tmp := make([]byte, binary.MaxVarintLen64)
			n := binary.PutVarint(tmp, v)
			data = append(data, tmp[:n]...)
		}
	}
	return data
}
