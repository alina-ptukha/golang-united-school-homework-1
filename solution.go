package solution

import (

    "github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
    hello := "Hello"
    world := ":world_map:"
    task := emoji.Sprint(hello, world)
    task = task + " !"
    return task

}
