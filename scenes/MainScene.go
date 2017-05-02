package scenes

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"gameProject/systems"
	"image/color"
	"engo.io/engo"
	"gameProject/entities"
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
