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
    first, _ := strconv.Atoi(minus_split[0])
    second, _ := strconv.Atoi(minus_split[1])

    n_letters := count_letter_in_password(password, letter, first, second)

    if (n_letters == 1) {
        return true
    } else {
        return false
    }
}

func count_letter_in_password(password string, letter string, first int, second int) int {
    n_letter := 0

    if (string(password[first - 1]) == letter) {
        n_letter++
    }
    
    if (string(password[second - 1]) == letter) {
        n_letter++
    }

    return n_letter
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
