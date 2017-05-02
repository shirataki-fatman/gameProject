package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"gameProject/components"
	"gameProject/systems"
)

type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	components.InputComponent
}

func NewPlayer(world *ecs.World) *Player {
	player := Player{
		BasicEntity: ecs.NewBasic(),
	}
	player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 400,
			Y: 300,
		},
		Width:  30,
		Height: 30,
	}

	texture, err := common.LoadedSprite("textures/player.png")
	if err != nil {
		panic(err)
	}

	player.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
		case *systems.InputSystem:
			sys.Add(&player.BasicEntity, &player.InputComponent)
		case *systems.PlayerUpdateSystem:
			sys.Add(&player.BasicEntity, &player.InputComponent, &player.SpaceComponent)
		}
	}

	return &player
}
