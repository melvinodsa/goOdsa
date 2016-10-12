package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/melvinodsa/goOdsa/modules/cmd"
	"github.com/melvinodsa/goOdsa/modules/forwardt"
)

func main() {
	args := os.Args
	//Default message when no cmd options are given
	if len(args) <= 1 {
		fmt.Println("You entered no options. Please use the help flag(-h) for knowinf more.\nEg:- odsa -h")
		fmt.Println("An examle conversion :- Text to be compressed - asdaaa\nCompressed :- ")
		x := forwardt.FTransform([]byte{'a', 's', 'd', 'a', 'a', 'a'})
		fmt.Println(x.ToString())
		return
	}
	//Mapping the cmd arguments
	cmdArgs := make(map[string]int)
	for i := 0; i < len(args); i++ {
		cmdArgs[args[i]] = i
	}
	//CHecking whether help is the option
	if cmdArgs[cmd.HELP] != 0 {
		fmt.Println(cmd.HelpOptions())
		return
	}
	//Checking whether the input file is found or not
	if cmdArgs[cmd.INPUTFILE] == 0 || cmdArgs[cmd.INPUTFILE]+1 >= len(args) {
		fmt.Println("No input file found")
		return
	}
	input, err := ioutil.ReadFile(args[cmdArgs[cmd.INPUTFILE]+1])
	if err != nil {
		fmt.Println("Couldn't read the file " + args[cmdArgs[cmd.INPUTFILE]+1])
		return
	}
	x := forwardt.FTransform(input)
	if cmdArgs[cmd.OUTPUTFILE] != 0 {
		ioutil.WriteFile(args[cmdArgs[cmd.OUTPUTFILE]+1], x.GetData(), 0644)
		return
	}
	fmt.Println(x.GetData())

}
