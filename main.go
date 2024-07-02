package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kkdai/youtube/v2"
)

func downloadVideo(url string) error {
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		return fmt.Errorf("error getting video info: %w", err)
	}

	// Select a format for MP4 with audio and video
	var format *youtube.Format
	for _, f := range video.Formats {
		if f.MimeType == "video/mp4" && f.AudioChannels > 0 {
			format = &f
			break
		}
	}

	if format == nil {
		return fmt.Errorf("no suitable format found")
	}

	stream, _, err := client.GetStream(video, format)
	if err != nil {
		return fmt.Errorf("error getting video stream: %w", err)
	}

	file, err := os.Create(video.Title + ".mp4")
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	_, err = file.ReadFrom(stream)
	if err != nil {
		return fmt.Errorf("error downloading video: %w", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ytdl <YouTube-URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	err := downloadVideo(url)
	if err != nil {
		log.Fatalf("YTDL couldn't download the video: %v\n", err)
	}

	fmt.Println("Download completed successfully.")
}
