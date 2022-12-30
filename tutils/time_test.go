package tutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnixToString(t *testing.T) {
	ut := int64(1672363746)
	s := UnixToString(int64(ut))
	assert.Equal(t, "2022-12-30 09:29:06", s)
}

func TestUnixToTime(t *testing.T) {
	ut := int64(1672363746)
	tm := UnixToTime(ut)
	assert.Equal(t, ut, tm.Unix())
}

func TestStringToTime(t *testing.T) {
	s := "2022-12-30 09:29:06"
	tm := StringToTime(s)
	assert.Equal(t, TimeToString(tm), s)
}

func TestStringToUnix(t *testing.T) {
	s := "2022-12-30 09:29:06"
	ut := StringToUnix(s)
	assert.Equal(t, int64(1672363746), ut)
}

func TestTimeToString(t *testing.T) {
	tm := UnixToTime(1672363746)
	assert.Equal(t, TimeToString(tm), "2022-12-30 09:29:06")
}

func TestTimeToUnix(t *testing.T) {
	tm := UnixToTime(1672363746)
	assert.Equal(t, TimeToUnix(tm), int64(1672363746))
}

func TestGetDateAgo(t *testing.T) {
	tm := StringToTime("2022-12-30 09:29:06")
	s := GetDayOffset(tm, -1)
	assert.Equal(t, "2022-12-29 09:29:06", s)
}

func TestGetDaysDiff(t *testing.T) {
	d1 := "2022-12-30 09:29:06"
	d2 := "2022-12-29 09:29:06"
	diff := GetDaysDiff(StringToTime(d1), StringToTime(d2))
	assert.Equal(t, 1, diff)
}
