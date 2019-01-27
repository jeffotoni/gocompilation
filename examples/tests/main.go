// Go in action
// @jeffotoni
// 2019-01-24

//-ldflags "-X main.x=2 -X main.y=3"
package main

import "strconv"

import (
	"github.com/jeffotoni/gocompilation/example/tests/pkg/math"
)

var (
	x, y   string
	xi, yi int
)

func init() {
	xi, _ = strconv.Atoi(x)
	yi, _ = strconv.Atoi(y)
}

func main() {

	println(math.Sum(xi, yi))
}
