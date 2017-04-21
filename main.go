package main

import (
	"engo.io/engo"
	"gameProject/scenes"
)

func main() {
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  800,
		Height: 600,
	}

	engo.RegisterScene(&scenes.MainScene{})

	engo.Run(opts, &scenes.TitleScene{})
}
