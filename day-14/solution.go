package main

import (
  "fmt"
  "math"
  "os"
  "io/ioutil"
  "strings"
  "strconv"
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

type Instruction struct {
  mask string
  index int
  value string
}

type Programme struct {
  memory map[int]string
  instructions []Instruction
}

func parse_programme(lines []string) *Programme {
  mask := strings.Split(lines[0], " = ")[1]
  instrs := []Instruction{}
  for _, line := range lines[1:] {
    if strings.Split(line, " = ")[0] == "mask" {
      mask = strings.Split(line, " = ")[1]
    } else {
      mem := strings.Split(strings.Split(line, " = ")[0], "[")[1]
      value, _ := strconv.Atoi(strings.Split(line, " = ")[1])
      index, _ := strconv.Atoi(strings.Split(mem, "]")[0])
      instrs = append(instrs, Instruction{mask, index, strconv.FormatInt(int64(value), 2)})
   }
  }
  return &Programme{map[int]string{}, instrs}
}

func pad_zeroes(num string, bits int) string {
  padded := ""
  for bits > len(num) {
    padded += "0"
    bits -= 1
  }
  padded += num
  return padded
}

func (i *Instruction) apply_value_mask() string {
  padded := pad_zeroes(i.value, len(i.mask))
  res := []rune{}
  for i, m := range i.mask {
    if m != 'X' {
      res = append(res, m)
    } else {
      res = append(res, []rune(padded)[i])
    }
  }
  return string(res)
}

func (i *Instruction) get_indexes() []int {
  addr := pad_zeroes(strconv.FormatInt(int64(i.index), 2), len(i.mask))
  res := []rune{}
  num_floating := 0
  for i, m := range i.mask {
    if m == '0' {
      res = append(res, []rune(addr)[i])
    } else {
      if m == 'X' {
        num_floating += 1
      }
      res = append(res, m)
    }
  }
  indexes := []int{}
  for i := 0; i < int(math.Pow(2, float64(num_floating))); i++ {
    floating_vals := []rune(pad_zeroes(strconv.FormatInt(int64(i), 2), num_floating))
    j := 0
    index := []rune{}
    for _, bit := range res {
      if bit == 'X' {
        index = append(index, floating_vals[j])
        j += 1
      } else {
        index = append(index, bit)
      }
    }
    index_int, _ := strconv.ParseInt(string(index), 2, 64)
    indexes = append(indexes, int(index_int))
  }
  return indexes
}

func (p *Programme) execute() {
  for _, instr := range p.instructions {
    indexes := instr.get_indexes()
    for _, index := range indexes {
      p.memory[index] = instr.value
    }
  }
}

func (p *Programme) sum_memory() int {
  sum := int64(0)
  for _, val := range p.memory {
    i, _ := strconv.ParseInt(val, 2, 64)
    sum += i
  }
  return int(sum)
}

func main() {
  prog := parse_programme(read_file_lines("input.txt"))
  prog.execute()
  fmt.Println("Part 2:", prog.sum_memory())
}
