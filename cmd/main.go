package main

import (
	"fmt"
	"time"
)

func main() {

	t, err := time.Parse("02/01/2006 15:04:05", "13/11/2023 18:30:57")
	if err != nil {

	}

	t = t.Add(-7 * time.Hour)

	fmt.Println(t.Unix())
}
