package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  file, err := os.Open("data.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  sum := 0
  for scanner.Scan() {
    scanInt, err := strconv.Atoi(scanner.Text())
    if err != nil {
      panic(err)
    }
    fuelWeight := checkFuelRequired(scanInt)
    for fuelWeight > 0 {
      sum += fuelWeight
      fuelWeight = checkFuelRequired(fuelWeight)
    }
  }
  fmt.Println(sum)
}

