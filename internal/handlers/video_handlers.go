package handlers

import (
	"io"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/samaita/clip-by-clip/pkg/response"
)

func ReverseVideo(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	command := DoReverseVideo(file.Filename, "outputFile.mp4")
	command.Run()

	return response.OKResponse(c, "OK Ready to Reverse")
}

func DoReverseVideo(inputFile, outputFile string) *exec.Cmd {
	return exec.Command("ffmpeg",
		"-t", "00:00:10",
		"-i", inputFile,
		"-vf", "reverse",
		// "-filter:v", "scale=720:-1", // force scale to 720
		outputFile,
	)
}
