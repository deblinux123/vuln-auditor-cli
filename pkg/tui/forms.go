package tui

import "github.com/charmbracelet/huh"

type FormResult struct {
	ModelName string
	FilePath  string
}

func RunInitialForm() (FormResult, error) {
	var result FormResult

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("🤖 Select the AI model to use (ollama)").
				Options(
					huh.NewOption("Qwen 2.5 Coder (Recommended)", "qwen2.5-coder"),
					huh.NewOption("Gemma 2", "gemma2"),
					huh.NewOption("Gemma 3", "gemma3"),
					huh.NewOption("Llama 3", "llama3"),
				).Value(&result.ModelName),

			huh.NewInput().
				Title("📂 Enter the path to the security (sqlmap) tool log file").
				Placeholder("./logs/sqlmap_output.txt").
				Value(&result.FilePath),
		),
	)

	err := form.Run()

	return result, err
}
