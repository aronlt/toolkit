package toolkit

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

type LineHandler func(string) error

var DefaultPermOpen = os.FileMode(0666)
var LineBreak = []byte{'\n'}

// ReadFile 读取文件的内容
func ReadFile(path string) ([]byte, error) {
	buf := bytes.Buffer{}
	writer := bufio.NewWriter(&buf)

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "call ReadFile fail, open file error")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	_, err = io.Copy(writer, reader)
	if err != nil {
		return nil, errors.Wrap(err, "call ReadFile fail, read file error")
	}
	return buf.Bytes(), nil
}

// WriteFile 将content的数据写入文件
func WriteFile(path string, content []byte, append bool) (int64, error) {
	var writer *os.File
	var err error
	if !append {
		writer, err = os.Create(path)
	} else {
		writer, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
	}
	if err != nil {
		return -1, errors.Wrap(err, "call WriteFile fail, open file error")
	}
	defer writer.Close()

	reader := bytes.NewReader(content)

	n, err := io.Copy(writer, reader)
	if err != nil {
		return -1, errors.Wrap(err, "call WriteFile fail, copy content error")
	}
	return n, nil
}

// ReadLine 按行读取文件内容
// 为了提升性能，使用前需将文件转换为buffer reader
// file, err := os.Open(path)
// if err != nil {
//      return err
// }
// defer file.Close()
// buffer := bufio.NewReader(file)
// ReadLine(buffer)
func ReadLine(buffer *bufio.Reader) ([]byte, error) {
	buf := bytes.Buffer{}
	for {
		data, isPrefix, err := buffer.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if len(data) != 0 {
					buf.Write(data)
				}
				return buf.Bytes(), io.EOF
			} else {
				return nil, errors.Wrap(err, "call ReadLine fail, read line error")
			}
		}
		buf.Write(data)
		if !isPrefix {
			break
		}
	}
	return buf.Bytes(), nil
}

// ReadLines 读取文件内容，按行返回
func ReadLines(path string) ([][]byte, error) {
	content, err := ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "call ReadLines fail")
	}
	return bytes.Split(content, LineBreak), nil
}

// ScanStd 从标准输入读取内容
func ScanStd(handler LineHandler, hints ...string) error {
	if len(hints) > 0 {
		fmt.Sprintln(hints[0])
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		scan := scanner.Text()
		err := handler(scan)
		if err != nil {
			return errors.Wrap(err, "call ScanStd fail, call handler error")
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "call ScanStd fail, scan error")
	}
	return nil
}
