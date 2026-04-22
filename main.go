package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type coords struct {
	x int
	y int
}

type model struct {
	ship   int
	aliens []coords
	width  int
	height int
}

func initialModel() model {
	return model{
		ship:   40,
		aliens: []coords{{x: 8, y: 10}},
		width:  80,
		height: 20,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left":
			if m.ship > 0 {
				m.ship--
			}
		case "right":
			if m.ship < m.width {
				m.ship++
			}
		case "space":
		}
	}
	return m, nil
}
func (m model) View() tea.View {
	title := "SPACE INVADERS"
	padding := (m.width - len(title)) / 2
	s := strings.Repeat(" ", padding) + title + strings.Repeat(" ", padding) + "\n\n"
	lines := []string{}

	// reserve last line for ship
	for range m.height - 1 {
		lines = append(lines, strings.Repeat(" ", m.width)+"\n")
	}

	for _, alien := range m.aliens {
		line := lines[alien.y]
		line = line[:alien.x] + "%" + line[alien.x+1:]
		lines[alien.y] = line
	}
	for _, line := range lines {
		s += line
	}
	lp := m.ship - 1
	if lp < 0 {
		lp = 0
	}
	rp := m.width - lp + 1
	if rp > m.width {
		rp = 0
	}
	s += strings.Repeat(" ", lp) + "q" + strings.Repeat(" ", rp)
	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
