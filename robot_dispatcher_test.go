package main

import "testing"

func TestMain(t *testing.T) {
  main()
}

func TestLoadPoints(t *testing.T) {
  bots := [2]Bot{*new(Bot), *new(Bot)}
  d := Dispatcher{bots}
  d.LoadPoints(5937)
  if !(d.bots[0].points[0].x == 51.4761057) {
    t.Log("First element is wrong!", d.bots[0].points[0].x)
    t.Fail()
  }

}


