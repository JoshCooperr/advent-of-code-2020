package main

import (
  "fmt"
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

func find_num_arrangements(list []int) int {
  sub_seqs := [][]int{}
  seq := []int{0}
  for _, a := range list {
    seq = append(seq, a)
    if len(seq) > 1 && (a - seq[len(seq) - 2]) == 3 {
      sub_seqs = append(sub_seqs, seq)
      seq = []int{}
    }
  }
  if len(seq) > 0 {
    sub_seqs = append(sub_seqs, seq)
  }
  fmt.Println(sub_seqs)
  return 0
}

func main() {
  nums := read_file_lines("test-input.txt")
  sort.Ints(nums)
  j1, j2, j3 := find_jolt_diffs(nums)
  fmt.Println("Jolt differences:", j1, j2, j3)
  fmt.Println(j1 * j3)
  find_num_arrangements(nums)
}
