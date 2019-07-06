//package main

package n_tape

import (
    "bufio"
    "fmt"
    "os"
)

type tape struct {
    memory []string
}

/**
 * @brief      Reads lines.
 *
 * @param      path  The path
 *
 * @return     { description_of_the_return_value }
 */
func (t *tape) readlines(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if scanner.Err() != nil {
        return scanner.Err()
    } else {
        t.memory = lines
        return scanner.Err()
    }
}

func main() {
    var t tape
    t.readlines("Fathers.txt")
    for i := 0; i < 10; i++ {
        fmt.Println(t.memory[i])
    }
}
