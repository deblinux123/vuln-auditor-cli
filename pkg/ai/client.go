package ai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/log"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	System string `json:"system,omitempty"`
	Stream bool   `json:"stream"`
}

type OllamaStreamResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func StreamOllamaResponse(model, systemPrompt, userPrompt string) (string, error) {
	// This function will handle the streaming response from Ollama.

	ollamaURL := "http://localhost:11434/api/generate"

	requestBody, err := json.Marshal(OllamaRequest{
		Model:  model,
		System: systemPrompt,
		Prompt: userPrompt,
		Stream: true,
	})

	if err != nil {
		return "", fmt.Errorf("Failed to marshal request body: %w", err)
	}

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return "", fmt.Errorf("Failed to connect to Ollama API: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama API returned non-OK status: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)

	var fullResponse strings.Builder

	fmt.Println("AI Response: ")
	fmt.Println(strings.Repeat("-", 50))

	for scanner.Scan() {
		line := scanner.Bytes()

		var streamResp OllamaStreamResponse

		if err := json.Unmarshal(line, &streamResp); err != nil {
			log.Errorf("Failed to unmarshal stream response: %v", err)
			continue
		}

		fmt.Print(streamResp.Response)

		os.Stdout.Sync()

		fullResponse.WriteString(streamResp.Response)

		if streamResp.Done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Error reading stream response: %w", err)
	}

	fmt.Println("\n" + strings.Repeat("-", 50))

	return fullResponse.String(), nil
}

func ReadMarkdown(text string) {
	// This function can be used to read and display the markdown response in a more user-friendly way.

	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(100),
	)

	if err != nil {
		fmt.Println(text)
		return
	}

	out, err := r.Render(text)

	if err != nil {
		fmt.Println(text)
		return
	}

	fmt.Println(out)
}
