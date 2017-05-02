package scenes

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"gameProject/systems"
	"image/color"
	"engo.io/engo"
)

type TitleScene struct{}

func (*TitleScene) Type() string {
	return "TitleScene"
}

func (*TitleScene) Preload() {}

func (*TitleScene) Setup(world *ecs.World) {
	engo.Input.RegisterButton("Push", engo.A)
	engo.Input.RegisterButton("Up", engo.ArrowUp)
	engo.Input.RegisterButton("Down", engo.ArrowDown)
	engo.Input.RegisterButton("Left", engo.ArrowLeft)
	engo.Input.RegisterButton("Right", engo.ArrowRight)

	common.SetBackground(color.Black)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.TitleSystem{})
}
