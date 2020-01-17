package main

import (
	"fmt"
	"math"

	"github.com/diegosepusoto/crash-course/03-packages/strutil"
)

func main() {
	fmt.Println("Packages")
	fmt.Println("PI: ", math.Pi)
	fmt.Println(strutil.Reverse("The quick brown fox jumps over the lazy dog"))
}
