package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"gameProject/components"
)

type inputEntity struct {
	*ecs.BasicEntity
	*components.InputComponent
}

type InputSystem struct {
	entities []inputEntity
}

func (is *InputSystem) Remove(basic ecs.BasicEntity) {
	var delete int = -1
	for index, entity := range is.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		is.entities = append(is.entities[:delete], is.entities[delete+1:]...)
	}
}

func (is *InputSystem) Update(dt float32) {
	for _, entity := range is.entities {
		if engo.Input.Button("Down").Down() {
			entity.DownButton = 1
		} else {
			entity.DownButton = 0
		}

		if engo.Input.Button("Up").Down() {
			entity.UpButton = 1
		} else {
			entity.UpButton = 0
		}

		if engo.Input.Button("Left").Down() {
			entity.LeftButton = 1
		} else {
			entity.LeftButton = 0
		}

		if engo.Input.Button("Right").Down() {
			entity.RightButton = 1
		} else {
			entity.RightButton = 0
		}
	}
}

func (is *InputSystem) Add(entity *ecs.BasicEntity, inputComponent *components.InputComponent) {
	is.entities = append(is.entities, inputEntity{entity, inputComponent})
}
