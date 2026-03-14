package main

import (
	"fmt"
	"time"
)

type Task struct {
	Priority int
	Text     string
	Fg       string
	Weight   int
}

const (
	reset   = "\x1b[0m"
	bgBlack = "\x1b[40m"
	fgRed   = "\x1b[31m"
	fgWhite = "\x1b[37m"
)

func clear()            { fmt.Print("\x1b[2J\x1b[H") }
func move(row, col int) { fmt.Printf("\x1b[%d;%dH", row, col) }

// Render one "label" in the SAME area (same cursor position), overwriting previous one.
func renderLabel(t Task) string {
	// fixed width so覆盖时不会残留： "  XXXX  " 8 chars
	label := fmt.Sprintf("  %s  ", t.Text)
	return bgBlack + t.Fg + label + reset
}

func main() {
	tasks := []Task{
		{Priority: 16, Text: "HUST", Fg: fgRed, Weight: 16},
		{Priority: 10, Text: "MRSU", Fg: fgWhite, Weight: 10},
		{Priority: 8, Text: "VERY", Fg: fgRed, Weight: 8},
		{Priority: 6, Text: "LOVE", Fg: fgWhite, Weight: 6},
	}

	const tick = 80 * time.Millisecond

	clear()
	fmt.Println("Priority scheduler demo: exclusive output in ONE shared area. Ctrl+C to stop.")
	fmt.Println("Shared area below will be overwritten by the running task:")
	fmt.Println()

	sharedRow, sharedCol := 4, 1 // 共享输出区域左上角（同一位置）
	step := 0

	for {
		for i := range tasks {
			for k := 0; k < tasks[i].Weight; k++ {
				// 在同一个位置输出当前任务字符串（覆盖之前的）
				move(sharedRow, sharedCol)
				fmt.Print(renderLabel(tasks[i]))

				// 下面打印状态（固定位置），避免滚屏
				move(sharedRow+2, 1)
				fmt.Print("\x1b[0K") // 清除到行尾
				fmt.Printf("step=%d  running=%s  prio=%d        ", step, tasks[i].Text, tasks[i].Priority)

				step++
				time.Sleep(tick)
			}
		}
	}
}
