package main

import (
  "fmt"
  "math"
  "os"
  "io/ioutil"
  "strconv"
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

type Instruction struct {
  action byte
  value int
}

type Waypoint struct {
  x int
  y int
}

type Ship struct {
  xPos int // positive: East / negative: West
  yPos int // positive: North / negative: South
  rotation byte // One of {N,E,S,W}
  waypoint Waypoint // Location relative to that of the ship
}

func (s *Ship) rotate(dir byte, amount int) {
  times := amount / 90
  for times > 0 {
    if dir == 'R' {
      switch s.rotation {
      case 'N':
        s.rotation = 'E'
      case 'E':
        s.rotation = 'S'
      case 'S':
        s.rotation = 'W'
      case 'W':
        s.rotation = 'N'
      }
    } else if dir == 'L' {
      switch s.rotation {
      case 'N':
        s.rotation = 'W'
      case 'E':
        s.rotation = 'N'
      case 'S':
        s.rotation = 'E'
      case 'W':
        s.rotation = 'S'
      }
    } else {
      fmt.Println("Unknown direction passed to ship.rotate()")
      os.Exit(1)
    }
    times -= 1
  }
}

func (s *Ship) rotate_waypoint(dir byte, amount int) {
  times := amount / 90
  for times > 0 {
    x, y := s.waypoint.x, s.waypoint.y
    if dir == 'R' {
      s.waypoint.x = y
      s.waypoint.y = -1 * x
    } else if dir == 'L' {
      s.waypoint.x = -1 * y
      s.waypoint.y = x
    } else {
      fmt.Println("Unknown direction passed to ship.rotate_waypoint()")
      os.Exit(1)
    }
    times -= 1
  }

}

func (s *Ship) move(amount int) {
  s.xPos += s.waypoint.x * amount
  s.yPos += s.waypoint.y * amount
}

func (s *Ship) move_waypoint(dir byte, amount int) {
  switch dir {
  case 'N':
    s.waypoint.y += amount
  case 'S':
    s.waypoint.y -= amount
  case 'E':
    s.waypoint.x += amount
  case 'W':
    s.waypoint.x -= amount
  }
}

func (s *Ship) navigate(instr *Instruction) {
  switch instr.action {
  case 'L':
    s.rotate_waypoint('L', instr.value)
  case 'R':
    s.rotate_waypoint('R', instr.value)
  case 'F':
    s.move(instr.value)
  default:
    s.move_waypoint(instr.action, instr.value)
  }
}

func parse_instructions(lines []string) []*Instruction {
  instructions := []*Instruction{}
  for _, line := range lines {
    value, _ := strconv.Atoi(line[1:])
    instructions = append(instructions, &Instruction{line[0], value})
  }
  return instructions
}

func main() {
  instrs := parse_instructions(read_file_lines("input.txt"))
  ship := &Ship{0, 0, 'E', Waypoint{10, 1}}
  for _, i := range instrs {
    ship.navigate(i)
  }
  fmt.Println("Manhattan Distance:", math.Abs(float64(ship.xPos)) + math.Abs(float64(ship.yPos)))
}
