package main

import (
    "fmt"
    "io/ioutil"
)

func read() {
    const filename = "abc.txt"
    if contents, err := ioutil.ReadFile(filename); err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }
}

func readByIf() {
    const filename = "abc.txt"
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }
}

func grade(score int) string {
    g := ""
    //自动break 除非使用fallthrough
    switch {
    case score < 0 || score > 100:
        panic(fmt.Sprintf("Wrong score: %d", score))
    case score < 60:
        g = "F"
    case score < 80:
        g = "C"
    case score < 90:
        g = "B"
    case score <= 100:
        g = "A"
    }
    return g
}

func main() {
    //readByIf()
    //read()
    fmt.Println(
        grade(0),
        grade(59),
        grade(82),
        grade(100),
    )
}