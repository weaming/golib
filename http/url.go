package http

import (
	"net/url"
)

//URLEncoded encodes a string to be used in a query part of a URL
func URLEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
