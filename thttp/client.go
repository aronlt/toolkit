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

	"github.com/aronlt/toolkit/treflect"
)

type Config struct {
	Decompress bool
	Header     map[string]string
}

type Option func(*Config)

func WithHeaderMap(header map[string]interface{}) Option {
	return func(config *Config) {
		if config == nil {
			return
		}
		if config.Header == nil {
			config.Header = make(map[string]string, len(header))
		}
		for key, value := range header {
			s, ok := treflect.ToString(value)
			if ok {
				config.Header[key] = s
			}
		}
	}
}

func WithHeaderKV(key string, value interface{}) Option {
	return func(config *Config) {
		if config == nil {
			return
		}
		if config.Header == nil {
			config.Header = make(map[string]string)
		}
		s, ok := treflect.ToString(value)
		if ok {
			config.Header[key] = s
		}
	}
}

func WithDecompress(v bool) Option {
	return func(config *Config) {
		if config == nil {
			return
		}
		config.Decompress = v
	}
}

func BuildConfig(options ...Option) Config {
	config := Config{
		Decompress: false,
		Header:     make(map[string]string),
	}
	for _, option := range options {
		option(&config)
	}
	return config
}

func Decompress(config Config, resp []byte) ([]byte, error) {
	if !config.Decompress {
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
func Get(u string, configs ...Config) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.WithMessagef(err, "call http.NewRequest fail, url:%s", u)
	}
	var config Config
	if len(configs) != 0 {
		config = configs[0]
	}

	for k, v := range config.Header {
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
	bs, err = Decompress(config, bs)
	if err != nil {
		return nil, errors.WithMessage(err, "call Decompress fail")
	}
	return bs, nil
}

// GetToMap 请求资源，以map形式返回结果
func GetToMap(u string, configs ...Config) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	bs, err := Get(u, configs...)
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
func GetToStruct[T any](u string, configs ...Config) (T, error) {
	var data T
	bs, err := Get(u, configs...)
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
func PostForm(u string, form url.Values, configs ...Config) ([]byte, error) {
	var resp *http.Response
	var err error
	if len(configs) == 0 {
		resp, err = http.PostForm(u, form)
	} else {
		req, err := http.NewRequest("POST", u, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		for k, v := range configs[0].Header {
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
	bs, err = Decompress(configs[0], bs)
	if err != nil {
		return nil, errors.WithMessage(err, "call Decompress fail")
	}
	return bs, nil
}

func PostJSONFromStruct(u string, obj interface{}, configs ...Config) ([]byte, error) {
	content, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.WithMessagef(err, "call json.Marshal fail, obj:%+v", obj)
	}
	return PostJSON(u, content, configs...)
}

// PostJSON 以json格式请求
func PostJSON(u string, jsonByte []byte, configs ...Config) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, errors.WithMessage(err, "call http.NewRequest fail")
	}
	req.Header.Set("Content-Type", "application/json")
	var config Config
	if len(configs) > 0 {
		config = configs[0]
	}

	for k, v := range config.Header {
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
	b, err = Decompress(config, b)
	if err != nil {
		return nil, errors.WithMessage(err, "call Decompress fail")
	}
	return b, nil
}

func PostToMapFromStruct(u string, obj interface{}, configs ...Config) (map[string]interface{}, error) {
	content, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.WithMessagef(err, "call json.Marshal fail, obj:%+v", obj)
	}
	return PostToMap(u, content, configs...)
}

func PostToMap(u string, jsonByte []byte, configs ...Config) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	bs, err := PostJSON(u, jsonByte, configs...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return nil, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}

func PostToStructFromStruct[T any](u string, obj interface{}, configs ...Config) (T, error) {
	content, err := json.Marshal(obj)
	if err != nil {
		var data T
		return data, errors.WithMessagef(err, "call json.Marshal fail, obj:%+v", obj)
	}
	return PostToStruct[T](u, content, configs...)
}

func PostToStruct[T any](u string, jsonByte []byte, configs ...Config) (T, error) {
	var data T
	bs, err := PostJSON(u, jsonByte, configs...)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return data, errors.WithMessage(err, "call json.Unmarshal fail")
	}
	return data, nil
}
