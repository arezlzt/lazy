package snowflake

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	worker, err := NewWorker(1)
	if err != nil {
		assert.Error(t, err)
	}
	for i := 0; i < 1000; i++ {
		t.Log(worker.GetId())
	}
}
