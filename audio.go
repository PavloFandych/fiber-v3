package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"io"
	"log"
	"net/http"
)

func audioHandler(c fiber.Ctx, client *http.Client, text *string, key *string) error {
	log.Printf("Input: %s", *text)
	requestBody, requestBodyErr := json.Marshal(map[string]interface{}{
		"model": "tts-1",
		"input": *text,
		"voice": "alloy",
	})
	if requestBodyErr != nil {
		panic(requestBodyErr)
	}

	request, requestErr := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/audio/speech", bytes.NewBuffer(requestBody))
	if requestErr != nil {
		panic(requestErr)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", *key)

	response, responseErr := client.Do(request)
	if responseErr != nil {
		panic(responseErr)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	c.Set("Content-Type", "audio/mpeg")
	c.Set("Content-Disposition", `attachment; filename="generated_audio.mp3"`)

	_, copyErr := io.Copy(c, response.Body)
	if copyErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Downloading audio failure")
	}

	return nil
}
