package radiobutton

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// 设置为淡灰色
var otherValueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

// 设置为白色
var selectedValueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))

type Item struct {
	Title string
	Url   string
	Style lipgloss.Style
}

func (i *Item) RenderTitle() string {
	return i.Style.Render(i.Title)
}

type Model struct {
	Items []Item // 选项列表
	// 选择的下标
	Selected int
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "left", "a":
			m.Prev()
		case "right", "d":
			m.Next()
		}
	}
	return m, nil
}

func (m *Model) View() string {
	currentValue := m.Items[m.Selected]
	//previousValue := m.Items[(m.Selected-1+len(m.Items))%len(m.Items)]
	nextValue := m.Items[(m.Selected+1)%len(m.Items)]

	return fmt.Sprintf("%s > %s",
		currentValue.RenderTitle(),
		otherValueStyle.Render(nextValue.Title))
}

func (m *Model) GetValue() Item {
	return m.Items[m.Selected]
}

// Next 选择下一个选项
func (m *Model) Next() {
	m.Selected = (m.Selected + 1) % len(m.Items)
}

// Prev 选择上一个选项
func (m *Model) Prev() {
	m.Selected = (m.Selected - 1 + len(m.Items)) % len(m.Items)
}

func NewModel(items []Item, defaultSelected int) *Model {
	return &Model{
		Items:    items,
		Selected: defaultSelected,
	}
}
