package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func read_file_lines(fileName string) []string {
  fileBytes, err := ioutil.ReadFile(fileName)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  list := strings.Split(string(fileBytes), "\n")
  return list[:len(list) - 1]
}

func traverse_slope(slope [][]rune, x_move, y_move int) int {
  x, y, trees := 0, 0, 0
  for y < len(slope) - 1 {
    y += y_move
    x = (x + x_move) % len(slope[y])
    if slope[y][x] == '#' {
      trees += 1
    }
  }
  return trees
}

func main() {
  file_lines := read_file_lines("input.txt")
  var slope [][]rune
  for _, line := range file_lines {
    slope = append(slope, []rune(line))
  }
  route_1 := traverse_slope(slope, 1, 1)
  route_2 := traverse_slope(slope, 3, 1)
  route_3 := traverse_slope(slope, 5, 1)
  route_4 := traverse_slope(slope, 7, 1)
  route_5 := traverse_slope(slope, 1, 2)
  fmt.Println("Multiplication of all trees:", route_1 * route_2 * route_3 * route_4 * route_5)
}
