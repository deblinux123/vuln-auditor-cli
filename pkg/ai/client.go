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

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type OllamaChatResponse struct {
	Message Message `json:"message"`
	Done    bool    `json:"done"`
}

func ChatWithOllama(model string, conversation []Message) (string, error) {
	ollamaURL := "http://localhost:11434/api/chat"

	requestBody, err := json.Marshal(OllamaChatRequest{
		Model:    model,
		Messages: conversation,
		Stream:   true,
	})
	if err != nil {
		return "", fmt.Errorf("error marshalling chat request: %w", err)
	}

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama returned non-200 status: %d", resp.StatusCode)
	}

	scanner := bufio.NewScanner(resp.Body)
	var fullResponse strings.Builder

	for scanner.Scan() {
		line := scanner.Bytes()
		var streamResp OllamaChatResponse

		if err := json.Unmarshal(line, &streamResp); err != nil {
			log.Error("Error parsing chat stream line", "error", err)
			continue
		}

		fmt.Print(streamResp.Message.Content)
		os.Stdout.Sync()

		fullResponse.WriteString(streamResp.Message.Content)

		if streamResp.Done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading stream data: %w", err)
	}

	return fullResponse.String(), nil
}

func RenderMarkdown(text string) {
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
	fmt.Print(out)
}
