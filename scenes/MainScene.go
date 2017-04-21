package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"gameProject/systems"
	"image/color"
)

type MainScene struct{}

func (*MainScene) Type() string {
	return "MainScene"
}

func (*MainScene) Preload() {}

func (*MainScene) Setup(world *ecs.World) {
	engo.Input.RegisterButton("Push", engo.A)

	common.SetBackground(color.White)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.MainSystem{})
}
