package main


import (
    "fmt";
    "strings";
	"time";
)

func main() {
	fmt.Printf("hello, world\n")
    s := []string{"this", "is", "a", "joined", "string\n"};
    fmt.Printf(strings.Join(s, " "));
	fmt.Printf("The time is", time.Now())
}