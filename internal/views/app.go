package views

import (
	"ds-protector/internal/dicts"
	"ds-protector/internal/style"
	"ds-protector/pkg/components/radiobutton"
	"ds-protector/pkg/tui"
	"ds-protector/pkg/utils"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"math/rand"
	"time"
)

const (
	target    = 0
	intensity = 1
)

var intensityValue string

type App struct {
	// 选项
	targetCheck        *radiobutton.Model
	attackIntensity    *radiobutton.Model
	currentSelectIndex int
	// 内容部分
	spinner  spinner.Model
	fixQueue *utils.FixedQueue
	// 状态
	attack bool
	// Message
	showMessage bool
	message     tea.Model
}
type result struct {
	duration time.Duration
	emoji    string
}

func (a *App) NextInput() {
	a.currentSelectIndex++
	if a.currentSelectIndex > 1 {
		a.currentSelectIndex = 0
	}
}

func (a *App) Init() tea.Cmd {
	return tea.Batch(
		a.spinner.Tick,
		runPretendProcess,
	)
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if a.currentSelectIndex == target {
		a.targetCheck, _ = a.targetCheck.Update(msg)
	}
	if a.currentSelectIndex == intensity {
		a.attackIntensity, _ = a.attackIntensity.Update(msg)
		// 更新后设置强度值
		intensityValue = a.attackIntensity.GetValue().Title
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEnter.String():
			a.attack = !a.attack
		case tea.KeyEsc.String():
			return a, tea.Quit
		case tea.KeyTab.String():
			a.NextInput()
		case tea.KeySpace.String():
			a.showMessage = !a.showMessage
		}
	case tea.WindowSizeMsg:
		// 更新窗口大小
		tui.SetWindowSize(msg.Width, msg.Height)

	case spinner.TickMsg:
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		return a, cmd
	case processFinishedMsg:
		if a.attack {
			a.fixQueue.Enqueue(dicts.GetRandomLog())
		}
		return a, runPretendProcess
	}
	return a, nil
}

func (a *App) View() string {
	// 判断显示的视图
	if a.showMessage {
		return a.message.View()
	}
	w, headerHeight, contentHeight := style.GetSelfWindowSize()
	// 获取头部内容
	headerBody := a.getHeaderBody(w, headerHeight)
	// 获取内容部分
	contentBody := a.getContentBody(w, contentHeight)

	return lipgloss.JoinVertical(0, headerBody, contentBody)
}

func (a *App) renderPrompt(index int, prompt string) string {
	if index == a.currentSelectIndex {
		return style.SelectedStyle.Render(fmt.Sprintf(">%s", prompt))
	} else {
		return fmt.Sprintf(" %s", prompt)
	}
}

func (a *App) getLogBody() string {
	_, _, h := style.GetSelfWindowSize()
	var logs []string
	logHeight := h - 2
	for _, log := range a.fixQueue.GetLastN(logHeight) {
		logs = append(logs, log.(string))
	}

	return style.LogBodyStyle.Height(logHeight).Render(lipgloss.JoinVertical(0, logs...))
}

func (a *App) getHeaderBody(w, h int) string {
	// 标题
	titleBody := style.TitleStyle.Render("【红客联盟487分盟】感谢您对DeepSeek保卫战的支持！")

	// 目标选择
	targetCheckBody := style.FormStyle.Render(fmt.Sprintf("%s: %s", a.renderPrompt(target, "攻击目标"), a.targetCheck.View()))
	speedCheckBody := style.FormStyle.Render(fmt.Sprintf("%s: %s", a.renderPrompt(intensity, "攻击强度"), a.attackIntensity.View()))
	// 帮助文档
	helpBody := style.HelpStyle.Render("[Tab: 切换] [←/→: 选择] [回车: 发起攻击/结束攻击] [空格: 红客声明] [ESC: 结束攻击]")
	return style.HeaderBoxStyle.Width(w).Height(h).Render(
		lipgloss.JoinVertical(0, titleBody, targetCheckBody, speedCheckBody, helpBody),
	)
}

func (a *App) getContentBody(w, h int) string {

	t := a.targetCheck.GetValue()
	i := a.attackIntensity.GetValue()
	var tips string
	if a.attack {
		tips = fmt.Sprintf(" %s 正在对[%s]发起攻击 攻击等级[%s]...", a.spinner.View(), t.RenderTitle(), i.RenderTitle())
	} else {
		tips = fmt.Sprintf(" %s 按下回车对[%s]发起攻击 攻击等级[%s]...", "|", t.RenderTitle(), i.RenderTitle())
	}

	logBody := a.getLogBody()
	return style.ContentBoxStyle.Width(w).Height(h).Render(lipgloss.JoinVertical(0, tips, logBody))
}

// processFinishedMsg is sent when a pretend process completes.
type processFinishedMsg time.Duration

// pretendProcess simulates a long-running process.
func runPretendProcess() tea.Msg {
	var pause time.Duration

	switch intensityValue {
	case "强力":
		// 设置较长的 pause 时间
		pause = time.Duration(rand.Int63n(100)+400) * time.Millisecond // 1s 到 2s
	case "破坏":
		// 设置中等的 pause 时间
		pause = time.Duration(rand.Int63n(100)+200) * time.Millisecond // 0.5s 到 1.5s
	case "猛攻":
		// 设置较短的 pause 时间
		pause = time.Duration(rand.Int63n(10)+50) * time.Millisecond // 100ms 到 1s
	default:
		// 如果没有匹配的状态，默认使用强力状态
		pause = time.Duration(rand.Int63n(100)+500) * time.Millisecond
	}

	time.Sleep(pause)
	return processFinishedMsg(pause)
}

func randomEmoji() string {
	emojis := []rune("🍦🍡🤠👾😭🦊🐯🦆🥨🎏🍔🍒🍥🎮📦🦁🐶🐸🍕🥐🧲🚒🥇🏆🌽")
	return string(emojis[rand.Intn(len(emojis))]) // nolint:gosec
}

func NewApp() tea.Model {

	targetCheck := radiobutton.NewModel(
		[]radiobutton.Item{
			{
				Title: "五角大楼",
				Url:   "https://www.google.com",
				Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")),
			},
			{
				Title: "美国白宫",
				Url:   "https://www.google.com",
				Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")),
			},
			{
				Title: "NASA卫星",
				Url:   "https://www.google.com",
				Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")),
			},
		}, 0)

	attackIntensity := radiobutton.NewModel(
		[]radiobutton.Item{
			{
				Title: "强力",
				Url:   "Normal",
				Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#4FE36A")),
			},
			{
				Title: "破坏",
				Url:   "Strong",
				Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5733")),
			},
			{
				Title: "猛攻",
				Url:   "Broken",
				Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#E40404")),
			},
		}, 0)

	sp := spinner.New()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("206"))

	return &App{
		targetCheck:     targetCheck,
		attackIntensity: attackIntensity,
		spinner:         sp,
		fixQueue:        utils.NewFixedQueue(50),
		message:         NewMessage(),
	}
}
