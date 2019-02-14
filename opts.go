package main

import (
  "fmt"
  "math/rand"
  "strconv"
)

func ask() {
  inp := read("What would you like to do?\n 1: Visit shop\n 2: Check inventory\n 3: Play a game\n > ")
  switch inp {
    case "1":
    shop()
    
    case "2":
    fmt.Printf("\nINVENTORY\n NAME:   %s\n CLASS:  %s\n MONEY:  %d\n HEALTH: %d\n POWER:  %d\n\n", player.name, player.class, player.money, player.health, player.power)

    case "3":
    game()

    default:
    fmt.Printf("'%s' is not a valid option.\n", inp)
  }
  ask()
}

func shop() {
  fmt.Println("Welcome to the shop!")
  choice := read("What would you like to buy?\n 1: Upgraded health ($50)\n 2: Upgraded power ($50)\n >>> ")
  switch choice {
    case "1":
    if player.money >= 50 {
      fmt.Println("Bought upgraded health")
      player.health += 3
      player.money -= 50
    } else {
      fmt.Println("Not enough money :/")
    }

    case "2":
    if player.money >= 50 {
      fmt.Println("Bought upgraded power")
      player.power++
      player.money -= 50
    } else {
      fmt.Println("Not enough money :/")
    }

    default:
    fmt.Printf("'%s' is not a valid choice...")
    shop()
  }
}

func game() {
  fmt.Println("Fight!")
  game_int := rand.Intn(2)
  switch game_int {
    case 0:
    fmt.Print("Game is **Guess**")
    numb := rand.Intn(10) + 1
    for i := 0; i < 3; i++ {
      guessstr := read("Take a guess! (1-10):\n >>>")
      guess, err := strconv.Atoi(guessstr)
      if err != nil{fmt.Printf("Err\n\t%v", err)}
      if guess > numb {
        fmt.Println("Guess lower!")
      } else if guess < numb {
        fmt.Println("Guess higher!")
      } else {
        fmt.Println("You win $10!")
        player.money += 10
        break
      }
    }
    fmt.Printf("Number was %d\n", numb)


    case 1:
    fmt.Println("Game is **Dice**")
    numb := rand.Intn(6) + 1
    for i := 0; i < 2; i++ {
      guessstr := read("Make your bet (1-6)\n >>> ")
      guess, err := strconv.Atoi(guessstr)
      if err != nil{fmt.Printf("Err\n\t%v", err)}
      if guess > numb {
        fmt.Println("Guess lower!")
      } else if guess < numb {
        fmt.Println("Guess higher!")
      } else {
        fmt.Println("You win $30!")
        player.money += 30
        break
      }
    }
    fmt.Printf("Number was %d\n", numb)
  }
}
