package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "SLACK BOT TOKEN HERE")
	os.Setenv("CHANNEL_ID", "CHANNEL ID HERE")

	//api connection with slack
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	//channel id to integrate with slack
	channelArray := []string{os.Getenv("CHANNEL_ID")}

	//file uploading to slack bot channel
	fileArray := []string{"testfile.jpg"}

	for i := 0; i < len(fileArray); i++ {

		//slack api function to upload files
		params := slack.FileUploadParameters{
			Channels: channelArray,
			File:     fileArray[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Nombre: %s, URL: %s\n", file.Name, file.URL)
	}

}
