package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <url>\n", os.Args[0])
        os.Exit(1)
    }
    url := os.Args[1]

    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Fprintf(os.Stderr, "fetch: %s returned status %d\n", url, resp.StatusCode)
        os.Exit(1)
    }

    _, err = io.Copy(os.Stdout, resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading body: %v\n", err)
        os.Exit(1)
    }
}
