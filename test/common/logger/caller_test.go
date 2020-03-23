package logger_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/ver13/ava/pkg/common/logger/formatter"
)

func TestCallerName(t *testing.T) {
	assert := require.New(t)

	assert.Equal("TestCallerName", CallerName(1))
	assert.Equal("TestCallerName", func() string { return CallerName(2) }())
}
