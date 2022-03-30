package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	t.Run("test cases", func(t *testing.T) {
		got := Convert(map[string]string{"id1": "1", "id2": "2"}, "=", "|")

		assert.Condition(t, func() (success bool) {
			return "id1=1|id2=2" == got || "id2=2|id1=|" == got
		})

		got = Convert(map[string]string{}, "=", "|")
		assert.Equal(t, "", got)

		got = Convert(map[string]string{"id1": "1"}, "=", "|")
		assert.Equal(t, "id1=1", got)
	})
}
