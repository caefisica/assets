package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "regexp"
)

func main() {
    directory := "./folder"

    re := regexp.MustCompile(`^link:\s*"(.*)"$`)

    // Recursive search
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
            return nil
        }

        if filepath.Ext(path) == ".md" {
            file, err := os.Open(path)
            if err != nil {
                fmt.Println(err)
                return nil
            }
            defer file.Close()

            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                line := scanner.Text()
                match := re.FindStringSubmatch(line)
                if len(match) == 2 {
                    link := match[1]
                    fmt.Printf("File: %s, Link: %s\n", path, link)
                    break
                }
            }

            if err := scanner.Err(); err != nil {
                fmt.Println(err)
            }
        }

        return nil
    })

    if err != nil {
        fmt.Println(err)
    }
}
