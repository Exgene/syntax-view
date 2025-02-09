package capture

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
)

type CarbonRequest struct {
	Code string `json:"code"`
}

func Screenshot(filepath string) (image.Image, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	code := string(content)
	requestBody := CarbonRequest{
		Code: code,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall request body: %w", err)
	}

	url := "https://carbonara.solopov.dev/api/cook"

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to get response from carbonara: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("carbonara API returned status code: %d", response.StatusCode)
	}

	img, err := png.Decode(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image from response: %w", err)
	}

	return img, nil
}
