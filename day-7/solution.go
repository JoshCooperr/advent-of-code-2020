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

type ExpandedBag struct {
  colour string
  contains map[string]int
  capacity int
}

func (b *Bag) can_hold(col string) bool {
  fmt.Println(b.colour, b.contains)
  for pos, _ := range b.contains {
    if pos.colour == col {
      return true
    } else if pos.can_hold(col) {
      return true
    }
  }
  return false
}

func (b *ExpandedBag) can_hold(col string) bool {
  _, ok := b.contains[col]
  return ok
}

func parse_bags(file_lines []string) map[string]*ExpandedBag {
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

  expanded_bags := map[string]*ExpandedBag{}
  for col, bag := range bags {
    contains := map[string]int{}
    capacity := 0
    to_process := []string{}
    for b, _ := range bag.contains {
      to_process = append(to_process, b.colour)
    }
    for len(to_process) > 0 {
      c := to_process[0]
      to_process = to_process[1:]
      contains[c] = 0
      for x, v := range bags[c].contains {
        if _, ok := contains[x.colour]; !ok {
          to_process = append(to_process, x.colour)
        }
        capacity += v
      }
    }
    fmt.Println(col, contains, capacity)
    expanded_bags[col] = &ExpandedBag{col, contains, capacity}
  }
  return expanded_bags
}

func main() {
  file_lines := read_file_lines("test-input.txt")
  bags := parse_bags(file_lines)
  desired := "shiny gold"
  count := 0
  for col, bag := range bags {
    if col != desired && bag.can_hold(desired) {
      count += 1
    }
  }
  fmt.Println("Number of bags that can hold", desired, ":", count)
}
