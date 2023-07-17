package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	ctx := context.Background()

	// Set up Google Cloud credentials
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "path/to/service/account/key.json")

	// Create a new SpeechClient
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Read audio file
	audioFile := "path/to/audio/file.flac"
	data, err := os.ReadFile(audioFile)
	if err != nil {
		log.Fatalf("Failed to read audio file: %v", err)
	}

	// Create a new RecognitionConfig
	config := &speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_FLAC,
		SampleRateHertz: 16000,   // Update sample rate if needed
		LanguageCode:    "en-US", // Update language code if needed
	}

	// Create a new RecognitionAudio
	audio := &speechpb.RecognitionAudio{
		AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
	}

	// Perform the speech recognition
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: config,
		Audio:  audio,
	})
	if err != nil {
		log.Fatalf("Failed to recognize speech: %v", err)
	}

	// Process the response
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("Transcript: %s\n", alt.Transcript)
		}
	}
}
