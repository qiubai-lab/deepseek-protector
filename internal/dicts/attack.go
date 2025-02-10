package dicts

import (
	"ds-protector/internal/style"
	"fmt"
	"math/rand"
	"time"
)

const (
	Attack  = 0
	Defense = 1
)

type Log struct {
	Type    int
	Emoji   string
	Content string
}

func GetRandomLog() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 创建新的随机数生成器
	randomIndex := r.Intn(len(logList))                  // 生成一个随机索引
	log := logList[randomIndex]                          // 获取随机日志

	var logString string
	if log.Type == Attack {
		logString = fmt.Sprintf("%s - [%s] 成功对目标发起%s攻击...", style.LogTypeAttack.Render(log.Emoji), time.Now().Format("2006/01/02/15:04:05"), style.VulStyle.Render(log.Content))
	} else {
		logString = fmt.Sprintf("%s - [%s] 成功防御目标发起的%s攻击...", style.LogTypeDefence.Render(log.Emoji), time.Now().Format("2006/01/02/15:04:05"), style.VulStyle.Render(log.Content))
	}
	return logString
}

var logList = []Log{
	{
		Type:    Attack,
		Emoji:   "攻击", // 代表破解或绕过
		Content: "序列化",
	},
	{
		Type:    Attack,
		Emoji:   "攻击", // 代表数据存储，适合 SQL 注入
		Content: "SQL注入",
	},
	{
		Type:    Attack,
		Emoji:   "攻击", // 代表延迟或被迫停止
		Content: "服务停止",
	},
	{
		Type:    Attack,
		Emoji:   "攻击", // 代表监视或洞察，适合天神之眼
		Content: "天神之眼",
	},
	{
		Type:    Attack,
		Emoji:   "攻击", // 代表拒绝服务或阻止
		Content: "服务阻断",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "数据洪流攻击",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "可视化漏洞挖掘",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "深入探测攻击",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "存储耗尽掠夺",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "虚拟机崩溃",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "全球网络轰炸",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "信号干扰攻击",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "链式反应安全漏洞",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "防御系统破解",
	},
	{
		Type:    Attack,
		Emoji:   "攻击",
		Content: "密钥暴力破解",
	},
	{
		Type:    Defense,
		Emoji:   "攻击",
		Content: "量子崩坏脉冲",
	},
	{
		Type:    Defense,
		Emoji:   "防御",
		Content: "网络过载风暴",
	},
	{
		Type:    Defense,
		Emoji:   "防御",
		Content: "数据吞噬",
	},
	{
		Type:    Defense,
		Emoji:   "防御",
		Content: "防火墙熔毁",
	},
	{
		Type:    Defense,
		Emoji:   "防御",
		Content: "协议漩涡",
	},
	{
		Type:    Defense,
		Emoji:   "防御",
		Content: "超频数据闪电",
	},
	{
		Type:    Attack,
		Emoji:   "防御",
		Content: "协议入侵",
	},
	{
		Type:    Defense,
		Emoji:   "防御",
		Content: "数据撕裂",
	},
	{
		Type:    Defense,
		Emoji:   "防御",
		Content: "预言式零日打击",
	},
	{
		Type:    Defense,
		Emoji:   "防御",
		Content: "死亡数据包",
	},
}
