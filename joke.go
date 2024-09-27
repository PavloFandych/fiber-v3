package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"io"
	"net/http"
)

func jokeHandler(c fiber.Ctx, client *http.Client, key *string) error {
	requestBody, requestBodyErr := json.Marshal(map[string]interface{}{
		"model": "gpt-4",
		"messages": []map[string]string{
			{"role": "user", "content": "Розкажи новий жарт!"},
		},
	})
	if requestBodyErr != nil {
		panic(requestBodyErr)
	}

	request, requestErr := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
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
