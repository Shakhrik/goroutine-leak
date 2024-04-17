package leak

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestProcessRecords(t *testing.T) {
	// // uber -> goleak
	// defer goleak.VerifyNone(t)
	// // leaktest
	// defer leaktest.Check(t)()
	recs := []string{"rec1", "rec2", "rec3"}
	s := processRecords(recs, 2)
	time.Sleep(5 * time.Second)
	assert.Equal(t, s, "some value")
}
