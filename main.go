package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type cell struct {
	Pos   rl.Vector2
	State bool
}

func update(grid *[50][50]bool) {

	m := make(map[cell]int)
	for y := range grid {
		for x := range grid[y] {
			friendsCounter := 0
			if y > 0 {
				if x > 0 {
					if grid[y-1][x-1] { //top left
						friendsCounter++
					}
				}
				if grid[y-1][x] { //top middle
					friendsCounter++
				}
				if x < len(grid[y])-1 {
					if grid[y-1][x+1] { //top right
						friendsCounter++
					}

				}
			}
			if x > 0 {
				if grid[y][x-1] { //middle left
					friendsCounter++
				}

			}
			if x < len(grid[y])-1 {
				if grid[y][x+1] { //middle right
					friendsCounter++
				}

			}
			if y < len(grid)-1 {
				if x > 0 {
					if grid[y+1][x-1] { //down left
						friendsCounter++
					}

				}
				if grid[y+1][x] { //down middle
					friendsCounter++
				}
				if x < len(grid[y])-1 {
					if grid[y+1][x+1] { //down right
						friendsCounter++
					}
				}
			}
			m[cell{rl.Vector2{X: float32(x), Y: float32(y)}, grid[y][x]}] = friendsCounter
		}
	}
	for k, friends := range m {
		if k.State {
			if friends < 2 {
				grid[int32(k.Pos.Y)][int32(k.Pos.X)] = false
			} else if friends == 2 || friends == 3 {
				grid[int32(k.Pos.Y)][int32(k.Pos.X)] = true
			} else if friends > 3 {
				grid[int32(k.Pos.Y)][int32(k.Pos.X)] = false
			}

		} else if !grid[int32(k.Pos.Y)][int32(k.Pos.X)] && friends == 3 {
			grid[int32(k.Pos.Y)][int32(k.Pos.X)] = true
		}
	}

}
func main() {
	var grid [50][50]bool
	rl.InitWindow(800, 800, "Life")
	rl.SetTargetFPS(60)
	// rl.SetWindowState(rl.FlagWindowResizable)
	// colors := []rl.Color{rl.Blue, rl.Red, rl.Orange, rl.Green}
	fpsCounter := 0
	start := false
	for !rl.WindowShouldClose() {
		fpsCounter++
		if rl.IsKeyPressed(rl.KeyS) {
			start = true
		} else if rl.IsKeyPressed(rl.KeyP) {
			fpsCounter = 0
			start = false
		}
		if start {
			fpsCounter++
			if fpsCounter >= 30 {
				fpsCounter = 0
				update(&grid)

			}
		}
		for y := range grid {
			for x := range grid[y] {
				if rl.IsMouseButtonDown(rl.MouseLeftButton) {
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{float32(x * rl.GetScreenWidth() / len(grid[y])), float32(y * rl.GetScreenHeight() / len(grid)), float32(rl.GetScreenWidth() / len(grid[y])), float32(rl.GetScreenHeight() / len(grid))}) {
						grid[y][x] = true
					}
				}
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] {
					rl.DrawRectangle(int32(x*rl.GetScreenWidth()/len(grid[y])), int32(y*rl.GetScreenHeight()/len(grid)), int32(rl.GetScreenWidth()/len(grid[y])), int32(rl.GetScreenHeight()/len(grid)), rl.Green)
				} else {
					rl.DrawRectangleLines(int32(x*rl.GetScreenWidth()/len(grid[y])), int32(y*rl.GetScreenHeight()/len(grid)), int32(rl.GetScreenWidth()/len(grid[y])), int32(rl.GetScreenHeight()/len(grid)), rl.Black)
				}
			}
		}
		if start {
			rl.DrawText("Started", 0, 0, 30, rl.DarkGreen)
		} else {
			rl.DrawText("stoped", 0, 0, 30, rl.Red)

		}
		// rl.DrawFPS(200, 0)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
