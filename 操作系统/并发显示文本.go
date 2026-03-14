package main

import (
	"context"
	"fmt"
	"time"
)

type Task struct {
	ID       int
	Priority int
	Text     string
	Fg       string
	Weight   int
}

type RenderMsg struct {
	TaskText string
	Priority int
	Styled   string
	Step     int
}

const (
	reset   = "\x1b[0m"
	bgBlack = "\x1b[40m"
	fgRed   = "\x1b[31m"
	fgWhite = "\x1b[37m"
)

func clear()            { fmt.Print("\x1b[2J\x1b[H") }
func move(row, col int) { fmt.Printf("\x1b[%d;%dH", row, col) }

func styleLabel(fg, text string) string {
	// 固定宽度，覆盖时不残留
	label := fmt.Sprintf("  %s  ", text) // 8 chars visible
	return bgBlack + fg + label + reset
}

// taskWorker: 模拟“任务”并发存在，但只有收到 token 才能“运行一次时间片并输出”
func taskWorker(ctx context.Context, t Task, token <-chan int, render chan<- RenderMsg) {
	for {
		select {
		case <-ctx.Done():
			return
		case step := <-token:
			render <- RenderMsg{
				TaskText: t.Text,
				Priority: t.Priority,
				Styled:   styleLabel(t.Fg, t.Text),
				Step:     step,
			}
		}
	}
}

// scheduler: 按权重(=优先级)给每个任务发 token，权重大者获得更多时间片
func scheduler(ctx context.Context, tasks []Task, tokenChans []chan int, tick time.Duration) {
	step := 0
	for {
		for i := range tasks {
			for k := 0; k < tasks[i].Weight; k++ {
				select {
				case <-ctx.Done():
					return
				case tokenChans[i] <- step:
					step++
				}
				time.Sleep(tick)
			}
		}
	}
}

// renderer: 单线程消费渲染消息，保证“独占输出到同一区域”
func renderer(ctx context.Context, render <-chan RenderMsg) {
	sharedRow, sharedCol := 4, 1
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-render:
			// 共享区域覆盖输出
			move(sharedRow, sharedCol)
			fmt.Print(msg.Styled)

			// 状态行固定位置输出
			move(sharedRow+2, 1)
			fmt.Print("\x1b[0K") // 清到行尾
			fmt.Printf("step=%d  running=%s  prio=%d        ", msg.Step, msg.TaskText, msg.Priority)
		}
	}
}

func main() {
	tasks := []Task{
		{ID: 0, Priority: 16, Text: "HUST", Fg: fgRed, Weight: 16},
		{ID: 1, Priority: 10, Text: "MRSU", Fg: fgWhite, Weight: 10},
		{ID: 2, Priority: 8, Text: "VERY", Fg: fgRed, Weight: 8},
		{ID: 3, Priority: 6, Text: "LOVE", Fg: fgWhite, Weight: 6},
	}

	const tick = 80 * time.Millisecond

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clear()
	fmt.Println("Priority scheduler demo (goroutines + tokens): exclusive output in ONE shared area. Ctrl+C to stop.")
	fmt.Println("Shared area below will be overwritten by the running task:")
	fmt.Println()

	// 每个任务一个 token channel（容量 1 防止调度器阻塞太久）
	tokenChans := make([]chan int, len(tasks))
	for i := range tasks {
		tokenChans[i] = make(chan int, 1)
	}

	// 渲染通道（缓冲避免任务阻塞）
	renderCh := make(chan RenderMsg, 32)

	// 启动渲染器
	go renderer(ctx, renderCh)

	// 启动4个“任务”
	for i := range tasks {
		go taskWorker(ctx, tasks[i], tokenChans[i], renderCh)
	}

	// 启动调度器
	go scheduler(ctx, tasks, tokenChans, tick)

	// 阻塞主协程
	select {}
}
