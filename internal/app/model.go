package app

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Width  int
	Height int
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}
