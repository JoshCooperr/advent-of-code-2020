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

type Group struct {
  size int
  answers map[rune]int
}

func (g *Group) num_yes() int {
  yes := 0
  for _, cnt := range g.answers {
    if cnt == g.size {
      yes += 1
    }
  }
  return yes
}

func parse_groups(lines []string) []*Group {
  var groups []*Group
  answers := map[rune]int{}
  size := 0
  for _, line := range lines {
    if line == "" {
      groups = append(groups, &Group{size, answers})
      answers = map[rune]int{}
      size = 0
    } else {
      qs := []rune(line)
      for _, q := range(qs) {
        if cnt, ok := answers[q]; ok {
          answers[q] = cnt + 1
        } else {
          answers[q] = 1
        }
      }
      size += 1
    }
  }
  return append(groups, &Group{size, answers})
}

func main() {
  file_lines := read_file_lines("input.txt")
  groups := parse_groups(file_lines)
  num_answers := 0
  for _, g := range(groups) {
    num_answers += g.num_yes()
  }
  fmt.Println("Answers sum:", num_answers)
}
