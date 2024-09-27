package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"io"
	"log"
	"net/http"
)

func imageHandler(c fiber.Ctx, client *http.Client, description *string, key *string) error {
	log.Printf("Input: %s", *description)
	requestBody, requestBodyErr := json.Marshal(map[string]interface{}{
		"prompt": description,
		"n":      1,
		"size":   "1024x1024",
	})
	if requestBodyErr != nil {
		panic(requestBodyErr)
	}

	request, requestErr := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/images/generations", bytes.NewBuffer(requestBody))
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

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		panic(readErr)
	}

	return c.SendString(string(body))
}
