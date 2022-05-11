package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateId(t *testing.T) {
	InitIdGenerator()
	id1 := GenerateId()
	id2 := GenerateId()
	assert.Equal(t, int64(0), sequenceIdStepMask&id1)
	assert.Equal(t, int64(1), sequenceIdStepMask&id2)
	assert.Equal(t, true, (id1>>sequenceIdStep) == (id2>>sequenceIdStep))
	for i := 0; i < 100000; i++ {
		id1 := GenerateId()
		id2 := GenerateId()
		//检验递增性
		assert.Equal(t, true, id1 < id2)
		//检验是否越界
		assert.Equal(t, true, id1 < (1<<53-1) && id1 > (-1<<53+1))
	}
}
