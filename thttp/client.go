package thttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Get 根据path请求资源
func Get(u string, headers ...map[string]string) ([]byte, error) {
	req, _ := http.NewRequest("GET", u, nil)

	for _, header := range headers {
		for k, v := range header {
			req.Header.Set(k, v)
		}
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
	return bs, nil
}

// GetToMap 请求资源，以map形式返回结果
func GetToMap(u string, header ...map[string]string) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	bs, err := Get(u, header...)
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
func GetToStruct[T any](u string, header ...map[string]string) (T, error) {
	var data T
	bs, err := Get(u, header...)
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
func PostForm(u string, form url.Values, headers ...map[string]string) ([]byte, error) {
	var resp *http.Response
	var err error
	if len(headers) == 0 {
		resp, err = http.PostForm(u, form)
	} else {
		req, err := http.NewRequest("POST", u, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		for _, header := range headers {
			for k, v := range header {
				req.Header.Set(k, v)
			}
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
	return bs, nil
}

// PostJSON 以json格式请求
func PostJSON(u string, jsonByte []byte, headers ...map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, errors.WithMessage(err, "call http.NewRequest fail")
	}
	req.Header.Set("Content-Type", "application/json")
	for _, header := range headers {
		for k, v := range header {
			req.Header.Add(k, v)
		}
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
	return b, nil
}

func PostToMap(u string, jsonByte []byte, header ...map[string]string) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	bs, err := PostJSON(u, jsonByte, header...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return nil, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}

func PostToStruct[T any](u string, jsonByte []byte, header ...map[string]string) (T, error) {
	var data T
	bs, err := PostJSON(u, jsonByte, header...)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return data, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}
