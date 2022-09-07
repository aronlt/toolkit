package toolkit

import (
	"net/url"
)

// AddQuery url增加query参数
func AddQuery(rawUrl string, key string, value string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl, err
	}
	values := u.Query()
	values.Add(key, value)
	u.RawQuery = values.Encode()
	return u.String(), nil
}

// QueryUrl 查询url的query参数
func QueryUrl(rawUrl string, key string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil || u.Host == "" {
		return "", err
	}
	values := u.Query()
	value := values.Get(key)
	return value, nil
}
