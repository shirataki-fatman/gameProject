package systems

import (
	"engo.io/ecs"
	"fmt"
	"engo.io/engo"
)

type MainSystem struct{}

func (*MainSystem) Remove(ecs.BasicEntity) {}

func (*MainSystem) Update(dt float32) {
	if engo.Input.Button("Push").JustPressed() {
		fmt.Println("Hoge")
		engo.SetSceneByName("TitleScene", true)
	}
}
