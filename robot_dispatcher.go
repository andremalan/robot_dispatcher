package main

import "fmt"

type Dispatcher struct {
}

type Robot struct {

    id int
    name string
}

func main() {
  robot := new(Robot)
  robot.id = 5
  robot.name = "AUTOBOT"
  fmt.Printf("bot name is %s and id is %d", robot.name, robot.id)
}

