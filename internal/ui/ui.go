package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	messages []string
}

func InitialModel() model {
	return model{
		messages: []string{"Welcome to TermByte!"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "Ctrl+c","q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string
	for _, message := range m.messages {
		s += message + "\n"
	}
	s += "\nPress 'q' quit."
	return lipgloss.NewStyle().Padding(1).Render(s)
}