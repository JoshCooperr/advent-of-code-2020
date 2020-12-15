package main

import (
  "fmt"
)

func find_nth_number(input []int, n int) int {
  turns_spoken := map[int][]int{}
  for i, num := range input {
    turns_spoken[num] = []int{i + 1}
  }
  prev := input[len(input) - 1]
  for turn := len(input) + 1; turn <= n; turn++ {
    if len(turns_spoken[prev]) == 1 {
      turns_spoken[0] = append(turns_spoken[0], turn)
      prev = 0
    } else {
      last_turns := turns_spoken[prev][len(turns_spoken[prev]) - 2:]
      diff := last_turns[1] - last_turns[0]
      turns_spoken[diff] = append(turns_spoken[diff], turn)
      prev = diff
    }
  }
  return prev
}

func main() {
  input := []int{8,13,1,0,18,9}
  fmt.Println(find_nth_number(input, 30000000))
}
