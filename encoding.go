package toolkit

import (
	"bytes"
	"io"
	"sort"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type EUnit[K any] struct {
	Raw   K
	Field []byte
}

func sortBytes[K any](units []EUnit[K]) {
	sort.Slice(units, func(i, j int) bool {
		switch bytes.Compare(units[i].Field, units[j].Field) {
		case -1:
			return true
		case 1, 0:
			return false
		default:
			panic("bytes.Compare invalid return")
		}
	})
}

func encodeToGBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// SortFieldsWithGbk 按照中文排序
func SortFieldsWithGbk[K any](units []EUnit[K], offset int, limit int, desc bool) ([]EUnit[K], error) {
	gbkUnits := make([]EUnit[K], 0, len(units))
	for _, unit := range units {
		content, err := encodeToGBK(unit.Field)
		if err != nil {
			return gbkUnits, err
		}
		gbkUnits = append(gbkUnits, EUnit[K]{
			Raw:   unit.Raw,
			Field: content,
		})
	}
	sortBytes(gbkUnits)
	if desc {
		ReverseSlice(gbkUnits)
	}
	if offset >= len(gbkUnits) {
		return gbkUnits, nil
	} else {
		if offset+limit >= len(gbkUnits) {
			return gbkUnits[offset:], nil
		} else {
			return gbkUnits[offset : offset+limit], nil
		}
	}
}
