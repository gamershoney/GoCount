package main

import (
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type mode string

const (
	ModeInputFilePath mode = "Typing"
	ModeSelect        mode = "Select"
	ModeOpenSheet     mode = "Opening"
)

type Input struct {
	path     string
	filename string
	ti       textinput.Model
	loading  bool
	wb       *WorkBook
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
	Spinner     spinner.Model
}

func initialModel() model {
	m := model{Input: NewInput(), CurrentMode: ModeInputFilePath}
	m.resetSpinner()
	return m
}

func (m *model) resetSpinner() {
	m.Spinner = spinner.New()
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

func (m model) ClearError() model {
	m.Error = nil
	return m
}

func (m model) ReadSheets() *WorkBook {
	if m.Input.path == "" {
		return nil
	}
	wb := GetSheets(m.Input.path)
	return wb
}

func (m model) RenderInput() string {
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

func (m model) HandleInputMode(msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m = m.FindFile(m.Input.ti.Value())
			if m.Error != nil {
				return m, nil
			}
			m.CurrentMode = ModeOpenSheet
			return m.HandleOpenMode(msg)
		}
	}

	m.Input.ti, cmd = m.Input.ti.Update(msg)

	return m, cmd
}

type loadFinishedMsg struct {
	result     *WorkBook
	transition mode
}

func (m model) RenderLoading() string {
	return m.Spinner.View()
}

func (m model) HandleOpenMode(msg tea.Msg) (model, tea.Cmd) {
	m.Input.loading = true

	openLoad := func() tea.Msg {
		result := m.ReadSheets()
		return loadFinishedMsg{result: result}
	}
	return m, tea.Batch(m.Spinner.Tick, openLoad)
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
	switch m.CurrentMode {
	case ModeInputFilePath:
		m, cmd = m.HandleInputMode(msg)
	case ModeOpenSheet:
		switch msg := msg.(type) {
		case loadFinishedMsg:
			m.Input.loading = false
			m.Input.wb = msg.result
			return m, nil
		case spinner.TickMsg:
			if !m.Input.loading {
				return m, nil
			}
			m.Spinner, cmd = m.Spinner.Update(msg)
			return m, cmd

		}
	}
	return m, cmd
}

func (m model) View() tea.View {
	m = m.ClearError()
	switch m.CurrentMode {
	case ModeInputFilePath:
		input := m.RenderInput()
		v := tea.NewView(lipgloss.JoinVertical(lipgloss.Left, HeaderText("Input the file location"), input))
		return v
	case ModeOpenSheet:
		spinner := m.RenderLoading()
		v := tea.NewView(lipgloss.JoinVertical(lipgloss.Left, HeaderText("Hang on a second"), spinner))
		return v

	default:
		return tea.NewView("ERROR!")
	}
}
