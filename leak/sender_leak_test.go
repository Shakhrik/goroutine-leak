package leak

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	// uber -> goleak
	// defer goleak.VerifyNone(t)
	// leaktest
	// defer leaktest.Check(t)()

	err := process("some value")
	time.Sleep(2 * time.Second)
	// assert.NoError(t, err)
	assert.Error(t, err, "search canceled")
}
