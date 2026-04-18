package utils

import (
	"github.com/manifoldco/promptui"
)

var defaultTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}:",
	Active:   "▶ {{ . | cyan | bold }}",
	Inactive: "  {{ . | faint }}",
	Selected: "✓ {{ . | green }}",
}

func SelectPrompt(label string, items []string) (int, string, error) {
	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		Templates: defaultTemplates,
		Size:      6,
	}
	return prompt.Run()
}