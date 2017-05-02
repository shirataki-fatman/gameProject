package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"gameProject/entities"
	"gameProject/systems"
	"image/color"
)

type MainScene struct{}

func (*MainScene) Type() string {
	return "MainScene"
}

func (m *MainScene) Preload() {
	err := engo.Files.Load("textures/player.png")
	if err != nil {
		panic(err)
	}
}

func (*MainScene) Setup(world *ecs.World) {
	common.SetBackground(color.White)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.MainSystem{})
	world.AddSystem(&systems.InputSystem{})
	world.AddSystem(&systems.PlayerUpdateSystem{})

	entities.NewPlayer(world)
}
