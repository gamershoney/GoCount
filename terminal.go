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
	ti.Focus()
	ti.SetWidth(20)
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
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.Input.ti, cmd = m.Input.ti.Update(msg)

	return m, cmd
}

func (m model) View() tea.View {
	str := m.Input.ti.View()

	v := tea.NewView(str)
	return v
}
