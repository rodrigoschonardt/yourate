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
	"github.com/labstack/echo/v4"
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
		logger echo.Logger
		//store
	}
)

func (vs *videoService) GetVideoDetails(videoURL string) (*models.Video, error) {
	video := &models.Video{
		Url: videoURL,
	}

	videoID, err := vs.getVideoID(videoURL)

	if err != nil {
		return video, err
	}

	// TODO: separate this logic
	err = godotenv.Load()

	if err != nil {
		vs.logger.Error(err)
		return video, fmt.Errorf("Internal Error, contact the support!")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")

	if apiKey == "" {
		vs.logger.Error(err)
		return video, fmt.Errorf("Internal Error, contact the support!")
	}

	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet,statistics&id=%s&key=%s", videoID, apiKey)

	resp, err := http.Get(apiURL)

	if err != nil {
		vs.logger.Error(err)
		return video, fmt.Errorf("Internal Error, contact the support!")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		vs.logger.Error(err)
		return video, fmt.Errorf("Internal Error, contact the support!")
	}

	var response struct {
		Items []videoJSON `json:"items"`
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		vs.logger.Error(err)
		return video, fmt.Errorf("Internal Error, contact the support!")
	}

	if len(response.Items) != 1 {
		return video, fmt.Errorf("Video not found!")
	}

	video.Id = response.Items[0].Id
	video.Title = response.Items[0].Snippet.Title
	video.ChannelTitle = response.Items[0].Snippet.ChannelTitle

	return video, nil
}

func (vs *videoService) getVideoID(videoURL string) (string, error) {
	u, err := url.Parse(videoURL)

	if err != nil {
		return "", fmt.Errorf("URL is not valid!")
	}

	query := u.Query()

	videoID := query.Get("v")

	if videoID == "" {
		return "", fmt.Errorf("Video ID not found in the URL!")
	}

	return videoID, nil
}
