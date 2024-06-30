package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tregalloway/termbyte/internal/ui"
)

func main() {
	p := tea.NewProgram(ui.InitialModel())
	if err := p.Star(); err != nil {
		log.Fatalf("Error starting program: %v", err)
	}
}