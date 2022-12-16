// copy from badger project
// https://github.com/dgraph-io/badger

package bytes

import (
	"fmt"
	"math"
	"strconv"
)

// IBytesToString converts size in bytes to human readable format.
// The code is taken from humanize library and changed to provide
// value upto custom decimal precision.
// IBytesToString(12312412, 1) -> 11.7 MiB
func IBytesToString(size uint64, precision int) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	base := float64(1024)
	if size < 10 {
		return fmt.Sprintf("%dB", size)
	}
	e := math.Floor(math.Log(float64(size)) / math.Log(base))
	suffix := sizes[int(e)]
	val := float64(size) / math.Pow(base, e)
	f := "%." + strconv.Itoa(precision) + "f%s"

	return fmt.Sprintf(f, val, suffix)
}
