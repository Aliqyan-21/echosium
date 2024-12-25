// package jamendo is the package that integrates the jamendo api in echosium
package jamendo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Track struct represents a Jamendo track
type Track struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	TrackUrl      string `json:"song"`
	Artist        string `json:"artist"`
	Album         string `json:"album"`
	Image         string `json:"image"`
	AudioFormat   string `json:"audioformat"`
	AudioDownload bool   `json:"audiodownload"`
}

func getTracks(clientID, mood string) ([]Track, error) {
	url := fmt.Sprintf("https://api.jamendo.com/v3.0/tracks/?client_id=%s&fuzzytags=%s&format=jsonpretty", clientID, mood)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	body, _ := io.ReadAll(resp.Body)
	// fmt.Printf("raw: %s", string(body))

	resp.Body = io.NopCloser(bytes.NewReader(body))

	var result struct {
		Results []Track `json:"results"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Results, nil
}
