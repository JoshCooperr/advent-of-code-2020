package main

import (
  "fmt"
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

func parse_timetable(lines []string) (int, []int) {
  depart, _ := strconv.Atoi(lines[0])
  buses := []int{}
  for _, b := range strings.Split(lines[1], ",") {
    if b == "x" {
      buses = append(buses, 1)
    } else {
      bus, _ := strconv.Atoi(b)
      buses = append(buses, bus)
    }
  }
  return depart, buses
}

func find_earliest_bus(time int, buses []int) (int, int) {
  first_bus := buses[0]
  min_time := first_bus - (time % first_bus)
  for _, bus := range buses {
    if bus != 1 {
      wait_time := bus - (time % bus)
      if wait_time < min_time {
        min_time = wait_time
        first_bus = bus
      }
    }
  }
  return min_time, first_bus
}

func part_two(buses []int) int {
  time := buses[0]
  interval := buses[0]
  for i, bus := range buses[1:] {
    for (time + i + 1) % bus != 0 {
      time += interval
    }
    interval *= bus
  }
  return time
}

func main() {
  departure, buses := parse_timetable(read_file_lines("input.txt"))
  time, bus := find_earliest_bus(departure, buses)
  fmt.Println("Part 1:", time * bus)
  fmt.Println("Part 2:", part_two(buses))
}
