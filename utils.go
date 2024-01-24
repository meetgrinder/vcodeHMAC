package vcodeHMAC

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net"
	"net/url"
	"time"
)

func getCurrentTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}

func generateNonce() string {
	token := make([]byte, 16)
	rand.Read(token)
	return hex.EncodeToString(token)
}

func getHost(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	host, _, err := net.SplitHostPort(u.Host)
	return host, nil
}

func getPathParams(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	if len(u.RawQuery) > 0 {
		return fmt.Sprintf("%s?%s", u.Path, u.RawQuery), nil
	}
	return string(u.Path), nil
}

func getScheme(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	return u.Scheme, nil
}
