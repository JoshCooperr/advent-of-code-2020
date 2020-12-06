package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "sort"
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

func binary_search(code []rune, limit int, lower, upper rune) int {
  s, e := 0, limit
  for _, r := range code[:len(code) - 1] {
    if r == lower {
      e = s + ((e - s) / 2)
    } else if r == upper {
      s = s + ((e - s) / 2) + 1
    } else {
      fmt.Println("Unrecognised character:", r)
    }
  }
  if code[len(code) - 1] == lower {
    return s
  } else if code[len(code) - 1] == upper {
    return e
  } else {
    fmt.Println("Unrecognised character:", code[len(code) - 1])
    return -1
  }
}

func calc_seat_id(code string, num_row, num_seat int) int {
  row := binary_search([]rune(code[:num_row]), 127, 'F', 'B')
  seat := binary_search([]rune(code[len(code) - num_seat:]), 7, 'L', 'R')
  return (row * 8) + seat
}

func main() {
  file_lines := read_file_lines("input.txt")
  max := 0
  seats := make([]int, len(file_lines))
  for i, line := range file_lines {
    seat_id := calc_seat_id(line, 7, 3)
    if seat_id > max {
      max = seat_id
    }
    seats[i] = seat_id
  }
  fmt.Println("Maximum seat number:", max)
  sort.Slice(seats, func(i, j int) bool {
    return seats[i] < seats[j]
  })
  prev_seat := seats[0]
  for _, seat := range seats[1:] {
    if seat - prev_seat != 1 {
      fmt.Println("Possible seat:", seat - 1)
    }
    prev_seat = seat
  }
}
