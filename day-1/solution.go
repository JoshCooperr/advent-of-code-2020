package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "sort"
  "strconv"
  "strings"
)

func read_file(fileName string) []int {
  fileBytes, err := ioutil.ReadFile(fileName)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  list := strings.Split(string(fileBytes), "\n")
  var list_ints = []int{}
  for _, i := range list {
    converted_i, _ := strconv.Atoi(i)
    list_ints = append(list_ints, converted_i)
  }
  return list_ints[:len(list_ints) - 1]
}

func get_pair(list []int, target int) [2]int {
  // Sort the list
  sort.Slice(list, func(i, j int) bool {
    return list[i] < list[j]
  })
  // Initialise found boolean and pointers
  found := false;
  i, j := 0, len(list) - 1
  // Search through the slice for the pair
  var answer [2]int
  for !found {
    sum := list[i] + list[j]
    if sum == target {
      answer = [2]int{list[i], list[j]}
      found = true
    } else if sum < target {
      i += 1
    } else { // sum > target
      j -= 1
    }
  }
  return answer
}

func get_triplet(list []int, target int) [3]int {
  // Sort the list
  sort.Slice(list, func(i, j int) bool {
    return list[i] < list[j]
  })
  // Initialise found boolean and pointers
  found := false;
  i, j, k := 0, 1, len(list) - 1
  var answer [3]int
  for !found {
    // i is fixed in each loop
    for j < k {
      sum := list[i] + list[j] + list[k]
      if sum == target {
        answer = [3]int{list[i], list[j], list[k]}
        found = true
        break
      } else if sum < target {
        j += 1
      } else { // sum > target
        k -= 1
      }
    }
    i += 1
    j, k = i + 1, len(list) - 1
  }
  return answer
}

func main() {
  list := read_file("input.txt")
  pair := get_pair(list, 2020)
  pair_answer := pair[0] * pair[1]
  triplet := get_triplet(list, 2020)
  triplet_answer := triplet[0] * triplet[1] * triplet[2]
  fmt.Println("Pair answer:", pair, pair_answer)
  fmt.Println("Triplet answer:", triplet, triplet_answer)
}
