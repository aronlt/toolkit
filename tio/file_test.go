package tio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestExistFile(t *testing.T) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "toolkit-")
	assert.Nil(t, err)
	ok, err2 := ExistFile(tmpFile.Name())
	assert.True(t, ok)
	assert.Nil(t, err2)
	os.Remove(tmpFile.Name())
}

func TestExistDir(t *testing.T) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "toolkit-")
	assert.Nil(t, err)
	ok, err2 := ExistDir(tmpDir)
	assert.True(t, ok)
	assert.Nil(t, err2)
	os.Remove(tmpDir)
}

func TestReadDir(t *testing.T) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "toolkit-")
	assert.Nil(t, err)
	err = os.Mkdir(path.Join(tmpDir, "test1"), 0755)
	assert.Nil(t, err)
	_, err = os.Create(path.Join(tmpDir, "test2"))
	assert.Nil(t, err)
	dirs, files, err := ReadDir(tmpDir)
	assert.Nil(t, err)
	assert.Equal(t, path.Join(tmpDir, "test1"), dirs[0])
	assert.Equal(t, path.Join(tmpDir, "test2"), files[0])
	os.Remove(path.Join(tmpDir, "test2"))
	os.Remove(path.Join(tmpDir, "test1"))
	os.Remove(tmpDir)
}
