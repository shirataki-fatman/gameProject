package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"gameProject/components"
)

const (
	PLAYER_MOVE_SPEED = 100
)

type playerEntity struct {
	*ecs.BasicEntity
	*components.InputComponent
	*common.SpaceComponent
}

type PlayerUpdateSystem struct {
	entities []playerEntity
}

func (pus *PlayerUpdateSystem) Remove(basic ecs.BasicEntity) {
	var delete int = -1
	for index, entity := range pus.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		pus.entities = append(pus.entities[:delete], pus.entities[delete+1:]...)
	}
}

func (is *PlayerUpdateSystem) Update(dt float32) {
	for _, entity := range is.entities {
		var vx, vy float32

		if entity.RightButton > 0 {
			vx += PLAYER_MOVE_SPEED
		}
		if entity.LeftButton > 0 {
			vx -= PLAYER_MOVE_SPEED
		}

		if entity.UpButton > 0 {
			vy -= PLAYER_MOVE_SPEED
		}
		if entity.DownButton > 0 {
			vy += PLAYER_MOVE_SPEED
		}

		entity.Position.X += vx * dt
		entity.Position.Y += vy * dt
	}
}

func (is *PlayerUpdateSystem) Add(entity *ecs.BasicEntity, inputComponent *components.InputComponent, spaceComponent *common.SpaceComponent) {
	is.entities = append(is.entities, playerEntity{entity, inputComponent, spaceComponent})
}
