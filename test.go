package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn :=false
  eludeGuards := rand.Intn(100)
  if eludeGuards <=50 {
  fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
  fmt.Println("Plan a better disguise next time?")
  }
  fmt.Printf("IsHeistOn is currently: %v",isHeistOn)
  fmt.Println(".")
  
openedVault := rand.Intn(100)
  if isHeistOn == true && openedVault >=70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn == false {
    fmt.Println("Vault can't be opened")
  }

  leftSafely := rand.Intn(5)
  if isHeistOn {
    switch leftSafely{
      case 0:
      isHeistOn = false
      fmt.Println("We fucked up")
      case 1:
      isHeistOn = false
      fmt.Println("We had some chances, but now we suck")
      case 2:
      isHeistOn = false
      fmt.Println("Guards have too much energy today")
      default:
      fmt.Println ("Buckle up!")
    }
      if isHeistOn {
		amtStolen := 10000+rand.Intn(1000000)
      fmt.Println("The gang got: ", amtStolen, "for today!")


    }
    
  }
}