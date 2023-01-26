package thttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

// Get 根据path请求资源
func Get(u string) ([]byte, error) {
	resp, err := http.Get(u)
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
	return bs, nil
}

// GetToMap 请求资源，以map形式返回结果
func GetToMap(u string) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	bs, err := Get(u)
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
func GetToStruct[T any](u string) (T, error) {
	var data T
	bs, err := Get(u)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return nil, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}

// PostForm 以form格式请求
func PostForm(u string, form url.Values) ([]byte, error) {
	resp, err := http.PostForm(u, form)
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
	return bs, nil
}

// PostJSON 以json格式请求
func PostJSON(u string, jsonByte []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, errors.WithMessage(err, "call http.NewRequest fail")
	}
	req.Header.Set("Content-Type", "application/json")
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
	return b, nil
}

func PostToMap(u string, jsonByte []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	bs, err := PostJSON(u, jsonByte)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return nil, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}

func PostToStruct[T any](u string, jsonByte []byte) (T, error) {
	var data T
	bs, err := PostJSON(u, jsonByte)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return data, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}
