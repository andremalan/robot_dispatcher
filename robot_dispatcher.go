package main

import "fmt"
//import "time"
import "os"
import "io"
import "encoding/csv"
import "strconv"

type Point struct {

  x float64
  y float64
  time string
}

func NewPoint(x float64, y float64, time string) *Point { p := Point{x, y, time}
    return &p
}


type Bot struct{
  robot Robot
  points [1100]Point
  index int
}

type Dispatcher struct {
  bots map[string]Bot
}

func (self *Dispatcher) GetPoints (botId string) {
  fmt.Printf("bot  id is %s", botId)
}

func (self *Dispatcher) LoadPoints (botId int) {
  stringId := strconv.Itoa(botId)
  filename :=  stringId + ".csv"
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  defer file.Close()
  reader := csv.NewReader(file)
  bot := self.bots[stringId]
  var index int
  for {
    record, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println("Error:", err)
      return
    }
    x, _ := strconv.ParseFloat(record[1], 64)
    y, _ := strconv.ParseFloat(record[2], 64)
    t := record[3]
    bot.points[index] =  Point{x,y,t}
    index++
  }
  self.bots[stringId] = bot
}


func main() {
  //botIds := [2]int{5937, 6043}
  bots := map[string]Bot{"5937" : *new(Bot), "6043" : *new(Bot)}
  dispatcher := Dispatcher{bots}
  dispatcher.LoadPoints(5937)
  dispatcher.LoadPoints(6043)
  fmt.Printf("dispatcher is: %v", dispatcher)

}

