package hls

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateHLS(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Error uploading file: %s", err.Error())
		return
	}

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "video/") {
		c.String(http.StatusBadRequest, "Error: only video files are allowed")
		return
	}

	err = c.SaveUploadedFile(file, "./videos/"+file.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error saving file: %s", err.Error())
		return
	}
	inputFile, err := filepath.Abs("./videos/" + file.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting file path: %s", err.Error())
		return
	}

	c.String(http.StatusOK, "Folder path: %s", inputFile)

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

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"Message": "Okay",
	})
}
