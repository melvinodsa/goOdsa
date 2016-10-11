/*Package cmd contains the list of the options
available for the cmd application. Also it has the
function which returns the help text for the application
*/
package cmd

//HELP is the option constant for help
const HELP = "-h"

//INPUTFILE is the input filename option
const INPUTFILE = "-if"

//OUTPUTFILE is the output filename option
const OUTPUTFILE = "-of"

//DONOTCOMPRESS is the option for showing the compress state of the data
const DONOTCOMPRESS = "-dc"

//Options is map of all the available options in the application
var Options = map[string]bool{"-if": true,
	"-of": true,
	"-dc": true,
	"-h":  true}

/*HelpOptions will give all the help options available for the
application
*/
func HelpOptions() string {
	helpOptions := "ODSA Help\n" +
		"=========\n" +
		"\tOptions format odsa -[option] [option value] ...{more option/option value pairs}\n" +
		"\t\tHelp Options         -h  \t\t----- for help\n" +
		"\t\tInput File           -if \t\t----- input file name\n" +
		"\t\tOutput File          -of \t\t----- output file name. Console output incase of no output file specified\n" +
		"\t\tDo not compress data -dc \t\t----- just print the compressed data state\n"
	return helpOptions
}
