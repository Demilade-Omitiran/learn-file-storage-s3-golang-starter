package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math"
	"os/exec"
)

type dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type commandOutput struct {
	Streams []dimension `json:"streams"`
}

func calculateAspectRatio(width, height int) string {
	ratio := float64(width) / float64(height)
	if math.Abs(ratio-(9.0/16.0)) < 0.01 {
		return "9:16"
	} else if math.Abs(ratio-(16.0/9.0)) < 0.01 {
		return "16:9"
	} else {
		return "other"
	}
}

func getVideoAspectRatio(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)
	var buffer bytes.Buffer
	cmd.Stdout = &buffer

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	outputStreams := commandOutput{}

	if err = json.Unmarshal(buffer.Bytes(), &outputStreams); err != nil {
		return "", err
	}

	width := outputStreams.Streams[0].Width
	height := outputStreams.Streams[0].Height

	return calculateAspectRatio(width, height), nil
}
