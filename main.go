package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	inputFile := "/home/afthab/Desktop/videoStreaming/videoFile/football.mp4"
	outputDir := "/home/afthab/Desktop/videoStreaming/outputFile"
	segmentTime := "10" // Segment duration in seconds

	cmd := exec.Command("ffmpeg",
		"-i", inputFile,
		"-c:v", "copy",
		"-c:a", "aac",
		"-hls_time", segmentTime,
		"-hls_list_size", "0",
		"-hls_segment_filename", outputDir+"/%03d.ts",
		outputDir+"/playlist.m3u8",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
