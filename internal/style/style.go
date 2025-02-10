package style

import "github.com/charmbracelet/lipgloss"

var TitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#2EC5AA")).PaddingBottom(1).PaddingLeft(2)

// HeaderBoxStyle 导航栏Box样式
var HeaderBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder())

// ContentBoxStyle 内容Box样式
var ContentBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder())

// HelpStyle 帮助文档的样式
var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingLeft(2).PaddingTop(1)

// FormStyle 表单基础样式
var FormStyle = lipgloss.NewStyle().PaddingLeft(2)

// SelectedStyle 选中样式
var SelectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#A62CD5"))

// LogBodyStyle 日志样式
var LogBodyStyle = lipgloss.NewStyle().PaddingLeft(4).PaddingTop(1)

// VulStyle 漏洞显示样式
var VulStyle = lipgloss.NewStyle().Background(lipgloss.Color("#FF2525"))

// LogTypeAttack 攻击日志
var LogTypeAttack = lipgloss.NewStyle().Foreground(lipgloss.Color("#F70202"))

// LogTypeDefence 攻击日志
var LogTypeDefence = lipgloss.NewStyle().Foreground(lipgloss.Color("#01C719"))

// MessageStyle 的主体样式
var MessageStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62"))

// MessageTitle 消息标题
var MessageTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0ABD87"))

// MessageContent 消息内容
var MessageContent = lipgloss.NewStyle().Align(lipgloss.Center).PaddingTop(1)

// MessageWechat 微信样式
var MessageWechat = lipgloss.NewStyle().Align(lipgloss.Center).Foreground(lipgloss.Color("#00E200")).PaddingTop(1)
