package main

import tea "charm.land/bubbletea/v2"

type mode string

const (
	ModeType   mode = "Typing"
	ModeSelect mode = "Select"
)

type Input struct {
	path     string
	filename string
}

type model struct {
	Input       *Input
	CurrentMode mode
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
	}
}
