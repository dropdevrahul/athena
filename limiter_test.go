package athena_test

import (
	"testing"
	"time"

	"github.com/dropdevrahul/athena"
	"github.com/stretchr/testify/assert"
)

func Test_TokenLimiter_Get(t *testing.T) {
	t.Parallel()
	l := athena.NewTokenLimiter(100, time.Millisecond*2000)
	for i := 0; i < 100; i++ {
		r := l.Get()
		assert.True(t, r)
	}

	r := l.Get()
	assert.False(t, r)

	time.Sleep(time.Millisecond * 1500)
	r = l.Get()
	assert.False(t, r)

	time.Sleep(time.Millisecond * 500)
	for i := 0; i < 100; i++ {
		r := l.Get()
		assert.True(t, r)
	}
}
