package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToString(t *testing.T) {
	t.Parallel()

	t.Run("When function receives a nil value", func(t *testing.T) {
		assert.Equal(t, "", ConvertToString(nil))
	})

	t.Run("When function receives a string as value", func(t *testing.T) {
		assert.Equal(t, "string", ConvertToString("string"))
	})

	t.Run("When function receives int", func(t *testing.T) {
		assert.Equal(t, "1", ConvertToString(1))
	})

	t.Run("When function receives int64", func(t *testing.T) {
		assert.Equal(t, "1", ConvertToString(int64(1)))
	})

	t.Run("When default behavior", func(t *testing.T) {
		assert.Equal(t, "", ConvertToString([]int{1}))
	})
}

func TestConvertToInt64(t *testing.T) {
	t.Parallel()

	t.Run("When function receives a nil value", func(t *testing.T) {
		assert.Equal(t, int64(0), ConvertToInt64(nil))
	})

	t.Run("When function receives a string as value", func(t *testing.T) {
		assert.Equal(t, int64(1), ConvertToInt64("1"))
	})

	t.Run("When function receives an invalid string value", func(t *testing.T) {
		assert.Equal(t, int64(0), ConvertToInt64("&"))
	})

	t.Run("When function receives an int as value", func(t *testing.T) {
		assert.Equal(t, int64(1), ConvertToInt64(1))
	})

	t.Run("When function receives an int64 as value", func(t *testing.T) {
		assert.Equal(t, int64(1), ConvertToInt64(int64(1)))
	})

}

func TestConvertNumberToBool(t *testing.T) {
	t.Parallel()

	t.Run("When function receives a '1' value", func(t *testing.T) {
		assert.Equal(t, true, ConvertNumberToBool(1))
	})

	t.Run("When function receives a value different from '1'", func(t *testing.T) {
		assert.Equal(t, false, ConvertNumberToBool(0))
	})
}
