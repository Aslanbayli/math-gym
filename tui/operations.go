package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("#c92d0a"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#068504"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	op := []string{"+", "-", "*", "/"}
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("(%s) %s", op[index], i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type Operations struct {
	list     list.Model
	choice   string
	quitting bool
}

func (op Operations) Init() tea.Cmd {
	return nil
}

func (op Operations) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		op.list.SetWidth(msg.Width)
		return op, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q":
			op.quitting = true
			return op, tea.Quit

		case "enter":
			i, ok := op.list.SelectedItem().(item)
			if ok {
				op.choice = string(i)
			}
			return op, tea.Quit
		}
	}

	var cmd tea.Cmd
	op.list, cmd = op.list.Update(msg)
	return op, cmd
}

func (op Operations) View() string {
	if op.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("Operation: %s", op.choice))
	}
	if op.quitting {
		return quitTextStyle.Render("----- ❤️ Have a good day ❤️ -----")
	}
	return "\n" + op.list.View()
}

func Run() *Operations {
	items := []list.Item{
		item("addition"),
		item("substraction"),
		item("multiplication"),
		item("division"),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Please choose the mathematical operation to perform (add | sub | mul | div): "
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return &Operations{list: l}
}
