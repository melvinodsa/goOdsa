/*Package cmd contains the list of the options
available for the cmd application. Also it has the
function which returns the help text for the application
*/
package cmd

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
		"\t\tHelp Options         -h  \t\t\t----- for help" +
		"\t\tInput File           -if \t\t\t----- input file name" +
		"\t\tOutput File          -of \t\t\t----- output file name. Console output incase of no output file specified" +
		"\t\tDo not compress data -dc \t\t\t----- just print the compressed data state"
	return helpOptions
}
