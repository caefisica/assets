package main

import (
    "bufio"
    "context"
    "fmt"
    "io/ioutil"
    "os"
    "strings"

    "github.com/chromedp/chromedp"
)

func main() {
    // Set the window size to 1920 x 1080
    width, height := 1920, 1080

    // Read the list of websites from a file
    file, err := os.Open("lists.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var websites []string
    for scanner.Scan() {
        website := strings.TrimSpace(scanner.Text())
        if website != "" {
            websites = append(websites, website)
        }
    }

    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", true),
        chromedp.Flag("disable-gpu", true),
        chromedp.Flag("no-sandbox", true),
        chromedp.Flag("disable-setuid-sandbox", true),
    )

    allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    if _, err := os.Stat("results"); os.IsNotExist(err) {
        os.Mkdir("results", os.ModePerm)
    }

    for i, website := range websites {
        err := chromedp.Run(ctx, chromedp.Navigate(website))
        if err != nil {
            fmt.Println(err)
            continue
        }

        err = chromedp.Run(ctx, chromedp.EmulateViewport(int64(width), int64(height)))
        if err != nil {
            fmt.Println(err)
            continue
        }

        var buf []byte
        err = chromedp.Run(ctx, chromedp.CaptureScreenshot(&buf))
        if err != nil {
            fmt.Println(err)
            continue
        }

        filename := fmt.Sprintf("results/screenshot%d.png", i)
        err = ioutil.WriteFile(filename, buf, os.ModePerm)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("Screenshot of %s saved to %s\n", website, filename)
    }
}