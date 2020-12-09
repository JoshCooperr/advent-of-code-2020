package main

import (
  "fmt"
  "os"
  "io/ioutil"
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

func check_sum_exists(nums []int, sum int) bool {
  for i := 0; i < len(nums); i++ {
    for j := i + 1; j < len(nums); j++ {
      if (nums[i] + nums[j]) == sum {
        return true
      }
    }
  }
  return false
}

func find_non_sum(nums []int, step int) int {
  i := step
  found := false
  for !found || i < len(nums) {
    if !check_sum_exists(nums[i - step:i], nums[i]) {
      return nums[i]
    }
    i += 1
  }
  return 0
}

func find_contiguous_set(nums []int, sum int) []int {
  for i := range nums {
    acc, j := 0, i
    ns := []int{}
    for acc < sum {
      acc += nums[j]
      ns = append(ns, nums[j])
      if acc == sum {
        return ns
      }
      j += 1
    }
  }
  return []int{}
}

func get_min_max(list []int) (int, int) {
  min := list[0]
  max := list[0]
  for _, n := range list {
    if n < min {
      min = n
    }
    if n > max {
      max = n
    }
  }
  return min, max
}

func main() {
  nums := read_file_lines("input.txt")
  non_sum := find_non_sum(nums, 25)
  fmt.Println("First number to break the rule:", non_sum)
  cont_set := find_contiguous_set(nums, non_sum)
  min, max := get_min_max(cont_set)
  fmt.Println("Encryption weakness:", min + max)
}
