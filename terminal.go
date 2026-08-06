package main

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type mode string

const (
	ModeType   mode = "Typing"
	ModeSelect mode = "Select"
)

type Input struct {
	path     string
	filename string
	ti       textinput.Model
}

func NewInput() *Input {
	ti := textinput.New()
	ti.Placeholder = "File path"
	return &Input{ti: ti}
}

type model struct {
	Input       *Input
	CurrentMode mode
}

func initialModel() model {
	return model{Input: NewInput()}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
	}
}
