package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
)

type TitleSystem struct{}

func (*TitleSystem) Remove(ecs.BasicEntity) {}

func (*TitleSystem) Update(dt float32) {
	if engo.Input.Button("Push").JustPressed() {
		fmt.Println("Test")
		engo.SetSceneByName("MainScene", true)
	}
}
