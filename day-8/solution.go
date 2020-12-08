package main

import (
  "fmt"
  "os"
  "io/ioutil"
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

type Instruction struct {
  operation string
  argument int
  executed bool
}

func (i *Instruction) String() string {
  return fmt.Sprintf("%s %d", i.operation, i.argument)
}

type Programme struct {
  instructions []*Instruction
  acc int
}

func (p *Programme) String() string {
  instrs := []string{}
  for _, i := range p.instructions {
    instrs = append(instrs, i.String())
  }
  return fmt.Sprintf(strings.Join(instrs, "\n"))
}

func (p *Programme) execute() (int, bool) {
  // Returns the value of the accumulator, and a bool (true if no loop)
  cur_i := 0
  finished := false
  for !finished {
    instr := p.instructions[cur_i]
    if instr.executed {
      // Loop found
      return p.acc, false
    }
    switch instr.operation {
      case "acc":
        p.acc += instr.argument
        cur_i += 1
      case "jmp":
        cur_i += instr.argument
      case "nop":
        cur_i += 1
      default:
        fmt.Println("Execution error:", instr)
    }
    instr.executed = true
    if cur_i == len(p.instructions) {
      finished = true
    }
  }
  return p.acc, true
}

func (p *Programme) reset(ind int, op string) {
  p.acc = 0
  for i := 0; i < len(p.instructions); i++ {
    p.instructions[i].executed = false
    if i == ind {
      p.instructions[i].operation = op
    }
  }
}

func (p *Programme) execute_and_fix() (int, bool) {
  for i := 0; i < len(p.instructions); i++ {
    if p.instructions[i].operation == "nop" {
      p.instructions[i].operation = "jmp"
      acc, ok := p.execute()
      if ok {
        return acc, ok
      }
      p.reset(i, "nop")
    } else if p.instructions[i].operation == "jmp" {
      p.instructions[i].operation = "nop"
      acc, ok := p.execute()
      if ok {
        return acc, ok
      }
      p.reset(i, "jmp")
    }
  }
  return p.acc, false
}

func parse_programme(input []string) *Programme {
  // Allowed ops: 'acc', 'jmp', 'nop'
  allowed_ops := map[string]bool{"acc": true, "jmp": true, "nop": true}
  instructions := []*Instruction{}
  for _, line := range input {
    instr := strings.Split(line, " ")
    op, arg := instr[0], instr[1]
    if _, ok := allowed_ops[op]; !ok {
      fmt.Println("Unknown instruction:", instr)
      break
    }
    amt, _ := strconv.Atoi(arg[1:])
    if arg[0] == '-' {
      amt *= -1
    }
    instructions = append(
      instructions,
      &Instruction{op, amt, false},
    )
  }
  return &Programme{instructions, 0}
}

func main() {
  programme := parse_programme(read_file_lines("input.txt"))
  acc, ok := programme.execute_and_fix()
  if ok {
    fmt.Println("Programme terminated, accumulator value:", acc)
  } else {
    fmt.Println("Loop found, accumulator value:", acc)
  }
}
