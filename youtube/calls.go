package youtube

import (
	"context"
	"fmt"

	"github.com/gabrielebnc/Watchlist-CLI/utils"
	"github.com/senseyeio/duration"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type VideoInfo struct {
	title       string
	channelName string
	duration    string
}

const (
	maxResults = 1
)

func (v *VideoInfo) Fstring() string {
	return fmt.Sprintf("Title: %v\nChannel:%v\nDuration:%v", v.title, v.channelName, v.duration)
}

func formatDuration(d duration.Duration) string {

	if d.Y > 0 || d.M > 0 || d.W > 0 || d.D > 0 {
		return "A lot of time"
	}

	return fmt.Sprintf("%v:%v:%v", d.TH, d.TM, d.TS)
}

func generateYoutubeService(apiKey string) *youtube.Service {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		utils.PrintfSTDERR("Error creating Youtube Service: %v", err)
	}
	return youtubeService
}

func getVideoInfoByVideoId(youtubeService *youtube.Service, videoId string) VideoInfo {

	call := youtubeService.Videos.
		List([]string{"snippet", "contentDetails"}).
		Id(videoId).
		MaxResults(maxResults)

	response, err := call.Do()
	if err != nil {
		utils.PrintfSTDERR("Error on executing call: %v", err)
	}

	videoItem := *(response.Items[0])

	var videoDuration string

	d, err := duration.ParseISO8601(videoItem.ContentDetails.Duration)
	if err != nil {
		utils.PrintfSTDERR("Error parsing video duration")
		videoDuration = "n/a"
	} else {
		videoDuration = formatDuration(d)
	}

	return VideoInfo{
		title:       videoItem.Snippet.Title,
		channelName: videoItem.Snippet.ChannelTitle,
		duration:    videoDuration,
	}
}

func SearchVideoById(videoId string, apiKey string) VideoInfo {

	youtubeService := generateYoutubeService(apiKey)

	return getVideoInfoByVideoId(youtubeService, videoId)
}

func SearchVideosInfoByIds(youtubeService *youtube.Service, videoId string) VideoInfo {
	call := youtubeService.Videos.
		List([]string{"snippet", "contentDetails"}).
		Id(videoId).
		MaxResults(maxResults)

	response, err := call.Do()
	if err != nil {
		utils.PrintfSTDERR("Error on executing call: %v", err)
	}

	videoItem := *(response.Items[0])

	var videoDuration string

	d, err := duration.ParseISO8601(videoItem.ContentDetails.Duration)
	if err != nil {
		utils.PrintfSTDERR("Error parsing video duration")
		videoDuration = "n/a"
	} else {
		videoDuration = formatDuration(d)
	}

	return VideoInfo{
		title:       videoItem.Snippet.Title,
		channelName: videoItem.Snippet.ChannelTitle,
		duration:    videoDuration,
	}
}

func SearchVideosByIds(videoIds []string, apiKey string) *[]VideoInfo {
	youtubeService := generateYoutubeService(apiKey)

	vInfos := make([]VideoInfo, 0)

	for _, vId := range videoIds {
		vInfos = append(vInfos, getVideoInfoByVideoId(youtubeService, vId))
	}

	return &vInfos
}
