package style

import "ds-protector/pkg/tui"

func GetSelfWindowSize() (width, headerHeight, contentHeight int) {
	w, h := tui.GetWindowSize()
	// 移除左右占用掉的2格border
	width = w - 2
	// 标题部分高度
	headerHeight = 6
	// content部分高度
	contentHeight = h - headerHeight - 4

	return width, headerHeight, contentHeight
}
