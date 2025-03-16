package curl

import (
	"io"
	"net/http"
)

func Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}