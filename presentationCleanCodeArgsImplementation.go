package main

import (
	"os"

	"github.com/adampresley/logging"
	"github.com/adampresley/presentationCleanCodeArgsImplementation/arguments"
)

func main() {
	log := logging.NewLogger("Clean Code Example")
	log.EnableColors()

	commandLineArgs := os.Args[1:]

	log.Debugf("Command line args: %v\n", commandLineArgs)

	var err error
	var argumentParser *arguments.ArgumentParser

	var loggingEnabled bool
	var directory string
	var port int

	/*
	 * Initialize our argument parser. We are looking for a boolean,
	 * an integer, and a string.
	 */
	if argumentParser, err = arguments.NewArgumentParser("l,p#,d*", commandLineArgs); err != nil {
		log.Error(err)
		return
	}

	if loggingEnabled, err = argumentParser.GetBoolean("l"); err != nil {
		log.Error(err)
	}

	if directory, err = argumentParser.GetString("d"); err != nil {
		log.Error(err)
	}

	if port, err = argumentParser.GetInteger("p"); err != nil {
		log.Error(err)
	}

	/*
	 * Display the results in the console
	 */
	log.Infof("Logging enabled: %v", loggingEnabled)
	log.Infof("Directory: %s", directory)
	log.Infof("Port: %d", port)
}
