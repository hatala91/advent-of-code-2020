package main


import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    content, _ := ioutil.ReadFile("input.txt")
    lines := strings.Split(string(content), "\n")
    
    for i, outer_element := range lines {
        for j, middle_element := range lines[i:] {
            for _, inner_element := range lines[j:] {
                outer_value, _ := strconv.Atoi(outer_element)
                middle_value, _ := strconv.Atoi(middle_element)
                inner_value, _ := strconv.Atoi(inner_element)
                
                if outer_value + middle_value + inner_value == 2020 {
                    result := outer_value * middle_value * inner_value
                    fmt.Printf("Result: %v\n", result)
                }
            }
        }
    }
}
