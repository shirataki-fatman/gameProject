package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
)

type MainSystem struct{}

func (*MainSystem) Remove(ecs.BasicEntity) {}

func (*MainSystem) Update(dt float32) {
	if engo.Input.Button("Push").JustPressed() {
		fmt.Println("Hoge")
		engo.SetSceneByName("TitleScene", true)
	}
}
