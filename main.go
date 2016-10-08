package main

import (
	"fmt"

	"github.com/melvinodsa/goOdsa/forwardt"
)

func main() {
	x := forwardt.FTransform("asdaaa")
	fmt.Println(x.ToString())
}
