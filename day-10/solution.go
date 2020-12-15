package main

import (
  "fmt"
  "math"
  "os"
  "io/ioutil"
  "sort"
  "strconv"
  "strings"
)

func read_file_lines(fileName string) []int {
  fileBytes, err := ioutil.ReadFile(fileName)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  list := strings.Split(string(fileBytes), "\n")
  nums := []int{}
  for _, l := range list[:len(list) - 1] {
    n, _ := strconv.Atoi(l)
    nums = append(nums, n)
  }
  return nums
}

func find_jolt_diffs(list []int) (int, int, int) {
  jolt_diffs := [3]int{}
  prev := 0
  for _, a := range list {
    jolt_diffs[a - prev - 1] += 1
    prev = a
  }
  return jolt_diffs[0], jolt_diffs[1], jolt_diffs[2] + 1
}

func count_arrangements(list []int) int {
  memo := map[int]int{}
  memo[0] = 1
  for _, n := range list {
    sum := 0
    if val, ok := memo[n-3]; ok {
      sum += val
    }
    if val, ok := memo[n-2]; ok {
      sum += val
    }
    if val, ok := memo[n-1]; ok {
      sum += val
    }
    memo[n] = sum
  }
  return memo[list[len(list) - 1]]
}

func main() {
  nums := read_file_lines("input.txt")
  sort.Ints(nums)
  j1, j2, j3 := find_jolt_diffs(nums)
  fmt.Println("Jolt differences:", j1, j2, j3)
  fmt.Println(j1 * j3)
  fmt.Println(count_arrangements(nums))
}
