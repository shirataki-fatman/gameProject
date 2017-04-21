package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"gameProject/systems"
	"image/color"
)

type TitleScene struct{}

func (*TitleScene) Type() string {
	return "TitleScene"
}

func (*TitleScene) Preload() {}

func (*TitleScene) Setup(world *ecs.World) {
	engo.Input.RegisterButton("Push", engo.A)

	common.SetBackground(color.Black)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.TitleSystem{})
}
