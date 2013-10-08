package main

import "testing"

func TestMain(t *testing.T) {
  main()
}

func TestLoadPoints(t *testing.T) {
  bots := [2]Bot{*new(Bot), *new(Bot)}
  d := Dispatcher{bots}
  d.LoadPoints(5937)
  firstPoint := d.bots[0].points[0]
  if (firstPoint.x != 51.476105 || firstPoint.y != -0.100224 || firstPoint.time != "2011-03-22 07:55:26") {
    t.Log("First element is wrong!", firstPoint.x, firstPoint.y, firstPoint.time)
    t.Fail()
  }

}


