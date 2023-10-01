package web

import (
	"io"
	"net/http"
	"strings"
	"time"
)

func GetCount(url, word string) (int, error) {
	resp, err := getResponse(url)
	if err != nil {
		return 0, err
	}
	return findWord(word, resp), nil
}

func getResponse(url string) (string, error) {
	client := http.Client{
		Timeout: 6 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func findWord(word string, resp string) int {
	return strings.Count(resp, word)
}
