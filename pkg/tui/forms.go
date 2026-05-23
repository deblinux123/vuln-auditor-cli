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
					huh.NewOption("Deepseek Coder (Recommended)", "deepseek-coder:6.7b"),
					huh.NewOption("Code Gemma", "codegemma:7b"),
					huh.NewOption("Gemma 4", "gemma4:e4b"),
					huh.NewOption("Gemma 3", "gemma3:4b"),
					huh.NewOption("Qwen 3.6", "qwen3.6:35b"),
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
