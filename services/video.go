package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"yourate/domain"

	"github.com/joho/godotenv"
)

type Snippet struct {
	Title string `json:"title"`
}

type Video struct {
	Id      string  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

func GetVideoDetails(videoURL string) (*domain.Video, error) {
	videoID, err := getVideoID(videoURL)

	if err != nil {
		return nil, fmt.Errorf("error getting video ID: %w", err)
	}

	err = godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("Error loading ENV")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("YouTube API key not found in environment")
	}

	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet,statistics&id=%s&key=%s", videoID, apiKey)

	resp, err := http.Get(apiURL)

	if err != nil {
		return nil, fmt.Errorf("error making request to YouTube API: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response struct {
		Items []Video `json:"items"`
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Error:", err)
		return nil, fmt.Errorf("error converting JSON: %w", err)
	}

	domainVideo := domain.Video{
		Id:    response.Items[0].Id,
		Title: response.Items[0].Snippet.Title,
	}

	return &domainVideo, nil
}

func getVideoID(videoURL string) (string, error) {
	u, err := url.Parse(videoURL)

	if err != nil {
		return "", err
	}

	query := u.Query()

	videoID := query.Get("v")

	if videoID == "" {
		return "", fmt.Errorf("video ID not found in URL")
	}
	return videoID, nil
}
