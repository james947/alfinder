package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var systemDirs = []string{
    "/bin",
    "/boot",
    "/dev",
    "/etc",
    "/lib",
    "/lib64",
    "/proc",
    "/root",
    "/run",
    "/sbin",
    "/sys",
    "/tmp",
    "/usr",
    "/var",
}

func findFileOrFolder(directory string, name string, fileType string, results chan<- string) {
    filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println("Error accessing:", path, "-", err)
            return nil
        }

        baseName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
        matchesName := strings.Contains(strings.ToLower(baseName), strings.ToLower(name))
        matchesType := true

        if fileType != "" {
            matchesType = false
            switch fileType {
            case "pdf":
                matchesType = strings.HasSuffix(strings.ToLower(path), ".pdf")
            case "img":
                matchesType = strings.HasSuffix(strings.ToLower(path), ".jpg") ||
                    strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
                    strings.HasSuffix(strings.ToLower(path), ".png") ||
                    strings.HasSuffix(strings.ToLower(path), ".gif") ||
                    strings.HasSuffix(strings.ToLower(path), ".bmp")
            case "txt":
                matchesType = strings.HasSuffix(strings.ToLower(path), ".txt")
            }
        }

        if matchesName && matchesType {
            results <- path
        }
        return nil
    })
}

func main() {
    var directory, name, fileType string

    switch len(os.Args) {
    case 2:
        directory = "/" // default to root for global search
        name = os.Args[1]
    case 3:
        directory = "/" // default to root for global search
        name = os.Args[1]
        fileType = os.Args[2][1:] // remove the leading '-'
    case 4:
        directory = os.Args[1]
        name = os.Args[2]
        fileType = os.Args[3][1:] // remove the leading '-'
    default:
        fmt.Printf("Usage: %s [directory] <name> [<file type>]\n", os.Args[0])
        os.Exit(1)
    }

    fmt.Printf("Searching for '%s' in directory '%s' with file type '%s'\n", name, directory, fileType)

    results := make(chan string, 100) // increase buffer size for better performance

    // Search user directories first(performace kinda thing :) )
    go func() {
        if directory == "/" {
            for _, dir := range systemDirs {
                filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
                    if err != nil {
                        return filepath.SkipDir
                    }
                    return nil
                })
            }
            filepath.Walk("/home", func(path string, info os.FileInfo, err error) error {
                if err == nil {
                    findFileOrFolder(path, name, fileType, results)
                }
                return nil
            })
            findFileOrFolder("/Users", name, fileType, results)
        } else {
            findFileOrFolder(directory, name, fileType, results)
        }
        close(results)
    }()

    found := false
    for result := range results {
        fmt.Println("Found:", result)
        found = true
    }

    if !found {
        fmt.Println("No matching files or folders found.")
    }
}
