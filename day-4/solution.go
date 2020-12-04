package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "strconv"
  "regexp"
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

type Passport struct {
  fields map[string]string
}

func (p *Passport) is_valid() bool {
  required_fields := []string{"byr","iyr","eyr","hgt","hcl","ecl","pid"}
  if len(p.fields) < 7 {
    return false
  }
  for _, f := range required_fields {
    if val, ok := p.fields[f]; ok {
      switch f {
      case "byr":
        // Year between 1920 and 2002
        year, err := strconv.Atoi(val)
        if err != nil || year < 1920 || year > 2002 {
          return false
        }

      case "iyr":
        // Year between 2010 and 2020
        year, err := strconv.Atoi(val)
        if err != nil || year < 2010 || year > 2020 {
          return false
        }

      case "eyr":
        // Year between 2020 and 2030
        year, err := strconv.Atoi(val)
        if err != nil || year < 2020 || year > 2030 {
          return false
        }

      case "hgt":
        // Number followed by "cm" or "in" and within ranges
        metric := val[len(val) - 2:]
        if metric == "cm" {
          height, err := strconv.Atoi(val[:len(val) - 2])
          if err != nil || height < 150 || height > 193 {
            return false
          }
        } else if metric == "in" {
          height, err := strconv.Atoi(val[:len(val) - 2])
          if err != nil || height < 59 || height > 76 {
            return false
          }
        } else {
          return false
        }

      case "hcl":
        // '#' followed by exactly 6 hex characters
        match, err := regexp.MatchString("^#[0-9a-f]{6}$", val)
        if !match || err != nil {
          return false
        }

      case "ecl":
        // One of the below allowed values
        allowed := []string{"amb","blu","brn","gry","grn","hzl","oth"}
        ok := false
        for _, col := range allowed {
          if col == val {
            ok = true
            break
          }
        }
        if !ok {
          return false
        }

      case "pid":
        // 9 digit number
        match, err := regexp.MatchString("^[0-9]{9}$", val)
        if !match || err != nil {
          return false
        }

      default:
        fmt.Println("Unrecognised field:", val)
      }
    } else {
      return false
    }
  }
  return true
}

func parse_passports(file_lines []string) []*Passport {
  fields := map[string]string{}
  var passports []*Passport
  for _, line := range file_lines {
    if line == "" {
      passports = append(passports, &Passport{fields})
      fields = map[string]string{}
    } else {
      fs := strings.Split(line, " ")
      for _, f := range fs {
        entry := strings.Split(f, ":")
        fields[entry[0]] = entry[1]
      }
    }
  }
  return append(passports, &Passport{fields})
}


func main() {
  file_lines := read_file_lines("input.txt")
  num_valid := 0
  for _, passport := range(parse_passports(file_lines)) {
    if passport.is_valid() {
      num_valid += 1
    }
  }
  fmt.Println("Number of valid passports:", num_valid)
}
