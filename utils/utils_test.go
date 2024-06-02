package utils_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/sunil-bagde/go-todo-app/utils"
)

func TestFetchJSON(t *testing.T) {
	expectedURL := "/"
	expectedData := map[string]interface{}{
		"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		"body":  "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path.Join(expectedURL) {
			t.Errorf("Expected URL: %s, got: %s", expectedURL, r.URL.Path)
		}

		jsonData, _ := json.Marshal(expectedData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}))
	defer mockServer.Close()

	// Call the FetchJSON function
	var result map[string]interface{}
	err := utils.FetchJSON(mockServer.URL, &result)

	if err != nil {
		t.Errorf("FetchJSON returned an error: %v", err)
	}
	// Check if the fetched data matches the expected data
	if !isEqual(result, expectedData) {
		t.Errorf("FetchJSON returned incorrect data. Expected: %v, got: %v", expectedData, result)
	}
}

// Helper function to compare two maps
func isEqual(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for key, valueA := range a {
		valueB, ok := b[key]
		if !ok || valueA != valueB {
			return false
		}
	}

	return true
}
