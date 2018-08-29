package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type scene int

const (
	start scene = iota
	menu
	end
)

func (s *scene) nextScene() {
	*s = (*s + 1) % 3
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Demo 1",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.White)

	// initialize the scene
	s := start

	for !win.Closed() {
		switch s {
		case start:
			win.Clear(colornames.Red)
		case menu:
			win.Clear(colornames.Blue)
		case end:
			win.Clear(colornames.Green)
		}

		if win.JustPressed(pixelgl.KeyEnter) {
			s.nextScene()
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
