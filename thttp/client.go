package thttp

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type Option struct {
	Decompress bool
	Header     map[string]string
}

func Decompress(option Option, resp []byte) ([]byte, error) {
	if !option.Decompress {
		return resp, nil
	}
	r, err := gzip.NewReader(bytes.NewReader(resp))
	if err != nil {
		return resp, err
	}
	defer r.Close()
	r2, err := io.ReadAll(r)
	return r2, err
}

// Get 根据path请求资源
func Get(u string, options ...Option) ([]byte, error) {
	req, _ := http.NewRequest("GET", u, nil)
	var option Option
	if len(options) != 0 {
		option = options[0]
	}

	for k, v := range option.Header {
		req.Header.Set(k, v)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "call http.Get() fail")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error-> status = %d", resp.StatusCode)
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "call io.ReadAll fail")
	}
	bs, err = Decompress(option, bs)
	if err != nil {
		return nil, errors.WithMessage(err, "call Decompress fail")
	}
	return bs, nil
}

// GetToMap 请求资源，以map形式返回结果
func GetToMap(u string, options ...Option) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	bs, err := Get(u, options...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return nil, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}

// GetToStruct 请求资源，以struct形式返回结果
func GetToStruct[T any](u string, options ...Option) (T, error) {
	var data T
	bs, err := Get(u, options...)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return data, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}

// PostForm 以form格式请求
func PostForm(u string, form url.Values, options ...Option) ([]byte, error) {
	var resp *http.Response
	var err error
	if len(options) == 0 {
		resp, err = http.PostForm(u, form)
	} else {
		req, err := http.NewRequest("POST", u, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		for k, v := range options[0].Header {
			req.Header.Set(k, v)
		}
		resp, err = (&http.Client{}).Post(u, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	}
	if err != nil {
		return nil, errors.WithMessage(err, "call http.PostForm fail")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post json error-> status = %d", resp.StatusCode)
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "call io.ReadAll fail")
	}
	bs, err = Decompress(options[0], bs)
	if err != nil {
		return nil, errors.WithMessage(err, "call Decompress fail")
	}
	return bs, nil
}

func PostJSONFromStruct(u string, obj interface{}, options ...Option) ([]byte, error) {
	content, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.WithMessagef(err, "call json.Marshal fail, obj:%+v", obj)
	}
	return PostJSON(u, content, options...)
}

// PostJSON 以json格式请求
func PostJSON(u string, jsonByte []byte, options ...Option) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, errors.WithMessage(err, "call http.NewRequest fail")
	}
	req.Header.Set("Content-Type", "application/json")
	var option Option
	if len(options) > 0 {
		option = options[0]
	}

	for k, v := range option.Header {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "call http client.Do fail")
	}
	defer rsp.Body.Close()
	b, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "call io.ReadAll fail")
	}
	b, err = Decompress(option, b)
	if err != nil {
		return nil, errors.WithMessage(err, "call Decompress fail")
	}
	return b, nil
}

func PostToMapFromStruct(u string, obj interface{}, options ...Option) (map[string]interface{}, error) {
	content, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.WithMessagef(err, "call json.Marshal fail, obj:%+v", obj)
	}
	return PostToMap(u, content, options...)
}

func PostToMap(u string, jsonByte []byte, options ...Option) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	bs, err := PostJSON(u, jsonByte, options...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return nil, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}

func PostToStructFromStruct[T any](u string, obj interface{}, options ...Option) (T, error) {
	content, err := json.Marshal(obj)
	if err != nil {
		var data T
		return data, errors.WithMessagef(err, "call json.Marshal fail, obj:%+v", obj)
	}
	return PostToStruct[T](u, content, options...)
}

func PostToStruct[T any](u string, jsonByte []byte, options ...Option) (T, error) {
	var data T
	bs, err := PostJSON(u, jsonByte, options...)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return data, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}
