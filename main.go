package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
)

const (
	WIDTH  = 80
	HEIGHT = 20
)

type coords struct {
	x int
	y int
}

type model struct {
	ship      int
	aliens    []coords
	direction bool
	width     int
	height    int
}

func initialModel() model {
	const (
		startingAliens = 30
		aliensPerLine  = 10
		spaceBetween   = 2
	)
	aliens := []coords{}
	sx := WIDTH/2 - (aliensPerLine*spaceBetween)/2
	x := sx
	y := 0
	for i := 1; i <= startingAliens; i++ {
		aliens = append(aliens, coords{x: x, y: y})
		x += spaceBetween
		if i%10 == 0 {
			y++
			x = sx
		}
	}
	return model{
		ship:   40,
		aliens: aliens,
		width:  WIDTH,
		height: HEIGHT,
	}
}

func (m model) Init() tea.Cmd {
	return tickEvery()
}

type TickMsg time.Time

func tickEvery() tea.Cmd {
	return tea.Every(time.Second/10, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
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
	case TickMsg:
		aliens := []coords{}
		atEdge := false
		for _, alien := range m.aliens {
			if alien.x == m.width || alien.x == 0 {
				atEdge = true
				m.direction = !m.direction
				break
			}
		}
		for _, alien := range m.aliens {
			if m.direction {
				alien.x--
			} else {
				alien.x++
			}
			if atEdge {
				alien.y++
			}
			aliens = append(aliens, alien)
		}
		m.aliens = aliens
		return m, tickEvery()
	}
	return m, nil
}
func (m model) View() tea.View {
	title := "SPACE INVADERS"
	padding := (m.width - len(title)) / 2
	s := "x" + strings.Repeat(" ", padding-2) + title + strings.Repeat(" ", padding) + "x" + "\n\n"
	lines := []string{}

	// reserve last line for ship
	for range m.height - 1 {
		lines = append(lines, strings.Repeat(" ", m.width)+"\n")
	}

	for _, alien := range m.aliens {
		line := lines[alien.y]
		line = line[:alien.x] + "m" + line[alien.x+1:]
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
	s += strings.Repeat(" ", lp) + "w" + strings.Repeat(" ", rp)
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
