package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "strconv"
)

func read_file_lines(file_name string) []string {
  file_bytes, err := ioutil.ReadFile(file_name)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  list := strings.Split(string(file_bytes), "\n")
  return list[:len(list) - 1]
}

type Bag struct {
  colour string
  contains map[*Bag]int
}

func parse_bags(file_lines []string) map[string]*Bag {
  bags := map[string]*Bag{}
  for _, line := range file_lines {
    info := strings.Split(line, " bags contain ")
    bag_colour := info[0]
    bags[bag_colour] = &Bag{bag_colour, map[*Bag]int{}}
    if info[1] != "no other bags." {
      contents := strings.Split(info[1], ", ")
      for _, con := range contents {
        inf := strings.Split(con, " ")
        col := inf[1] + " " + inf[2]
        amt, _ := strconv.Atoi(inf[0])
        if b, ok := bags[col]; ok {
          bags[bag_colour].contains[b] = amt
        } else {
          bags[col] = &Bag{col, map[*Bag]int{}}
          bags[bag_colour].contains[bags[col]] = amt
        }
      }
    }
  }
  return bags
}

func bag_contains(bags map[string]*Bag, colour string, desired string) bool {
  for pos, _ := range bags[colour].contains {
    if pos.colour == desired {
      return true
    } else if bag_contains(bags, pos.colour, desired) {
      return true
    }
  }
  return false
}

func num_bags(bags map[string]*Bag, bag *Bag) int {
  sum := 1
  for b, amt := range bag.contains {
    sum += amt * num_bags(bags, bags[b.colour])
  }
  return sum
}

func main() {
  file_lines := read_file_lines("input.txt")
  bags := parse_bags(file_lines)
  desired := "shiny gold"
  count := 0
  for col, _ := range bags {
    if col != desired && bag_contains(bags, col, desired) {
      count += 1
    }
  }
  fmt.Println("Number of bags that can hold", desired, ":", count)
  fmt.Println("Total number of bags", num_bags(bags, bags[desired]) - 1)
}
