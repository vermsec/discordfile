package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	version := flag.Bool("version", false, "current version")
	help := flag.Bool("help", false, "usage info")
	webhookURL := flag.String("webhook", "whook", "Webhook URL to send the data to")
	fileName := flag.String("file", "nofile", "Path to file which is to be sent")
	flag.Parse()

	if *version {
		fmt.Println("0.1.1")
		os.Exit(0)
	}
	if *help {
		fmt.Printf("Usage:\n▶ discordfile -webhook https://discord.com/api/webhooks/213123 -file /path/to/file \n")
		os.Exit(0)
	}

	bodyWriter, bodyBuffer := createFile(*fileName)

	request, err := http.NewRequest("POST", *webhookURL, bodyBuffer)
	handleError(err, true)

	request.Header.Add("Content-Type", bodyWriter.FormDataContentType())

	response, err := httpClient.Do(request)
	handleError(err, true)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		fmt.Println("File successfully Uploaded")
		return
	}

	fmt.Println("File Upload Failed")
}

func createFile(fileName string) (*multipart.Writer, *bytes.Buffer) {
	file, err := os.Open(fileName)
	handleError(err, true)
	defer file.Close()

	body := &bytes.Buffer{}
	bodywriter := multipart.NewWriter(body)

	part, err := bodywriter.CreateFormFile("multipart/form-data", filepath.Base(file.Name()))
	handleError(err, true)

	io.Copy(part, file)
	bodywriter.Close()

	return bodywriter, body
}

func handleError(err error, serious bool) {
	if err == nil {
		return
	}

	if serious {
		fmt.Printf("Usage:\n▶ discordfile -webhook https://discord.com/api/webhooks/213123 -file /path/to/file \n")
		os.Exit(0)
	}

	fmt.Printf("[Error]: %s\n", err.Error())
}

var (
	httpClient = &http.Client{}
)
