package main

import (
  "fmt"
  "os"
  "io/ioutil"
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

type Ferry struct {
  seats [][]rune
}

func (f *Ferry) view_adjacent_seat(row, seat, ver, hor int) rune {
  // Returns '.' if a seat cannot be seen
  r, s := row + ver, seat + hor
  for r >= 0 && r < len(f.seats) && s >= 0 && s < len(f.seats[r]) {
    if f.seats[r][s] != '.' {
      return f.seats[r][s]
    }
    r += ver
    s += hor
  }
  return '.'
}

func (f *Ferry) check_adjacent_seats(row, seat int) int {
  cnt := 0
  for r := row - 1; r <= row + 1; r++ {
    for s := seat - 1; s <= seat + 1; s++ {
      if r == row && s == seat {
        continue
      }
      if f.view_adjacent_seat(row, seat, r - row, s - seat) == '#' {
        cnt += 1
      }
    }
  }
  return cnt
}

func (f *Ferry) simulate_passengers() {
  finished := false
  for !finished {
    finished = true
    next := [][]rune{}
    for r, row := range f.seats {
      next_row := make([]rune, len(row))
      for s, seat := range row {
        if seat == 'L' && f.check_adjacent_seats(r, s) == 0 {
          next_row[s] = '#'
          finished = false
        } else if seat == '#' && f.check_adjacent_seats(r, s) >= 5 {
          next_row[s] = 'L'
          finished = false
        } else {
          next_row[s] = row[s]
        }
      }
      next = append(next, next_row)
    }
    f.seats = next
  }
}

func (f *Ferry) count_seats() (int, int) {
  occ, not := 0, 0
  for _, row := range f.seats {
    for _, seat := range row {
      if seat == 'L' {
        not += 1
      } else if seat == '#' {
        occ += 1
      }
    }
  }
  return occ, not
}

func create_ferry(layout []string) *Ferry {
  seats := [][]rune{}
  for _, line := range layout {
    seats = append(seats, []rune(line))
  }
  return &Ferry{seats}
}

func main() {
  file_lines := read_file_lines("input.txt")
  ferry := create_ferry(file_lines)
  ferry.simulate_passengers()
  occ, _ := ferry.count_seats()
  fmt.Println("Number of occupied seats:", occ)
}
