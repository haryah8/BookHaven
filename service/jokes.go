package service

import (
	"BookHaven/config"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// FetchRandomJoke fetches a joke from the joke API
func FetchRandomJoke() (string, error) {
	url := "https://api.api-ninjas.com/v1/jokes"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Api-Key", config.NINJA_API_KEY)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch joke, status code: " + resp.Status)
	}

	var jokeResponses []struct {
		Joke string `json:"joke"`
	}

	if err := json.Unmarshal(body, &jokeResponses); err != nil {
		return "", err
	}

	if len(jokeResponses) == 0 {
		return "", errors.New("no jokes found")
	}

	return jokeResponses[0].Joke, nil
}
