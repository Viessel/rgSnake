package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Down
	Right
	Left
)
const HEIGHT = 800
const WIDTH = 800
const STEP = HEIGHT / 20

type Node struct {
	x    int32
	y    int32
	next *Node
}

func (n *Node) CheckCollision() int {
	i := n.next
	for i != nil {
		if i.x == n.x && i.y == n.y {
			return 1
		}
		i = i.next
	}
	return 0
}
func (n *Node) Push() *Node {
	node := Node{x: n.x, y: n.y}
	n.next = &node
	return &node
}
func (n *Node) Move(dir Direction) {
	n.Propagate(n.x, n.y)
	switch dir {
	case Up:
		{
			n.y += STEP
			if n.y >= HEIGHT {
				n.y = 0
			}

		}
	case Down:
		{
			n.y -= STEP
			if n.y < 0 {
				n.y = HEIGHT - STEP
			}
		}
	case Left:
		{
			n.x -= STEP
			if n.x < 0 {
				n.x = WIDTH - STEP
			}
		}
	case Right:
		{
			n.x += STEP
			if n.x >= WIDTH {
				n.x = 0
			}
		}
	}
}

func (n *Node) Propagate(x, y int32) {
	// not tail recursive :')
	if n.next != nil {
		n.next.Propagate(n.x, n.y)
	}
	n.x = x
	n.y = y
}

func drawSnake(n *Node) {
	i := n
	for {
		rl.DrawRectangle(i.x, i.y, STEP, STEP, rl.Orange)
		i = i.next
		if i == nil {
			break
		}
	}
}

func spawnFood(food *Node) {
	h := STEP * rand.Intn(HEIGHT/STEP)
	w := STEP * rand.Intn(WIDTH/STEP)
	food.x = int32(w)
	food.y = int32(h)
}
func main() {
	rl.InitWindow(WIDTH, HEIGHT, "raylib example")
	defer rl.CloseWindow()
	rl.SetTargetFPS(10)

	var direction Direction = Left
	var head Node = Node{x: HEIGHT / 2, y: WIDTH / 2}
	tail := &head
	var food Node = Node{}
  GAME_OVER := 0
	spawnFood(&food)

  input := rl.KeyD
	for !rl.WindowShouldClose() {
    input = int(rl.GetKeyPressed())
    if input == (rl.KeyS) {direction = Up}
    if input == (rl.KeyW) {direction = Down}
    if input == (rl.KeyD) {direction = Right}
    if input == (rl.KeyA) {direction = Left}
		head.Move(direction)
    if head.CheckCollision() == 1 {
      GAME_OVER = 1
    }
		if head.x == food.x && head.y == food.y {
			spawnFood(&food)
			tail = tail.Push()
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGreen)
		drawSnake(&head)
		rl.DrawRectangle(food.x, food.y, STEP, STEP, rl.DarkBlue)
		if GAME_OVER == 1 {
      w := rl.MeasureText("GAME OVER", 100)
			rl.DrawText("GAME OVER", (WIDTH - w)/2, HEIGHT/2-50, 100, rl.Red)
      rl.EndDrawing()
			time.Sleep(5 * time.Second)
			//    head.Restart()
			direction = Left
			head.x =  HEIGHT / 2
      head.y = WIDTH / 2 
      head.next = nil
			tail = &head
			spawnFood(&food)
      GAME_OVER = 0
      continue
		}
		//time.Sleep(150 * time.Millisecond)
		rl.EndDrawing()
	}
}
