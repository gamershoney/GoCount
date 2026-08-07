package main

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type mode string

const (
	ModeInputFilePath mode = "Typing"
	ModeSelect        mode = "Select"
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
	ti.SetWidth(40)
	styles := ti.Styles()
	styles.Focused.Text = styles.Focused.Text.
		Foreground(lipgloss.BrightGreen)
	ti.SetStyles(styles)

	return &Input{ti: ti}
}

type model struct {
	Input       *Input
	CurrentMode mode
	Error       error
}

func initialModel() model {
	return model{Input: NewInput(), CurrentMode: ModeInputFilePath}
}

func (m model) FindFile(p string) model {
	input, err := FindSheet(p)
	if err != nil {
		m.Error = err
		return m
	}
	m.Input = &Input{path: input.path, filename: input.filename, ti: m.Input.ti}
	return m
}

func (m model) Render() string {
	input := m.Input.ti.View()
	input = styleInput.Render(input)
	if m.Error != nil {
		strerr := styleError.Render(m.Error.Error())
		input = lipgloss.JoinVertical(lipgloss.Left, input, strerr)
		return input
	}
	return input
}

func HeaderText(s string) string {
	return styleHeader.Render(s)
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m = m.FindFile(m.Input.ti.Value())
			if m.Error == nil {
				return m, cmd
			}
		}
	}

	m.Input.ti, cmd = m.Input.ti.Update(msg)

	return m, cmd
}

func (m model) View() tea.View {
	switch m.CurrentMode {
	case ModeInputFilePath:
		input := m.Render()
		v := tea.NewView(lipgloss.JoinVertical(lipgloss.Left, HeaderText("Input the file location"), input))
		return v
	default:
		return tea.NewView("ERROR!")
	}
}
