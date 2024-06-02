package utils

import (
	"encoding/json"
	"net/http"
)

type NoOpLogger struct{}

/**
 * FetchJSON fetches JSON from a URL
 * @param url string
 * @param target interface{}
 * @return error
 */
func FetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

/**
 * Printf does nothing
 */

func (NoOpLogger) Printf(string, ...interface{}) {}
