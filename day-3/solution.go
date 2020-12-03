package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

type direction struct {
  x int
  y int
}

func read_file_lines(fileName string) []string {
  fileBytes, err := ioutil.ReadFile(fileName)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  list := strings.Split(string(fileBytes), "\n")
  return list[:len(list) - 1]
}

func traverse_slope(slope [][]rune, move direction, c chan int) {
  x, y, trees := 0, 0, 0
  for y < len(slope) - 1 {
    y += move.y
    x = (x + move.x) % len(slope[y])
    if slope[y][x] == '#' {
      trees += 1
    }
  }
  c <- trees
}

func product(c chan int, num int) int {
  prod := 1
  for i := 0; i < num; i++ {
    val := <-c
    prod *= val
  }
  return prod
}

func main() {
  file_lines := read_file_lines("input.txt")
  var slope [][]rune
  for _, line := range file_lines {
    slope = append(slope, []rune(line))
  }
  moves := []direction{
    direction{1,1},
    direction{3,1},
    direction{5,1},
    direction{7,1},
    direction{1,2},
  }
  c := make(chan int, len(moves))
  for _, move := range moves {
    go traverse_slope(slope, move, c)
  }
  fmt.Println("Multiplication of all trees:", product(c, len(moves)))
}
