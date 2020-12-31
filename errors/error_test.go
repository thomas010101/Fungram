package errors

import (
	"errors"
	"fmt"
	"testing"
	"wallet/types"

	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	e := fmt.Errorf("err : %w", UnknownErr)
	c := As(e)
	require.True(t, unknown == c.code)
	e = fmt.Errorf("success: %w", Success)
	c = As(e)
	require.True(t, success == c.code)
}

func TestMError_With(t *testing.T) {
	require.True(t, "成功%!(EXTRA int=1, int=2, int=3)" == Success.With(1, 2, 3).Msg(types.ZHCNLanguage))
}

func TestMError_WithError(t *testing.T) {
	require.True(t, Success.Msg(types.ZHCNLanguage) == "成功")
	require.True(t, Success.Format("format: %w", errors.New("err")).Error() == "format: err")
	require.True(t, Success.Format("format err").Error() == "format err")
	require.True(t, Success.Format("format err %d", 12).Error() == "format err 12")
}
