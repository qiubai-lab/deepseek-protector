package views

import (
	"ds-protector/internal/style"
	"ds-protector/pkg/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Message struct {
}

func (m Message) Init() tea.Cmd {
	return nil
}

func (m Message) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Message) View() string {
	w, h := tui.GetWindowSize()
	title := style.MessageTitle.Width(w / 3).Align(lipgloss.Center).Render("感谢您为DeepSeek保卫战共享自己的力量!")
	content := style.MessageContent.Width(w / 3).Render("欢迎您关注我们的微信公众号获取更多网络安全资讯以及AI相关技巧和知识!")
	wechat := style.MessageWechat.Width(w / 3).Render("公众号: 487Donkey Sec")
	body := style.MessageStyle.Render(lipgloss.JoinVertical(0, title, content, wechat))

	return lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, body)
}

func NewMessage() tea.Model {
	return &Message{}
}
