package timehelpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	time0 = time.Unix(0, 0)
	time1 = time.Unix(1526432013, 123456789)
)

func TestUnixMillisecond(t *testing.T) {
	ms0 := UnixMillisecond(time0)
	assert.EqualValues(t, int64(0), ms0)

	ms1 := UnixMillisecond(time1)
	assert.EqualValues(t, int64(1526432013123), ms1)

	assert.EqualValues(t, int64(0), FromUnixMillisecond(ms0).UnixNano())
	assert.EqualValues(t, int64(1526432013123000000), FromUnixMillisecond(ms1).UnixNano())
}

func TestUnixMicrosecond(t *testing.T) {
	ms0 := UnixMicrosecond(time0)
	assert.EqualValues(t, int64(0), ms0)

	ms1 := UnixMicrosecond(time1)
	assert.EqualValues(t, int64(1526432013123456), ms1)

	assert.EqualValues(t, int64(0), FromUnixMicrosecond(ms0).UnixNano())
	assert.EqualValues(t, int64(1526432013123456000), FromUnixMicrosecond(ms1).UnixNano())
}

func TestUnixNanosecond(t *testing.T) {
	ms0 := UnixNanosecond(time0)
	assert.EqualValues(t, int64(0), ms0)

	ms1 := UnixNanosecond(time1)
	assert.EqualValues(t, int64(1526432013123456789), ms1)

	assert.EqualValues(t, int64(0), FromUnixNanosecond(ms0).UnixNano())
	assert.EqualValues(t, int64(1526432013123456789), FromUnixNanosecond(ms1).UnixNano())
}

func TestUnixSecond(t *testing.T) {
	ms0 := UnixSecond(time0)
	assert.EqualValues(t, int64(0), ms0)

	ms1 := UnixSecond(time1)
	assert.EqualValues(t, int64(1526432013), ms1)

	assert.EqualValues(t, int64(0), FromUnixSecond(ms0).UnixNano())
	assert.EqualValues(t, int64(1526432013000000000), FromUnixSecond(ms1).UnixNano())
}
