package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFloatRegexp(t *testing.T) {
	args := []struct {
		s      string
		result bool
	}{
		{s: "123", result: true},
		{s: "12.3", result: true},
		{s: "123.", result: true},
		{s: ".123", result: true},
		{s: "0.123", result: true},
		{s: "a123", result: false},
		{s: "12a3", result: false},
		{s: "1a23", result: false},
		{s: "123.a", result: false},
		{s: "123b", result: false},
	}
	for _, v := range args {
		require.True(t, floatRegexp.MatchString(v.s) == v.result)
	}
}

func TestStringToUint64(t *testing.T) {
	args := []struct {
		s         string
		precision int
		result    uint64
		err       error
	}{
		{s: "1.43", precision: 8, result: 143000000, err: nil},
		{s: "1,43", precision: 8, result: 0, err: fmt.Errorf("parse wrong: %s", "1,43")},
		{s: "a143", precision: 8, result: 0, err: fmt.Errorf("parse wrong: %s", "a143")},
		{s: "143a", precision: 8, result: 0, err: fmt.Errorf("parse wrong: %s", "143a")},
		{s: "143", precision: 8, result: 14300000000, err: nil},
		{s: ".143", precision: 8, result: 14300000, err: nil},
		{s: "143.", precision: 8, result: 14300000000, err: nil},
		{s: "143.", precision: 17, result: 0, err: integerTooLargeErr},
		{s: "143", precision: 17, result: 0, err: integerTooLargeErr},
		{s: ".143", precision: 20, result: 0, err: integerTooLargeErr},
		{s: "0.143", precision: 20, result: 0, err: integerTooLargeErr},
		{s: ".143", precision: 19, result: 1430000000000000000, err: nil},
		{s: "0.143", precision: 19, result: 1430000000000000000, err: nil},
		{s: "143987654323", precision: 8, result: 0, err: integerTooLargeErr},
		{s: "143.98765", precision: 3, result: 143987, err: nil},
		{s: "0.98765", precision: 3, result: 987, err: nil},
		{s: ".98765", precision: 3, result: 987, err: nil},
		{s: "", precision: 20, result: 0, err: nil},
		{s: "0", precision: 20, result: 0, err: nil},
		{s: "0.1", precision: 20, result: 0, err: integerTooLargeErr},
	}
	for _, v := range args {
		r, err := StringToUint64(v.s, v.precision)
		if err != nil {
			require.NotNil(t, v.err, fmt.Sprintf("s: %s, err :%s, expectErr: %s", v.s, err, v.err))
			require.True(t, err.Error() == v.err.Error())
		}
		require.True(t, r == v.result, r)
	}
}

func TestUint64ToString(t *testing.T) {
	args := []struct {
		u             uint64
		precision     int
		showPrecision int
		result        string
	}{
		{u: 123, precision: 2, showPrecision: 2, result: "1.23"},
		{u: 123, precision: 5, showPrecision: 3, result: "0.001"},
		{u: 123, precision: 2, showPrecision: 5, result: "1.23"},
		{u: 123, precision: 0, showPrecision: 5, result: "123"},
		{u: 12300, precision: 4, showPrecision: -1, result: "1.23"},
		{u: 12300, precision: 2, showPrecision: -1, result: "123"},
		{u: 12300, precision: 0, showPrecision: -1, result: "12300"},
	}
	for k, v := range args {
		r := Uint64ToString(v.u, v.precision, v.showPrecision)
		require.True(t, r == v.result, fmt.Sprintf("index(%d): %s != %s", k, r, v.result))
	}
}
