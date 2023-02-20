package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "regexp"
)

func main() {
    // Change to your preferred folder/files!
    directory := filepath.Join(os.Getenv("GITHUB_WORKSPACE"), "content/experimental")
    outputPath := filepath.Join(os.Getenv("GITHUB_WORKSPACE"), "assets/go/lists.txt")

    // We use Regex to get the link values
    re := regexp.MustCompile(`^link:\s*"(.*)"$`)

    outputFile, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer outputFile.Close()

    // Recursive search
    err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
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
                    linkString := fmt.Sprintf("%s\n", link)
                    _, err := outputFile.WriteString(linkString)
                    if err != nil {
                        fmt.Println(err)
                        return nil
                    }
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
