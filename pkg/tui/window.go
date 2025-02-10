package tui

import "sync"

// 宽度
var width int

// 高度
var height int

var sizeLock sync.Mutex

func SetWindowSize(w, h int) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	width = w
	height = h
}

func GetWindowSize() (int, int) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	return width, height
}
