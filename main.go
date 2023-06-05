package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/videoStreaming/hls"
)

func main() {
	r := gin.Default()
	r.POST("/select/file", hls.CreateHLS)
	r.Static("/stream", "/home/afthab/Desktop/videoStreaming/outputFile/")
	fmt.Println("Port running at :9998/stream/playlist.m3u8")
	r.Run(":9998")
}

//https://hlsjs-dev.video-dev.org/demo/
