package systems

import (
	"engo.io/ecs"
)

type MainSystem struct{}

func (*MainSystem) Remove(ecs.BasicEntity) {}

func (*MainSystem) Update(dt float32) {
}
