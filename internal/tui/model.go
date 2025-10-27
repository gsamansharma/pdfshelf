package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"pdfshelf/internal/model"
)

var (
	docStyle      = lipgloss.NewStyle().Margin(1, 2)
	titleStyle    = lipgloss.NewStyle().MarginLeft(2)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item struct {
	entry model.PDFEntry
	index int 
}

func (i item) Title() string { return i.entry.Name }
func (i item) Description() string {
	duration := i.entry.TotalTimeSpent.Round(time.Second).String()
	return fmt.Sprintf("Time: %s | Path: %s", duration, i.entry.FilePath)
}
func (i item) FilterValue() string { return i.entry.Name }

type TUIModel struct {
	List           list.Model
	SelectedEntry  *model.PDFEntry 
	SelectedIndex  int            
	quitting       bool
}

func New(pdfEntries []model.PDFEntry) TUIModel {
	items := make([]list.Item, len(pdfEntries))
	for i, entry := range pdfEntries {
		items[i] = item{entry: entry, index: i}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Your PDF Shelf"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return TUIModel{List: l}
}

func (m TUIModel) Init() tea.Cmd {
	return nil
}

func (m TUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		m.List.SetHeight(msg.Height - 1) 
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.List.SelectedItem().(item)
			if ok {
				m.SelectedEntry = &i.entry
				m.SelectedIndex = i.index
			}
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m TUIModel) View() string {
	if m.quitting {
		if m.SelectedEntry != nil {
			return quitTextStyle.Render(fmt.Sprintf("Opening %s...", m.SelectedEntry.Name))
		}
		return quitTextStyle.Render("Bye!")
	}
	return docStyle.Render(m.List.View())
}

func StartTUI(pdfEntries []model.PDFEntry) (*model.PDFEntry, int, error) {
	m := New(pdfEntries)
	p := tea.NewProgram(m, tea.WithAltScreen())

	finalModel, err := p.Run()
	if err != nil {
		return nil, -1, err
	}

	finalTUIModel, ok := finalModel.(TUIModel)
	if !ok {
		return nil, -1, fmt.Errorf("Could not cast final model")
	}

	if finalTUIModel.SelectedEntry != nil {
		return finalTUIModel.SelectedEntry, finalTUIModel.SelectedIndex, nil
	}

	return nil, -1, nil
}