package thttp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddQuery(t *testing.T) {
	rawUrl := "https://www.baidu.com"
	url, err := AddQuery(rawUrl, "name", "lin")
	assert.Nil(t, err)
	assert.Equal(t, url, rawUrl+"?name=lin")

	url, err = AddQuery(rawUrl, "name", "lin==")
	assert.Nil(t, err)
	assert.Equal(t, url, rawUrl+"?name=lin%3D%3D")
}

func TestQueryUrl(t *testing.T) {
	rawUrl := "https://www.baidu.com"
	url, err := AddQuery(rawUrl, "name", "lin")
	assert.Nil(t, err)
	v, err := QueryUrl(url, "name")
	assert.Nil(t, err)
	assert.Equal(t, v, "lin")

	url, err = AddQuery(rawUrl, "name", "lin==")
	assert.Nil(t, err)
	v, err = QueryUrl(url, "name")
	assert.Nil(t, err)
	assert.Equal(t, v, "lin==")
}
