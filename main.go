package main

import (
	"fmt"
	"os"

	"github.com/Aslanbayli/math-gym/tui"
	tea "github.com/charmbracelet/bubbletea"
)

var p *tea.Program

type sessionState int

const (
	operationsView   sessionState = 1
	digitsView       sessionState = 2
	calculationsView sessionState = 3
)

type MainModel struct {
	state       sessionState
	operations  *tui.Operations
	digits      *tea.Model
	calculation *tea.Model
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch kepress := msg.String(); kepress {
		case "o":
			m.state = operationsView
		}
	}

	return m, nil
}

func (m MainModel) View() string {
	switch m.state {
	case operationsView:
		m.operations = tui.Run()
		return m.operations.View()
	default:
		return ""
	}
}

func main() {
	if _, err := tea.NewProgram(MainModel{}).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
