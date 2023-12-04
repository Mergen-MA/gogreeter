package main

import "fmt"

func Greet(name string) string {
    cleaned_name = strings.TrimSpace(name)
    if len(cleaned_name) == 0 {
        return "Hello World!"
    }
    return fmt.Sprintf("Hello %s!", name)
}
