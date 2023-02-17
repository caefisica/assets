package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {
    file, err := os.Open("example.md")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Using Regex to get the value
    re := regexp.MustCompile(`^link:\s*"(.*)"$`)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        match := re.FindStringSubmatch(line)
        if len(match) == 2 {
            link := match[1]
            fmt.Printf("Link: %s\n", link)
            break
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
