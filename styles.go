package main

import "charm.land/lipgloss/v2"

var styleHeader = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.BrightBlue).
	MarginTop(5).
	MarginBottom(2)

var styleInput = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder())

var styleError = lipgloss.NewStyle().
	Foreground(lipgloss.BrightRed).
	Bold(true).
	Underline(true).
	UnderlineColor(lipgloss.White)
