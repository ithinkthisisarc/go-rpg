package main

import (
  "fmt"
)

type Player struct {
  name   string
  money  int
  health int
  power  int
  class  string
}

var player = Player{ "", 0, 5, 3, ""}

func main() {
  name := read("Enter your name: ")
  player.name = name
  fmt.Printf("Welcome, %s\n", name)
  class_numb := read("Select your class:\n 1: Thief\n 2: Warrior\n 3: Tank\n > ")
  switch class_numb {
    case "1":
    player.class = "Thief"
    player.money = 25

    case "2":
    player.class = "Warrior"
    player.power = 6

    case "3":
    player.class = "Tank"
    player.health = 10
  }
  fmt.Printf("You selected %s!\n", player.class)
  ask()
}
