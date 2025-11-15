package main

// An example demonstrating advanced progress bar styling techniques.
//
// This example shows sophisticated gradient effects using the current API.
// With the upcoming ColorFunc feature (PR #838), these effects could be
// achieved more elegantly with dynamic color functions.
//
// See: https://github.com/charmbracelet/bubbles/pull/838

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time

type model struct {
	// Four different progress bars showcasing various gradient effects
	warmGradient    progress.Model // Red to yellow
	coolGradient    progress.Model // Blue to cyan
	scaledGradient  progress.Model // Gradient scales with fill
	vibrantGradient progress.Model // Purple to pink (default)

	percent float64
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width := msg.Width - padding*2 - 4
		if width > maxWidth {
			width = maxWidth
		}
		m.warmGradient.Width = width
		m.coolGradient.Width = width
		m.scaledGradient.Width = width
		m.vibrantGradient.Width = width
		return m, nil

	case tea.KeyMsg:
		return m, tea.Quit

	case tickMsg:
		if m.percent >= 1.0 {
			return m, tea.Quit
		}

		// Increment progress
		m.percent += 0.01

		// Update all progress bars
		cmd1 := m.warmGradient.SetPercent(m.percent)
		cmd2 := m.coolGradient.SetPercent(m.percent)
		cmd3 := m.scaledGradient.SetPercent(m.percent)
		cmd4 := m.vibrantGradient.SetPercent(m.percent)

		return m, tea.Batch(tickCmd(), cmd1, cmd2, cmd3, cmd4)

	case progress.FrameMsg:
		progressModel, cmd := m.warmGradient.Update(msg)
		m.warmGradient = progressModel.(progress.Model)

		progressModel, cmd2 := m.coolGradient.Update(msg)
		m.coolGradient = progressModel.(progress.Model)

		progressModel, cmd3 := m.scaledGradient.Update(msg)
		m.scaledGradient = progressModel.(progress.Model)

		progressModel, cmd4 := m.vibrantGradient.Update(msg)
		m.vibrantGradient = progressModel.(progress.Model)

		return m, tea.Batch(cmd, cmd2, cmd3, cmd4)

	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)

	s := "\n" +
		pad + "Advanced Progress Bar Gradients\n\n" +
		pad + "1. Warm Gradient (Red → Yellow)\n" +
		pad + m.warmGradient.View() + "\n\n" +
		pad + "2. Cool Gradient (Blue → Cyan)\n" +
		pad + m.coolGradient.View() + "\n\n" +
		pad + "3. Scaled Gradient (fits fill width)\n" +
		pad + m.scaledGradient.View() + "\n\n" +
		pad + "4. Vibrant Gradient (Purple → Pink)\n" +
		pad + m.vibrantGradient.View() + "\n\n" +
		pad + helpStyle("Press any key to quit") + "\n\n" +
		pad + lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Render(
		"Note: With ColorFunc (PR #838), you'll be able to create\n"+
			pad+"dynamic color changes, stripes, and custom gradients!",
	)

	return s
}

// tickCmd generates a tick message every 80ms
func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*80, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	// Create progress bars with different gradient effects

	// 1. Warm gradient (red to yellow) - great for error/warning states
	warmGradient := progress.New(
		progress.WithGradient("#FF0000", "#FFFF00"),
		progress.WithWidth(40),
	)

	// 2. Cool gradient (blue to cyan) - great for info/processing states
	coolGradient := progress.New(
		progress.WithGradient("#0000FF", "#00FFFF"),
		progress.WithWidth(40),
	)

	// 3. Scaled gradient - gradient scales to always fit the filled portion
	scaledGradient := progress.New(
		progress.WithScaledGradient("#00FF00", "#00AA00"),
		progress.WithWidth(40),
	)

	// 4. Vibrant gradient using the default colors
	vibrantGradient := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
	)

	m := model{
		warmGradient:    warmGradient,
		coolGradient:    coolGradient,
		scaledGradient:  scaledGradient,
		vibrantGradient: vibrantGradient,
		percent:         0.0,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
