package main

import "fmt"
//import "time"
//import "encoding/csv"

type Point struct {

  x float32
  y float32
  time string
}

func NewPoint(x float32, y float32, time string) *Point {
    p := Point{x, y, time}
    return &p
}


type Bot struct{
  robot Robot
  points [1100]Point
  index int
}

type Dispatcher struct {
  bots [2]Bot
}

func (self *Dispatcher) GetPoints (botId string) {
  fmt.Printf("bot  id is %s", botId)
}

func main() {
  botIds := [2]int{5937, 6043}
  dispatcher := new(Dispatcher)
  dispatcher.bots[0] = *new(Bot)
  dispatcher.bots[1] = *new(Bot)
  dispatcher.bots[0].robot.id = botIds[0]
  dispatcher.bots[1].robot.id = botIds[1]
  dispatcher.bots[0].points[0] = *NewPoint(3.0,4.0, "Yesterday")

  robot := dispatcher.bots[0].robot
  fmt.Printf("bot name is %s and id is %d", robot.name, robot.id)
  fmt.Printf("Dispatcher outputs:" )
  dispatcher.GetPoints("MYID")
  point := dispatcher.bots[0].points[0]
  fmt.Printf("the first point is: %f %f %s", point.x, point.y, point.time)
}

