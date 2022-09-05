package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	DIRECTION_LEFT int = iota
	DIRECTION_RIGHT
	DIRECTION_UP
	DIRECTION_DOWN
)

type Location struct {
	X int
	Y int
}

type Snake struct {
	Body      []Location
	Food      Location
	FoodEated bool
	Direction int
	MaxX      int
	MaxY      int
	Score     int
}

// 获取蛇头位置
func (s *Snake) GetHead() Location {
	return s.Body[len(s.Body)-1]
}

// 移动一步, 如果碰壁返回false, 否则返回true
func (s *Snake) Move() bool {
	head := s.GetHead()
	switch s.Direction {
	case DIRECTION_UP:
		s.Body = append(s.Body, Location{head.X, head.Y - 1})
	case DIRECTION_DOWN:
		s.Body = append(s.Body, Location{head.X, head.Y + 1})
	case DIRECTION_LEFT:
		s.Body = append(s.Body, Location{head.X - 1, head.Y})
	case DIRECTION_RIGHT:
		s.Body = append(s.Body, Location{head.X + 1, head.Y})
	}
	head = s.GetHead()

	// 蛇头到达食物位置时标记食物已吃，并且追加到蛇尾(s.Body[0]不用剔除, 否则剔除)
	if head == s.Food {
		s.FoodEated = true
		s.RandomFood()
		s.Score += 10
	} else {
		s.Body = s.Body[1:]
	}
	return 0 <= head.X && head.X <= s.MaxX && 0 <= head.Y && head.Y <= s.MaxY
}

// 判断生成的食物坐标是否在蛇身上
func (s *Snake) isFoodInSnake(location Location) bool {
	for _, l := range s.Body {
		if l == location {
			return true
		}
	}
	return false
}

// 生成食物
func (s *Snake) RandomFood() {
	w, h := termbox.Size()
	// 上下两边留点空隙
	location := Location{rand.Intn(w-10) + 5, rand.Intn(h-10) + 5}
	for s.isFoodInSnake(location) {
		location = Location{rand.Intn(w), rand.Intn(h)}
	}
	s.Food = location
}

func Draw(s *Snake) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, location := range s.Body {
		termbox.SetCell(location.X, location.Y, '●', termbox.ColorGreen, termbox.ColorDefault)
	}
	termbox.SetCell(s.Food.X, s.Food.Y, '●', termbox.ColorRed, termbox.ColorDefault)
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	w, h := termbox.Size()

	// 初始给它三个长度吧, 太小不好看
	snake := Snake{
		Body:      []Location{{0, 0}, {1, 0}, {2, 0}},
		Direction: DIRECTION_RIGHT,
		MaxX:      w,
		MaxY:      h,
		FoodEated: false,
	}
	snake.RandomFood()
	Draw(&snake)

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	gameFinished := false
	msgPrinted := false
	msg := `\n
*****************************************
		Game Over !
		Score: %d
		Press Esc to exit!
*****************************************
`
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			} else if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyArrowUp:
					snake.Direction = DIRECTION_UP
				case termbox.KeyArrowDown:
					snake.Direction = DIRECTION_DOWN
				case termbox.KeyArrowLeft:
					snake.Direction = DIRECTION_LEFT
				case termbox.KeyArrowRight:
					snake.Direction = DIRECTION_RIGHT
				}
			}
		default:
			time.Sleep(300 * time.Millisecond)
			if gameFinished && !msgPrinted {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				termbox.Flush()
				fmt.Printf(msg, snake.Score)
				msgPrinted = true
			} else {
				if success := snake.Move(); !success {
					gameFinished = true
				}
				Draw(&snake)
			}
		}
	}
}
