package main
import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  fmt.Println("===== Part 1 =====")
  file, err := os.Open("data.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var input []int
  for scanner.Scan() {
    line := scanner.Text()
    elements := strings.Split(line, ",")
    for j := range elements {
      scanInt, err := strconv.Atoi(elements[j])
      if err != nil {
        panic(err)
      }
      input = append(input, scanInt)
    }
  }


  i := 0
  ic := make([]int, len(input))
  copy(ic, input);

  ic[1] = 12
  ic[2] = 2

  calc: for true {
     if ic[i] == 99 {
      break calc
    } else if ic[i] == 1 {
      ic[ic[i + 3]] = ic[ic[i + 1]] + ic[ic[i + 2]]
    } else if ic[i] == 2 {
      ic[ic[i + 3]] = ic[ic[i + 1]] * ic[ic[i + 2]]
    }
    i += 4
  }
  fmt.Println(ic[0])


  fmt.Println("===== Part 2 =====")

  for noun := 1; noun < 100; noun++ {
    for verb := 1; verb < 100; verb++ {
      ic := make([]int, len(input))
      copy(ic, input);
      ic[1] = noun
      ic[2] = verb
      i = 0
      loop: for true {
	if ic[i + 3] >= len(ic) || ic[i] == 99 {
	  break loop
	} else if ic[i] == 1 {
	  ic[ic[i + 3]] = ic[ic[i + 1]] + ic[ic[i + 2]]
	} else if ic[i] == 2 {
	  ic[ic[i + 3]] = ic[ic[i + 1]] * ic[ic[i + 2]]
	}
	i += 4
      }
      if ic[0] == 19690720 {
	fmt.Println((100 * noun) + verb)
      }
    }
  }
}

