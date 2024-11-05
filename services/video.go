package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"yourate/models"

	"github.com/joho/godotenv"
)

type snippetJSON struct {
	Title        string `json:"title"`
	ChannelTitle string `json:"channelTitle"`
}

type videoJSON struct {
	Id      string      `json:"id"`
	Snippet snippetJSON `json:"snippet"`
}

type (
	VideoService interface {
		GetVideoDetails(string) (*models.Video, error)
		getVideoID(string) (string, error)
	}

	videoService struct {
		//store
	}
)

func (vs *videoService) GetVideoDetails(videoURL string) (*models.Video, error) {
	videoID, err := vs.getVideoID(videoURL)

	video := &models.Video{
		Url: videoURL,
	}

	if err != nil {
		return video, fmt.Errorf("error getting video ID: %w", err)
	}

	err = godotenv.Load()

	if err != nil {
		return video, fmt.Errorf("Error loading ENV")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")

	if apiKey == "" {
		return video, fmt.Errorf("Google API key not found in environment")
	}

	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet,statistics&id=%s&key=%s", videoID, apiKey)

	resp, err := http.Get(apiURL)

	if err != nil {
		return video, fmt.Errorf("error making request to YouTube API: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return video, fmt.Errorf("error reading response body: %w", err)
	}

	var response struct {
		Items []videoJSON `json:"items"`
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Error:", err)
		return video, fmt.Errorf("error converting JSON: %w", err)
	}

	video.Id = response.Items[0].Id
	video.Title = response.Items[0].Snippet.Title
	video.ChannelTitle = response.Items[0].Snippet.ChannelTitle

	return video, nil
}

func (vs *videoService) getVideoID(videoURL string) (string, error) {
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
