package main

import (
	"fmt"
	"os"

	"github.com/melvinodsa/goOdsa/modules/forwardt"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("You entered no options. Please use the help flag(-h) for knowinf more.\nEg:- odsa -h")
		fmt.Println("An examle conversion :- Text to be compressed - asdaaa\nCompressed :- ")
		x := forwardt.FTransform("asdaaa")
		fmt.Println(x.ToString())
		return
	}
	cmdArgs := make(map[string]int)
	for i := 0; i < len(args); i++ {
		cmdArgs[args[i]] = i
	}
}
