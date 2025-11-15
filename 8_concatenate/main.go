package main

import (
	"fmt"
	"strings"
)

func main() {
	ott := []string{"Netflix", "Hulu", "Disney+", "Amazon Prime"}
	fmt.Println(strings.Join(ott, ", "))

}
