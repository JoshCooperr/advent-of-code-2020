package main

import (
  "fmt"
  "io/ioutil"
  "os"
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

// part 1
func valid_password(line string) bool {
  parts := strings.Split(line, " ")
  min, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
  max, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])
  char := strings.Split(parts[1], ":")[0]
  password := parts[2]
  count := strings.Count(password, char)
  return (count >= min && count <= max)
}

// part 2
func valid_password_2(line string) bool {
  parts := strings.Split(line, " ")
  i, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
  j, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])
  char := []rune(strings.Split(parts[1], ":")[0])[0]
  password := []rune(parts[2])
  return (char == password[i - 1]) != (char == password[j - 1])
}

func main() {
  lines := read_file_lines("input.txt")
  count := 0
  count2 := 0
  for _, line := range lines {
    if valid_password(line) {
      count += 1
    }
    if valid_password_2(line) {
      count2 += 1
    }
  }
  fmt.Println("Criteria one:", count)
  fmt.Println("Criteria two:", count2)
}
