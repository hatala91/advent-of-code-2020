package main


import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func is_password_valid(line string) bool {
    double_point_split := strings.Split(line, ":")
    convention, password := double_point_split[0], double_point_split[1]

    space_split := strings.Split(convention, " ")
    cardinality, letter := space_split[0], space_split[1]
    
    minus_split := strings.Split(cardinality, "-")
    lower, _ := strconv.Atoi(minus_split[0])
    upper, _ := strconv.Atoi(minus_split[1])

    n_letters := count_letter_in_password(password, letter)

    if (n_letters >= lower && n_letters <= upper) {
        return true
    } else {
        return false
    }
}

func count_letter_in_password(password, letter string) int {
    return strings.Count(password, letter)
}

func main() {
    content, _ := ioutil.ReadFile("input.txt")
    lines := strings.Split(string(content), "\n")

    var valid_passwords int = 0

    for _, line := range lines {
        if (len(line) > 0 && is_password_valid(line)) {
            valid_passwords++
        }
    }

    fmt.Printf("%v valid passwords\n", valid_passwords)
}
