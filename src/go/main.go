package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02_15:04:05_") + fmt.Sprint(1))
}
