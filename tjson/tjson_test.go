package tjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testV = []byte(`
{
  "key1": 1,
  "key2": "value2",
  "key3": true,
  "key4": 0.1,
  "key5": {
    "key6": "value6"
  },
  "key7": [
    "key7_1",
    "key7_2",
    "key7_3"
  ],
  "key8": {
    "key9": [
      1,
      2,
      3
    ]
  },
  "key10": {
    "key11": [
      {
        "key12": "value12",
        "key13": "value13"
      },
      {
        "key14": "value14",
        "key15": "value15"
      }
    ]
  }
}
`)

func TestGetFloat(t *testing.T) {
	v, err := GetFloat(testV, "key4")
	assert.Nil(t, err)
	assert.Equal(t, v, 0.1)
}

func TestGetInt(t *testing.T) {
	v, err := GetInt(testV, "key1")
	assert.Nil(t, err)
	assert.Equal(t, v, 1)
	v2, err := GetIntSlice(testV, "key8.key9")
	assert.Nil(t, err)
	assert.Equal(t, v2, []int{1, 2, 3})

	v3 := GetIntWithDefault(testV, "key5.kk", 7)
	assert.Equal(t, v3, 7)
}

func TestGetString(t *testing.T) {
	v, err := GetString(testV, "key5.key6")
	assert.Nil(t, err)
	assert.Equal(t, v, "value6")
	v2, err := GetStringSlice(testV, "key7")
	assert.Nil(t, err)
	assert.Equal(t, v2, []string{"key7_1", "key7_2", "key7_3"})

	v3 := GetStringWithDefault(testV, "key5.kk", "abc")
	assert.Equal(t, v3, "abc")

	v4, err := GetString(testV, "key7[1]")
	assert.Nil(t, err)
	assert.Equal(t, v4, "key7_2")
}

var testV2 = []byte(`
{
	"key1": [
       {
			"key2": "a"
		},
		{
			"key2": "b"
		}
	]
}
`)

func TestGetEvery(t *testing.T) {
	v, err := GetEverySlice(testV2, "key1[*].key2")
	assert.Nil(t, err, nil)
	all := make([]string, 0)
	for _, v2 := range v.([]interface{}) {
		v3 := v2.(json.RawMessage)
		var v4 string
		err = json.Unmarshal(v3, &v4)
		assert.Nil(t, err)
		all = append(all, v4)
	}
	assert.Equal(t, all, []string{"a", "b"})
}
